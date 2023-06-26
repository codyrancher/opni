package cli

import (
	"bytes"
	"fmt"
	"io"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/iancoleman/strcase"
	"github.com/rancher/opni/pkg/util/flagutil"
	"google.golang.org/genproto/googleapis/api/annotations"
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/descriptorpb"
)

func NewGenerator() *Generator {
	return &Generator{
		allFlagSets:    make(map[string]*flagSet),
		generatedFiles: make(map[string]*protogen.GeneratedFile),
	}
}

type Generator struct {
	plugin          *protogen.Plugin
	allFlagSets     map[string]*flagSet
	orderedFlagSets []*flagSet
	generatedFiles  map[string]*protogen.GeneratedFile
}

func (cg Generator) Name() string {
	return "cli"
}

func (cg *Generator) Generate(gen *protogen.Plugin) error {
	cg.plugin = gen
	for _, file := range gen.Files {
		if !file.Generate {
			continue
		}
		if !proto.HasExtension(file.Desc.Options(), E_Generator) {
			continue
		}
		ext := proto.GetExtension(file.Desc.Options(), E_Generator).(*GeneratorOptions)
		if !ext.GetGenerate() {
			continue
		}
		cg.generatedFiles[file.Desc.Path()] = cg.generateFile(gen, file)
	}
	return nil
}

// FileDescriptorProto.package field number
const fileDescriptorProtoPackageFieldNumber = 2

// FileDescriptorProto.syntax field number
const fileDescriptorProtoSyntaxFieldNumber = 12

var (
	_time      = protogen.GoImportPath("time")
	_fmt       = protogen.GoImportPath("fmt")
	_context   = protogen.GoImportPath("context")
	_io        = protogen.GoImportPath("io")
	_os        = protogen.GoImportPath("os")
	_strings   = protogen.GoImportPath("strings")
	_cli       = protogen.GoImportPath("github.com/rancher/opni/internal/codegen/cli")
	_proto     = protogen.GoImportPath("google.golang.org/protobuf/proto")
	_protojson = protogen.GoImportPath("google.golang.org/protobuf/encoding/protojson")
	_cobra     = protogen.GoImportPath("github.com/spf13/cobra")
	_pflag     = protogen.GoImportPath("github.com/spf13/pflag")
	_emptypb   = protogen.GoImportPath("google.golang.org/protobuf/types/known/emptypb")
	_flagutil  = protogen.GoImportPath("github.com/rancher/opni/pkg/util/flagutil")
	_enumflag  = protogen.GoImportPath("github.com/thediveo/enumflag/v2")
	_errors    = protogen.GoImportPath("errors")
)

func genLeadingComments(g *protogen.GeneratedFile, loc protoreflect.SourceLocation) {
	for _, s := range loc.LeadingDetachedComments {
		g.P(protogen.Comments(s))
		g.P()
	}
	if s := loc.LeadingComments; s != "" {
		g.P(protogen.Comments(s))
		g.P()
	}
}

func (cg *Generator) generateFile(gen *protogen.Plugin, file *protogen.File) *protogen.GeneratedFile {
	opts := GeneratorOptions{}
	applyOptions(file.Desc, &opts)
	if !opts.GetGenerate() {
		return nil
	}

	filename := file.GeneratedFilenamePrefix + "_cli.pb.go"
	g := gen.NewGeneratedFile(filename, file.GoImportPath)
	// Attach all comments associated with the syntax field.
	genLeadingComments(g, file.Desc.SourceLocations().ByPath(protoreflect.SourcePath{fileDescriptorProtoSyntaxFieldNumber}))
	g.P("// Code generated by cli_gen.go DO NOT EDIT.")
	g.P("// source: ", file.Desc.Path())

	g.P()
	// Attach all comments associated with the package field.
	genLeadingComments(g, file.Desc.SourceLocations().ByPath(protoreflect.SourcePath{fileDescriptorProtoPackageFieldNumber}))
	g.P("package ", file.GoPackageName)
	g.P()

	anyServices := len(file.Services) > 0
	cg.generateServices(&opts, file, g)
	if !anyServices {
		g.Skip()
	}

	for _, fs := range cg.orderedFlagSets {
		if !fs.wrote {
			fs.wrote = true
			fs.buf.g.P()
			fs.buf.g.Unskip()
			io.Copy(fs.buf.g, &fs.buf.buf)
		}
	}

	if opts.GenerateDeepcopy {
		cg.generateDeepcopyFunctions(file, g)
	}
	return g
}

func (cg *Generator) generateServices(opts *GeneratorOptions, file *protogen.File, g *protogen.GeneratedFile) {
	svcCtx := serviceGenWriters{}
	switch opts.ClientDependencyInjection {
	case ClientDependencyInjectionStrategy_InjectIntoContext:
		generateContextInjectionFunctions(file, g)
		svcCtx = serviceGenWriters{
			PrintAddCmd: func(method *protogen.Method, g *protogen.GeneratedFile) {
				g.P("cmd.AddCommand(Build", method.GoName, "Cmd())")
			},
			PrintCmdBuilderSignature: func(methodName, _ string, g *protogen.GeneratedFile) {
				g.P("func Build", methodName, "Cmd() *", _cobra.Ident("Command"), " {")
			},
			PrintObtainClient: func(service *protogen.Service, g *protogen.GeneratedFile) {
				g.P("client, ok := ", service.GoName+"ClientFromContext(cmd.Context())")
				g.P("if !ok {")
				g.P(" cmd.PrintErrln(\"failed to get client from context\")")
				g.P(" return nil")
				g.P("}")
			},
		}

	case ClientDependencyInjectionStrategy_InjectAsArgument:
		svcCtx = serviceGenWriters{
			PrintAddCmd: func(method *protogen.Method, g *protogen.GeneratedFile) {
				g.P("cmd.AddCommand(Build", method.GoName, "Cmd(client))")
			},
			PrintCmdBuilderSignature: func(methodName, svcName string, g *protogen.GeneratedFile) {
				g.P("func Build", methodName, "Cmd(client ", svcName+"Client", ") *", _cobra.Ident("Command"), " {")
			},
			PrintObtainClient: func(service *protogen.Service, g *protogen.GeneratedFile) {},
		}
	}
	g.P()
	for _, service := range file.Services {
		cg.generateServiceTopLevelCmd(service, g, svcCtx)
		g.P()
		for _, method := range service.Methods {
			cg.generateMethodCmd(service, method, g, svcCtx)
			g.P()
		}
	}
}

type serviceGenWriters struct {
	PrintAddCmd              func(method *protogen.Method, g *protogen.GeneratedFile)
	PrintCmdBuilderSignature func(methodName, svcName string, g *protogen.GeneratedFile)
	PrintObtainClient        func(service *protogen.Service, g *protogen.GeneratedFile)
}

func generateContextInjectionFunctions(file *protogen.File, g *protogen.GeneratedFile) {
	for _, service := range file.Services {
		g.P("type contextKey_", service.GoName, "_type struct{}")
		g.P("var contextKey_", service.GoName, " contextKey_", service.GoName, "_type")
		g.P()
		g.P("func ContextWith", service.GoName, "Client(ctx ", _context.Ident("Context"), ", client ", service.GoName, "Client) context.Context {")
		g.P(" return context.WithValue(ctx, contextKey_", service.GoName, ", client)")
		g.P("}")
		g.P()
		g.P("func ", service.GoName, "ClientFromContext(ctx ", _context.Ident("Context"), ") (", service.GoName, "Client, bool) {")
		g.P(" client, ok := ctx.Value(contextKey_", service.GoName, ").(", service.GoName, "Client)")
		g.P(" return client, ok")
		g.P("}")
		g.P()
	}
}

func (cg *Generator) generateServiceTopLevelCmd(service *protogen.Service, g *protogen.GeneratedFile, writers serviceGenWriters) {
	leadingComments := formatComments(service.Comments)

	opts := CommandGroupOptions{
		Use: strcase.ToKebab(service.GoName),
	}
	applyOptions(service.Desc, &opts)

	g.P("var extraCmds_", service.GoName, " []*", _cobra.Ident("Command"))
	g.P()
	g.P("func addCortexOpsCommand(custom *", _cobra.Ident("Command"), ") {")
	g.P(" extraCmds_", service.GoName, " = append(extraCmds_", service.GoName, ", custom)")
	g.P("}")
	g.P()

	writers.PrintCmdBuilderSignature(service.GoName, service.GoName, g)

	g.P("cmd := &", _cobra.Ident("Command"), "{")
	g.P(" Use:   \"", opts.Use, "\",")
	if len(leadingComments) > 0 {
		g.P(" Short: `", leadingComments[0], "`,")
	}
	if len(leadingComments) > 1 {
		g.P(" Long: `")
		for _, c := range leadingComments[1:] {
			g.P(c)
		}
		g.P("`[1:],")
	}
	if opts.GroupId != "" {
		g.P(" GroupID: \"", opts.GroupId, "\",")
	}

	g.P(" Args: cobra.NoArgs,")
	g.P(" ValidArgsFunction: cobra.NoFileCompletions,")

	g.P("}")
	g.P()

	for _, method := range service.Methods {
		writers.PrintAddCmd(method, g)
	}
	g.P("for _, extraCmd := range extraCmds_", service.GoName, " {")
	g.P(" cmd.AddCommand(extraCmd)")
	g.P("}")

	g.P(_cli.Ident("AddOutputFlag(cmd)"))

	g.P("return cmd")
	g.P("}")
}

func (cg *Generator) generateMethodCmd(service *protogen.Service, method *protogen.Method, g *protogen.GeneratedFile, writers serviceGenWriters) {
	writers.PrintCmdBuilderSignature(method.GoName, service.GoName, g)
	isEmpty := method.Input.Desc.FullName() == "google.protobuf.Empty"
	if !isEmpty {
		g.P("in := &", g.QualifiedGoIdent(method.Input.GoIdent), "{}")
	}
	methodOptions := method.Desc.Options().(*descriptorpb.MethodOptions)
	leadingComments := formatComments(method.Comments)

	opts := CommandOptions{
		Use: strcase.ToKebab(method.GoName),
	}
	applyOptions(method.Desc, &opts)

	g.P("cmd := &", _cobra.Ident("Command"), "{")
	g.P(" Use:   \"", opts.Use, "\",")
	if len(leadingComments) > 0 {
		g.P(" Short: ", fmt.Sprintf("%q", leadingComments[0]), ",")
	}
	httpExt, hasHttpExt := getExtension[*annotations.HttpRule](method.Desc, annotations.E_Http)
	if len(leadingComments) > 1 || hasHttpExt {
		g.P(" Long: `")
		for _, c := range leadingComments[1:] {
			g.P(c)
		}
		if hasHttpExt {
			if len(leadingComments) > 1 {
				g.P()
			}
			g.P("HTTP handlers for this method:")
			bindings := append([]*annotations.HttpRule{httpExt}, httpExt.GetAdditionalBindings()...)
			for _, httpExt := range bindings {
				var httpMethod, path string
				switch pattern := httpExt.Pattern.(type) {
				case *annotations.HttpRule_Get:
					httpMethod = "GET"
					path = pattern.Get
				case *annotations.HttpRule_Post:
					httpMethod = "POST"
					path = pattern.Post
				case *annotations.HttpRule_Put:
					httpMethod = "PUT"
					path = pattern.Put
				case *annotations.HttpRule_Delete:
					httpMethod = "DELETE"
					path = pattern.Delete
				case *annotations.HttpRule_Patch:
					httpMethod = "PATCH"
					path = pattern.Patch
				case *annotations.HttpRule_Custom:
					httpMethod = strings.ToUpper(httpExt.GetCustom().Kind)
					path = httpExt.GetCustom().Path
				}
				g.P(fmt.Sprintf("- %s %s", httpMethod, path))
			}
		}
		g.P("`[1:],")
	}
	if methodOptions.GetDeprecated() {
		g.P(" Deprecated: \"", method.GoName, " is deprecated.\",")
	}

	g.P(" Args: cobra.NoArgs,")
	g.P(" ValidArgsFunction: cobra.NoFileCompletions,")

	cg.generateRun(service, method, g, writers)

	g.P("}")

	// Generate flags for input fields recursively
	if !isEmpty {
		flagSet := cg.generateFlagSet(g, method.Input)
		g.P("cmd.Flags().AddFlagSet(in.FlagSet())")

		cg.generateFlagCompletionFuncs(g, flagSet)
	}

	for _, requiredFlag := range opts.RequiredFlags {
		g.P("cmd.MarkFlagRequired(\"", requiredFlag, "\")")
	}

	g.P("return cmd")
	g.P("}")
}

type fieldTypeDefaults struct {
	defaultValue  string
	flagsFunction string
}

var defaultsForType = map[protoreflect.Kind]fieldTypeDefaults{
	protoreflect.BoolKind:     {"false", "BoolVar"},
	protoreflect.Int32Kind:    {"0", "Int32Var"},
	protoreflect.Sint32Kind:   {"0", "Int32Var"},
	protoreflect.Sfixed32Kind: {"0", "Int32Var"},
	protoreflect.Uint32Kind:   {"0", "Uint32Var"},
	protoreflect.Fixed32Kind:  {"0", "Uint32Var"},
	protoreflect.Int64Kind:    {"0", "Int64Var"},
	protoreflect.Sint64Kind:   {"0", "Int64Var"},
	protoreflect.Sfixed64Kind: {"0", "Int64Var"},
	protoreflect.Uint64Kind:   {"0", "Uint64Var"},
	protoreflect.Fixed64Kind:  {"0", "Uint64Var"},
	protoreflect.FloatKind:    {"0.0", "Float32Var"},
	protoreflect.DoubleKind:   {"0.0", "Float64Var"},
	protoreflect.StringKind:   {`""`, "StringVar"},
	protoreflect.BytesKind:    {"nil", "BytesHexVar"},
}

var stringToMapValueTypeDefaults = map[protoreflect.Kind]fieldTypeDefaults{
	protoreflect.StringKind:   {"nil", "StringToStringVar"},
	protoreflect.Int32Kind:    {"nil", "StringToIntVar"},
	protoreflect.Sint32Kind:   {"nil", "StringToIntVar"},
	protoreflect.Sfixed32Kind: {"nil", "StringToIntVar"},
	protoreflect.Uint32Kind:   {"nil", "StringToIntVar"},
	protoreflect.Fixed32Kind:  {"nil", "StringToIntVar"},
	protoreflect.Int64Kind:    {"nil", "StringToInt64Var"},
	protoreflect.Sint64Kind:   {"nil", "StringToInt64Var"},
	protoreflect.Sfixed64Kind: {"nil", "StringToInt64Var"},
	protoreflect.Uint64Kind:   {"nil", "StringToInt64Var"},
	protoreflect.Fixed64Kind:  {"nil", "StringToInt64Var"},
}

var typeHints = map[string]string{
	"ipNet": "IPNetValue", // type:flagutil.F
}

var specialCaseReplacements = map[string]string{
	"s-3": "s3",
}

type flagCompletion struct {
	name string
	fn   string
}

type flagSet struct {
	receiver *protogen.Message
	buf      *buffer
	deps     map[string]*flagSet
	wrote    bool

	secretFields         []*protogen.Field
	depsWithSecretFields []*protogen.Field

	flagCompletionFuncs []flagCompletion
}

type buffer struct {
	g   *protogen.GeneratedFile
	buf bytes.Buffer
}

func (b *buffer) P(v ...interface{}) {
	// note that this logic is different than the standard P method, it supports
	// recursive expansion of sub-slices
	var loop func(v ...interface{})
	loop = func(v ...interface{}) {
		for _, x := range v {
			switch x := x.(type) {
			case protogen.GoIdent:
				fmt.Fprint(&b.buf, b.g.QualifiedGoIdent(x))
			case []any:
				loop(x...)
			default:
				fmt.Fprint(&b.buf, x)
			}
		}
	}
	loop(v...)
	fmt.Fprintln(&b.buf)
}

func (b *buffer) QualifiedGoIdent(ident protogen.GoIdent) string {
	return b.g.QualifiedGoIdent(ident)
}

func (cg *Generator) generateFlagSet(g *protogen.GeneratedFile, message *protogen.Message) *flagSet {
	if existing, ok := cg.allFlagSets[g.QualifiedGoIdent(message.GoIdent)]; ok {
		return existing
	}

	deps := make(map[string]*flagSet)
	fs := &flagSet{
		receiver: message,
		deps:     deps,
	}

	cg.allFlagSets[g.QualifiedGoIdent(message.GoIdent)] = fs
	cg.orderedFlagSets = append(cg.orderedFlagSets, fs)

	if dest, ok := cg.generatedFiles[message.Desc.ParentFile().Path()]; ok {
		fs.buf = &buffer{g: dest}
	} else {
		fs.buf = &buffer{g: g}
	}
	{
		g := fs.buf
		g.P("func (in *", message.GoIdent.GoName, ") FlagSet(prefix ...string) *", _pflag.Ident("FlagSet"), " {")
		g.P("fs := ", _pflag.Ident("NewFlagSet("), "\""+message.GoIdent.GoName+"\"", ", ", _pflag.Ident("ExitOnError"), ")")
		g.P("fs.SortFlags = true")

		for _, field := range message.Fields {
			// Skip google.protobuf.Empty message
			if field.Message != nil && isEmptypb(field.Message.Desc) {
				continue
			}

			kebabName := formatKebab(field.GoName)
			commentLines := formatComments(field.Comments)
			comment := strings.Join(commentLines, " ")

			var hasCustomDefault bool
			var defaultValue any

			flagOpts := FlagOptions{}
			applyOptions(field.Desc, &flagOpts)
			if flagOpts.Skip {
				continue
			}

			if flagOpts.Default != "" {
				hasCustomDefault = true
				if field.Desc.Kind() == protoreflect.StringKind {
					if field.Desc.IsList() {
						defaultValue = unparseStringSlice(flagOpts.Default)
					} else {
						defaultValue = fmt.Sprintf("%q", strings.Trim(flagOpts.Default, `"`))
					}
				} else {
					defaultValue = flagOpts.Default
				}
			} else if flagOpts.Env != "" {
				hasCustomDefault = true
				defaultValue = []any{_os.Ident("Getenv"), fmt.Sprintf("(%q)", flagOpts.Env)}
				comment += fmt.Sprintf(" ($%s)", flagOpts.Env)
			}

			if flagOpts.Secret {
				comment = fmt.Sprintf("\033[31m[secret]\033[0m %s", comment)
				fs.secretFields = append(fs.secretFields, field)
			}

			if field.Desc.Options().(*descriptorpb.FieldOptions).GetDeprecated() {
				g.P(`fs.MarkDeprecated(`, _strings.Ident("Join"), `(append(prefix, "`, kebabName, `"), "."),`, fmt.Sprintf("%q", comment), `)`)
			}

			defaults, ok := defaultsForType[field.Desc.Kind()]
			if !ok && field.Desc.IsMap() && field.Desc.MapKey().Kind() == protoreflect.StringKind {
				// support for map[string]float64 which is missing from pflag
				if defaults, ok = stringToMapValueTypeDefaults[field.Desc.MapValue().Kind()]; !ok {
					switch field.Desc.MapValue().Kind() {
					case protoreflect.DoubleKind:
						if defaultValue == nil {
							defaultValue = "nil"
						}
						// todo: this might not work
						if str, ok := defaultValue.(string); ok && strings.HasPrefix(str, "{") {
							defaultValue = "map[string]float64" + str
						}
						g.P(`fs.Var(`, _flagutil.Ident("StringToFloat64Value"), `(`, defaultValue, `, &in.`, field.GoName, `), `, _strings.Ident("Join"), `(append(prefix, "`, kebabName, `"), "."),`, fmt.Sprintf("%q", comment), `)`)
					default:
						panic(fmt.Sprintf("unimplemented: map[string]%s", field.Desc.MapValue().Kind()))
					}
					continue
				}
			}

			if flagOpts.TypeOverride != "" {
				flagUtilValue, ok := typeHints[flagOpts.TypeOverride]
				if !ok {
					panic("unknown type override: " + flagOpts.TypeOverride)
				}
				if field.Desc.IsList() {
					flagUtilValue = strings.Replace(flagUtilValue, "Value", "SliceValue", 1)
				}
				if defaultValue == nil {
					if field.Desc.IsList() {
						defaultValue = "nil"
					} else {
						defaultValue = `""`
					}
				}
				g.P(`fs.Var(`, _flagutil.Ident(flagUtilValue), `(`, defaultValue, `, &in.`, field.GoName, `), `, _strings.Ident("Join"), `(append(prefix, "`, kebabName, `"), "."),`, fmt.Sprintf("%q", comment), `)`)
				continue
			}

			if ok {
				if field.Desc.HasPresence() {
					switch field.Desc.Kind() {
					case protoreflect.BoolKind:
						g.P(`fs.Var(`, _flagutil.Ident("BoolPtrValue"), `(&in.`, field.GoName, `), `, _strings.Ident("Join"), `(append(prefix, "`, kebabName, `"), "."),`, fmt.Sprintf("%q", comment), `)`)
						continue
					default:
						panic("unimplemented: *" + field.Desc.Kind().String())
					}
				}
				if !hasCustomDefault {
					defaultValue = defaults.defaultValue
				}
				flagsFunction := defaults.flagsFunction
				if field.Desc.IsList() {
					flagsFunction = strings.Replace(flagsFunction, "Var", "SliceVar", 1)
					if !hasCustomDefault {
						defaultValue = "nil"
					}
				}
				if defaultValue == nil {
					defaultValue = `""`
				}

				g.P(`fs.`, flagsFunction, `(&in.`, field.GoName, `, `, _strings.Ident("Join"), `(append(prefix, "`, kebabName, `"), "."),`, defaultValue, `,`, fmt.Sprintf("%q", comment), `)`)
				continue
			}

			switch field.Desc.Kind() {
			case protoreflect.EnumKind:
				g.P(`fs.Var(`, _enumflag.Ident("New"), `(&in.`, field.GoName, `, "`, field.Enum.Desc.Name(), `", map[`+g.QualifiedGoIdent(field.Enum.GoIdent)+`][]string{`)
				for _, v := range field.Enum.Values {
					g.P(g.QualifiedGoIdent(v.GoIdent), `: {`, fmt.Sprintf("%q", string(v.Desc.Name())), `},`)
				}
				g.P("},", _enumflag.Ident("EnumCaseSensitive"), `), `, _strings.Ident("Join"), `(append(prefix, "`, kebabName, `"), "."),`, fmt.Sprintf("%q", comment), `)`)
				var allValues []string
				for _, v := range field.Enum.Values {
					allValues = append(allValues, strconv.Quote(string(v.Desc.Name())))
				}
				fs.flagCompletionFuncs = append(fs.flagCompletionFuncs, flagCompletion{
					name: kebabName,
					fn:   `return []string{` + strings.Join(allValues, ", ") + `}, cobra.ShellCompDirectiveDefault`,
				})
				continue
			case protoreflect.MessageKind:
				if custom, ok := customFieldGenerators[string(field.Message.Desc.FullName())]; ok {
					custom(g, field, defaultValue, kebabName, fmt.Sprintf("%q", comment))
					continue
				}
				// Add flag sets for nested messages
				if !field.Desc.IsList() {
					g.P("if in.", field.GoName, " == nil {")
					g.P(" in.", field.GoName, " = &", field.Message.GoIdent, "{}")
					g.P("}")
					g.P("fs.AddFlagSet(in.", field.GoName, `.FlagSet(append(prefix,"`, kebabName, `")...))`)
					flagSetOpts := FlagSetOptions{}
					applyOptions(field.Desc, &flagSetOpts)

					flagSetOpts.ForEachDefault(field.Message, func(fd protoreflect.FieldDescriptor, v protoreflect.Value) bool {
						fdOpts := FlagOptions{}
						applyOptions(fd, &fdOpts)
						if fdOpts.Skip {
							return true
						}
						g.P(`fs.Lookup(`, _strings.Ident("Join"), `(append(prefix, "`, kebabName, `", "`, formatKebab(fd.Name()), `"), ".")).DefValue = `, fmt.Sprintf("%q", v.String()))
						return true
					})

					depFs := cg.generateFlagSet(g.g, field.Message)
					deps[kebabName] = depFs
					if len(depFs.secretFields) > 0 || len(depFs.depsWithSecretFields) > 0 {
						fs.depsWithSecretFields = append(fs.depsWithSecretFields, field)
					}
				}
				continue
			}
		}

		g.P("return fs")
		g.P("}")

		cg.genSecretMethods(g, fs)
	}

	return fs
}

func (cg *Generator) genSecretMethods(g *buffer, fs *flagSet) {
	if len(fs.secretFields) == 0 && len(fs.depsWithSecretFields) == 0 {
		return
	}

	g.P()
	g.P("func (in *", fs.receiver.GoIdent, ") RedactSecrets() {")
	g.P("if in == nil {")
	g.P(" return")
	g.P("}")
	for _, field := range fs.secretFields {
		g.P("if in.", field.GoName, " != \"\" {")
		g.P(" in.", field.GoName, " = \"***\"")
		g.P("}")
	}
	for _, dep := range fs.depsWithSecretFields {
		g.P("in.", dep.GoName, ".RedactSecrets()")
	}
	g.P("}")

	g.P()
	g.P("func (in *", fs.receiver.GoIdent, ") UnredactSecrets(unredacted *", fs.receiver.GoIdent, ") error {")
	g.P("if in == nil {")
	g.P(" return nil")
	g.P("}")
	for _, field := range fs.secretFields {
		g.P("if in.", field.GoName, " == \"***\" {")
		g.P(" if unredacted.Get", field.GoName, "() == \"\" {")
		g.P(`  return `, _errors.Ident("New"), `("cannot unredact: missing value for secret field: `, field.GoName, `")`)
		g.P(" }")
		g.P(" in.", field.GoName, " = unredacted.", field.GoName)
		g.P("}")
	}
	for _, dep := range fs.depsWithSecretFields {
		g.P("if err := in.", dep.GoName, ".UnredactSecrets(unredacted.Get", dep.GoName, "()); err != nil {")
		g.P(" return err")
		g.P("}")
	}
	g.P("return nil")
	g.P("}")
}

func (cg *Generator) generateDeepcopyFunctions(file *protogen.File, g *protogen.GeneratedFile) {
	for _, msg := range file.Messages {
		g.P()
		g.P("func (in *", msg.GoIdent, ") DeepCopyInto(out *", msg.GoIdent, ") {")
		g.P(_proto.Ident("Merge"), "(in, out)")
		g.P("}")
		g.P()
		g.P("func (in *", msg.GoIdent, ") DeepCopy() *", msg.GoIdent, " {")
		g.P(" return ", _proto.Ident("Clone"), "(in).(*", msg.GoIdent, ")")
		g.P("}")
	}
}

var customFieldGenerators = map[string]func(g *buffer, field *protogen.Field, defaultValue any, flagName, usage string){
	"google.protobuf.Duration": func(g *buffer, field *protogen.Field, defaultValue any, flagName, usage string) {
		var identName string
		if field.Desc.IsList() {
			identName = "DurationpbSliceValue"
		} else {
			identName = "DurationpbValue"
		}
		values := []any{"0"}
		switch defaultValue := defaultValue.(type) {
		case string:
			if field.Desc.IsList() {
				values = unparseDurationList(defaultValue)
			} else {
				values = unparseDuration(defaultValue)
			}
		case []any:
			values = defaultValue
		}
		g.P(`fs.Var(`, _flagutil.Ident(identName), `(`, values, `, &in.`, field.GoName, `), `, _strings.Ident("Join"), `(append(prefix, "`, flagName, `"), "."),`, usage, `)`)
	},
	"google.protobuf.Timestamp": func(g *buffer, field *protogen.Field, defaultValue any, flagName, usage string) {
		if defaultValue == "" {
			defaultValue = `"now"`
		}
		g.P(`fs.Var(`, _flagutil.Ident("TimestamppbValue"), `(`, defaultValue, `, &in.`, field.GoName, `), `, _strings.Ident("Join"), `(append(prefix, "`, flagName, `"), "."),`, usage, `)`)
	},
}

func (cg *Generator) generateFlagCompletionFuncs(g *protogen.GeneratedFile, fs *flagSet, prefix ...string) {
	for _, comp := range fs.flagCompletionFuncs {
		flagName := strings.Join(append(prefix, comp.name), ".")
		g.P(`cmd.RegisterFlagCompletionFunc("`, flagName, `", func(cmd *`, _cobra.Ident("Command"), `, args []string, toComplete string) ([]string, `, _cobra.Ident("ShellCompDirective"), `) {`)
		g.P(comp.fn)
		g.P(`})`)
	}

	orderedDeps := make([]string, 0, len(fs.deps))
	for k := range fs.deps {
		orderedDeps = append(orderedDeps, k)
	}
	sort.Strings(orderedDeps)
	for _, depName := range orderedDeps {
		dep := fs.deps[depName]
		cg.generateFlagCompletionFuncs(g, dep, append(prefix, depName)...)
	}
}

func formatKebab[T ~string](name T) string {
	kebabName := strcase.ToKebab(string(name))
	for orig, replacement := range specialCaseReplacements {
		if kebabName == orig {
			kebabName = replacement
		} else if strings.Contains(kebabName, orig+"-") {
			kebabName = strings.ReplaceAll(kebabName, orig+"-", replacement+"-")
		} else if strings.Contains(kebabName, "-"+orig) {
			kebabName = strings.ReplaceAll(kebabName, "-"+orig, "-"+replacement)
		}
	}
	return kebabName
}

// converts a duration string (e.g. "1h", "15m", 2h30m) into the equivalent go syntax
// string (e.g. "1*time.Hour", "15*time.Minute", "2*time.Hour+30*time.Minute")
func unparseDuration(durationStr string) (tokens []any) {
	if durationStr == "" {
		tokens = append(tokens, "0")
		return
	}

	duration, err := flagutil.ParseDurationWithExtendedUnits(durationStr)
	if err != nil {
		panic(fmt.Sprintf("invalid duration %q: %v", durationStr, err))
	}

	if duration == 0 {
		tokens = append(tokens, "0")
		return
	}

	h := duration / time.Hour
	m := (duration - (h * time.Hour)) / time.Minute
	s := (duration - (h * time.Hour) - (m * time.Minute)) / time.Second
	ms := (duration - (h * time.Hour) - (m * time.Minute) - (s * time.Second)) / time.Millisecond
	us := (duration - (h * time.Hour) - (m * time.Minute) - (s * time.Second) - (ms * time.Millisecond)) / time.Microsecond
	ns := (duration - (h * time.Hour) - (m * time.Minute) - (s * time.Second) - (ms * time.Millisecond) - (us * time.Microsecond)) / time.Nanosecond

	plus := []any{}
	if h != 0 {
		tokens = append(tokens, fmt.Sprintf("%d*", h), _time.Ident("Hour"))
		plus = []any{"+"}
	}
	if m != 0 {
		tokens = append(append(tokens, plus...), fmt.Sprintf("%d*", m), _time.Ident("Minute"))
		plus = []any{"+"}
	}
	if s != 0 {
		tokens = append(append(tokens, plus...), fmt.Sprintf("%d*", s), _time.Ident("Second"))
		plus = []any{"+"}
	}
	if ms != 0 {
		tokens = append(append(tokens, plus...), fmt.Sprintf("%d*", ms), _time.Ident("Millisecond"))
		plus = []any{"+"}
	}
	if us != 0 {
		tokens = append(append(tokens, plus...), fmt.Sprintf("%d*", us), _time.Ident("Microsecond"))
		plus = []any{"+"}
	}
	if ns != 0 {
		tokens = append(append(tokens, plus...), fmt.Sprintf("%d*", ns), _time.Ident("Nanosecond"))
	}
	return
}

func unparseDurationList(commaSeparatedDurations string) (tokens []any) {
	tokens = append(tokens, "[]", _time.Ident("Duration"), "{")
	for _, durationStr := range strings.Split(commaSeparatedDurations, ",") {
		tokens = append(tokens, unparseDuration(durationStr)...)
		tokens = append(tokens, ",")
	}
	return append(tokens, "}")
}

func unparseStringSlice(commaSeparatedStrings string) (tokens []any) {
	tokens = append(tokens, "[]string{")
	commaSeparatedStrings = strings.Trim(commaSeparatedStrings, "[]")
	for _, s := range strings.Split(commaSeparatedStrings, ",") {
		if s == "" {
			continue
		}
		tokens = append(tokens, strconv.Quote(s), ",")
	}
	return append(tokens, "}")
}

func formatComments(comments protogen.CommentSet) (leadingComments []string) {
	lines := strings.Split(strings.TrimSuffix(comments.Leading.String(), "\n"), "\n")
	for _, line := range lines {
		line := strings.TrimRight(strings.TrimLeft(line, " /"), " ")
		if strings.HasPrefix(line, "+") {
			continue // skip directives
		}
		leadingComments = append(leadingComments, line)
	}
	return
}

func (cg *Generator) generateRun(service *protogen.Service, method *protogen.Method, g *protogen.GeneratedFile, writers serviceGenWriters) {
	g.P(" RunE: func(cmd *", _cobra.Ident("Command"), ", args []string) error {")
	writers.PrintObtainClient(service, g)
	requestIsEmpty := isEmptypb(method.Desc.Input())
	responseIsEmpty := isEmptypb(method.Desc.Output())

	responseVarName := "_"
	if !responseIsEmpty {
		responseVarName = "response"
	}
	rpcCall := []any{}
	if requestIsEmpty {
		rpcCall = append(rpcCall, "(cmd.Context(), &", _emptypb.Ident("Empty"), "{})")
	} else {
		rpcCall = append(rpcCall, "(cmd.Context(), in)")
	}

	g.P(append([]any{responseVarName, ", err := client.", method.GoName}, rpcCall...)...)

	g.P("if err != nil {")
	g.P(" return err")
	g.P("}")

	if !responseIsEmpty {
		g.P(_cli.Ident("RenderOutput"), "(cmd, response)")
	}

	g.P("return nil")
	g.P("},")
}

func isEmptypb(t protoreflect.MessageDescriptor) bool {
	return t.FullName() == "google.protobuf.Empty"
}
