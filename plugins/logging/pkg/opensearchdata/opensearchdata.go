package opensearchdata

import (
	"context"
	"errors"
	"os"
	"time"

	"github.com/cenkalti/backoff"
	backoffv2 "github.com/lestrrat-go/backoff/v2"
	"github.com/nats-io/nats.go"
	loggingutil "github.com/rancher/opni/plugins/logging/pkg/util"
	"go.uber.org/zap"
)

const (
	pendingValue = "job pending"
)

type DeleteStatus int

const (
	DeletePending DeleteStatus = iota
	DeleteRunning
	DeleteFinished
	DeleteFinishedWithErrors
	DeleteError
)

type ClusterStatus int

const (
	ClusterStatusGreen = iota
	ClusterStatusYellow
	ClusterStatusRed
	ClusterStatusError
)

type Manager struct {
	*loggingutil.AsyncOpensearchClient

	kv     *loggingutil.AsyncJetStreamClient
	logger *zap.SugaredLogger
}

func NewManager(logger *zap.SugaredLogger) *Manager {
	return &Manager{
		kv:                    loggingutil.NewAsyncJetStreamClient(),
		AsyncOpensearchClient: loggingutil.NewAsyncOpensearchClient(),
		logger:                logger,
	}
}

func (m *Manager) keyExists(keyToCheck string) (bool, error) {
	keys, err := m.kv.Keys()
	if err != nil {
		if errors.Is(err, nats.ErrNoKeysFound) {
			return false, nil
		}
		return false, err
	}
	for _, key := range keys {
		if key == keyToCheck {
			return true, nil
		}
	}
	return false, nil
}

func (m *Manager) newNatsConnection() (*nats.Conn, error) {
	natsURL := os.Getenv("NATS_SERVER_URL")
	natsSeedPath := os.Getenv("NKEY_SEED_FILENAME")

	opt, err := nats.NkeyOptionFromSeed(natsSeedPath)
	if err != nil {
		return nil, err
	}

	retryBackoff := backoff.NewExponentialBackOff()
	return nats.Connect(
		natsURL,
		opt,
		nats.MaxReconnects(-1),
		nats.CustomReconnectDelay(
			func(i int) time.Duration {
				if i == 1 {
					retryBackoff.Reset()
				}
				return retryBackoff.NextBackOff()
			},
		),
		nats.DisconnectErrHandler(
			func(nc *nats.Conn, err error) {
				m.logger.With(
					"err", err,
				).Warn("nats disconnected")
			},
		),
	)
}

func (m *Manager) setJetStream() nats.KeyValue {
	var (
		nc  *nats.Conn
		err error
	)

	retrier := backoffv2.Exponential(
		backoffv2.WithMaxRetries(0),
		backoffv2.WithMinInterval(5*time.Second),
		backoffv2.WithMaxInterval(1*time.Minute),
		backoffv2.WithMultiplier(1.1),
	)
	b := retrier.Start(context.TODO())
	for backoffv2.Continue(b) {
		nc, err = m.newNatsConnection()
		if err == nil {
			break
		}
		m.logger.Error("failed to connect to nats, retrying")
	}

	mgr, err := nc.JetStream()
	if err != nil {
		panic(err)
	}

	kv, err := mgr.CreateKeyValue(&nats.KeyValueConfig{
		Bucket:      "pending-delete",
		Description: "track pending deletes",
	})
	if err != nil {
		panic(err)
	}

	return kv
}