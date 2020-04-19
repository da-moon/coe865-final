(c *Client) EstimateCost(req *model.UpdateRequest) (*model.UpdateResponse, error)
(c *Client) KeyExchange(req *model.KeyExchangeRequest) (*model.KeyExchangeResponse, error)
(s *Server) EstimateCost(_req *model.UpdateRequest, _resp *model.UpdateResponse) error
(s *Server) KeyExchange(_req *model.KeyExchangeRequest, _resp *model.KeyExchangeResponse) error
(p *Plugin) Client(b *plugin.MuxBroker, c *rpc.Client) (interface{}, error)
(p *Plugin) Server(*plugin.MuxBroker) (interface{}, error)
(c *Client) EstimateCost(req *model.UpdateRequest) (*model.UpdateResponse, error)
(c *Client) KeyExchange(req *model.KeyExchangeRequest) (*model.KeyExchangeResponse, error)
(s *Server) EstimateCost(ctx context.Context, _req *model.UpdateRequest) (*model.UpdateResponse, error)
(s *Server) KeyExchange(ctx context.Context, _req *model.KeyExchangeRequest) (*model.KeyExchangeResponse, error)
(p *Plugin) GRPCClient(ctx context.Context, broker *plugin.GRPCBroker, c *grpcx.ClientConn) (interface{}, error)
(p *Plugin) GRPCServer(broker *plugin.GRPCBroker, s *grpcx.Server) error
Print(program string) string
Info() string
BuildContext() string
LevelFilter() *logutils.LevelFilter
ValidateLevelFilter(minLevel logutils.LogLevel, filter *logutils.LevelFilter) bool
NewLogWriter(buf int) *LogWriter
(l *LogWriter) RegisterHandler(lh LogHandler)
(l *LogWriter) DeregisterHandler(lh LogHandler)
(l *LogWriter) Write(p []byte) (n int, err error)
(w *GatedWriter) Flush()
(w *GatedWriter) Write(p []byte) (n int, err error)
UUID() string
ToBase64(s string) string
FromBase64(s string) string
ToHex(s string) string
FromHex(s string) string
(w *Writer) Write(p []byte) (n int, err error)
(w *Writer) MD5() []byte
(w *Writer) SHA256() []byte
(w *Writer) MD5HexString() string
(w *Writer) MD5Base64String() string
(w *Writer) SHA256HexString() string
(w *Writer) SHA256Base64String() string
(r *Reader) Read(p []byte) (n int, err error)
(r *Reader) MD5() []byte
(r *Reader) SHA256() []byte
(r *Reader) MD5HexString() string
(r *Reader) MD5Base64String() string
(r *Reader) SHA256HexString() string
(r *Reader) SHA256Base64String() string
(c *Config) SaveAsJSON(path string) error
(e ConfigExtension) String() string
(c *ConfigFactory) ReadConfigPaths(paths []string, extension ConfigExtension) (map[string]Config, error)
(d dirEnts) Len() int
(d dirEnts) Less(i, j int) bool
(d dirEnts) Swap(i, j int)
DefaultConfigFactory() *ConfigFactory
MergeFactory(a, b *ConfigFactory) *ConfigFactory
DecodeJSONConfig(r io.Reader) (*Config, error)
(c *ConfigFactory) DecodeRawConfig(r io.Reader) (*Config, error)
ExtractRouteControllerFromLine(input string) (*RouteController, error)
ExtractAutonomousSystemFromLine(input string) (*AutonomousSystem, error)
SanitizeAndSplitLine(input string) []string
trimComment(s string) string
SetTestLogger(t *testing.T)
SetDefaultLogger()
Debug(msg string)
Info(msg string)
Warn(msg string)
Error(msg string)
newDefaultLogger() *defaultLogger
newTestLogger(t *testing.T) *defaultLogger
(d *defaultLogger) addLogLevel(level, msg string) string
(d *defaultLogger) addExtraFields(extraFields map[string]interface{}, msg string) string
(d *defaultLogger) Debug(msg string, extraFields map[string]interface{})
(d *defaultLogger) Info(msg string, extraFields map[string]interface{})
(d *defaultLogger) Warn(msg string, extraFields map[string]interface{})
(d *defaultLogger) Error(msg string, extraFields map[string]interface{})
(tw testWriter) Write(p []byte) (n int, err error)
GenerateRPC2Routes(routes []JSON2) *mux.Router
GenerateRoutes(routes []Route) *mux.Router
JWT(secret string) func(next http.HandlerFunc) http.HandlerFunc
parse(secret, tokenString string) (jwt.MapClaims, error)
Cors(next http.HandlerFunc) http.HandlerFunc
Log(next http.HandlerFunc) http.HandlerFunc
LogErrorResponse(r *http.Request, err error, code int, message string)
WriteErrorJSON(w *http.ResponseWriter, r *http.Request, code int, message string)
WriteSuccessfulJSON(w *http.ResponseWriter, r *http.Request, data interface{})
LogSuccessfulResponse(r *http.Request, data interface{})
HealthCheck(w http.ResponseWriter, r *http.Request)
Preflight(w http.ResponseWriter, r *http.Request)
EncodeJSONWithoutErr(in interface{}) []byte
EncodeJSON(in interface{}) ([]byte, error)
DecodeJSON(data []byte, out interface{}) error
EncodeJSONWithIndentation(in interface{}) ([]byte, error)
EncodeJSONToWriter(w io.Writer, in interface{}, prefix, indent string) error
DecodeJSONFromReader(r io.Reader, out interface{}) error
(OverlayNetwork) EstimateCost(req *model.UpdateRequest) (*model.UpdateResponse, error)
(OverlayNetwork) KeyExchange(req *model.KeyExchangeRequest) (*model.KeyExchangeResponse, error)
getAPIListener(addr string) (net.Listener, error)
tcpAddress(ip string, port int) *net.TCPAddr
tcpAddressFromString(addr string) (*net.TCPAddr, error)
(ln tcpKeepAliveListener) Accept() (c net.Conn, err error)
main()
(c *ParseConfigCommand) Run(args []string) int
(c *ParseConfigCommand) Synopsis() string
(c *ParseConfigCommand) Help() string
(c *Command) readConfig() *config.Config
Create(conf *config.Config, logOutput io.Writer) (*Core, error)
(a *Core) Start() error
(a *Core) Shutdown() error
(a *Core) ShutdownCh() <-chan struct{}
(a *Core) EstimateCost() func()
(c *Command) Run(args []string) int
(c *Command) handleSignals(config *config.Config, core *Core) int
(c *Command) handleReload(config *config.Config, core *Core) *config.Config
(c *Command) Synopsis() string
(c *Command) Help() string
(c *Command) setupCore(config *config.Config, logOutput io.Writer) *Core
(c *Command) setupLoggers(config *config.Config) (*view.GatedWriter, *view.LogWriter, io.Writer)
(c *KeygenCommand) Run(_ []string) int
(c *KeygenCommand) Synopsis() string
(c *KeygenCommand) Help() string
(c *VersionCommand) Help() string
(c *VersionCommand) Run(_ []string) int
(c *VersionCommand) Synopsis() (s string)
CostEstimatorPathFlag(f *flag.FlagSet) *string
ConfigFilePathFlag(f *flag.FlagSet) *string
LogLevelFlag(f *flag.FlagSet) *string
DevFlag(f *flag.FlagSet) *bool
RPCPortFlag(f *flag.FlagSet) *int
CronFlag(f *flag.FlagSet) *string
(s *AppendSliceValue) String() string
(s *AppendSliceValue) Set(value string) error
init()
makeShutdownCh() <-chan struct{}
main()
