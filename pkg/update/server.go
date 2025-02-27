package update

import (
	"context"
	"sync"

	"github.com/prometheus/client_golang/prometheus"
	controlv1 "github.com/rancher/opni/pkg/apis/control/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

var _ controlv1.UpdateSyncServer = (*UpdateServer)(nil)

type UpdateServer struct {
	controlv1.UnsafeUpdateSyncServer
	updateHandlers map[string]UpdateTypeHandler
	handlerMu      sync.Mutex
}

func NewUpdateServer() *UpdateServer {
	return &UpdateServer{
		updateHandlers: make(map[string]UpdateTypeHandler),
		handlerMu:      sync.Mutex{},
	}
}

func (s *UpdateServer) RegisterUpdateHandler(strategy string, handler UpdateTypeHandler) {
	s.handlerMu.Lock()
	defer s.handlerMu.Unlock()
	s.updateHandlers[strategy] = handler
}

// SyncManifest implements UpdateSync.  It expects a manifest with a single
// type and a single strategy.  It will return an error if either of these
// conditions are not met.  The package URN must be in the following format
// urn:opni:<type>:<strategy>:<name>
func (s *UpdateServer) SyncManifest(ctx context.Context, manifest *controlv1.UpdateManifest) (*controlv1.SyncResults, error) {
	strategy, err := getStrategy(manifest.GetItems())
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	s.handlerMu.Lock()
	defer s.handlerMu.Unlock()
	handler, ok := s.updateHandlers[strategy]
	if !ok {
		return nil, status.Errorf(codes.Unimplemented, "no handler for update strategy: %q", strategy)
	}

	patchList, desired, err := handler.CalculateUpdate(ctx, manifest)
	if err != nil {
		return nil, err
	}

	return &controlv1.SyncResults{
		RequiredPatches: patchList,
		DesiredState:    desired,
	}, nil
}

func (s *UpdateServer) StreamServerInterceptor() grpc.StreamServerInterceptor {
	return func(srv interface{}, stream grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
		md, ok := metadata.FromIncomingContext(stream.Context())
		if !ok {
			return handler(srv, stream)
		}

		strategy := md.Get(controlv1.UpdateStrategyKey)
		if len(strategy) == 0 {
			return handler(srv, stream)
		}

		interceptors := make([]grpc.StreamServerInterceptor, 0, len(strategy))
		s.handlerMu.Lock()
		defer s.handlerMu.Unlock()
		for _, strat := range strategy {
			updateHandler, ok := s.updateHandlers[strat]
			if !ok {
				continue
			}

			if interceptor, ok := updateHandler.(UpdateStreamInterceptor); ok {
				interceptors = append(interceptors, interceptor.StreamServerInterceptor())
			}
		}

		switch len(interceptors) {
		case 0:
			return handler(srv, stream)
		case 1:
			return interceptors[0](srv, stream, info, handler)
		default:
			return chainStreamInterceptors(interceptors)(srv, stream, info, handler)
		}
	}
}

func chainStreamInterceptors(interceptors []grpc.StreamServerInterceptor) grpc.StreamServerInterceptor {
	return func(srv interface{}, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
		return interceptors[0](srv, ss, info, getChainStreamHandler(interceptors, 0, info, handler))
	}
}

func getChainStreamHandler(
	interceptors []grpc.StreamServerInterceptor,
	curr int,
	info *grpc.StreamServerInfo,
	finalHandler grpc.StreamHandler,
) grpc.StreamHandler {
	if curr == len(interceptors)-1 {
		return finalHandler
	}
	return func(srv interface{}, stream grpc.ServerStream) error {
		return interceptors[curr+1](srv, stream, info, getChainStreamHandler(interceptors, curr+1, info, finalHandler))
	}
}

func (s *UpdateServer) Collectors() []prometheus.Collector {
	var collectors []prometheus.Collector
	s.handlerMu.Lock()
	defer s.handlerMu.Unlock()
	for _, handler := range s.updateHandlers {
		collectors = append(collectors, handler.Collectors()...)
	}
	return collectors
}
