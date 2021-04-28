package configs

var (
	ServerSetting   *Server
	DatabaseSetting *Database
)

type Server struct {
	RunMode   string
	HttpPort  string
	JWTSecret string
	JWTExpire int // 启动时 x time.Second
	JWTIssuer string
}

type Database struct {
	Username,
	Password,
	Host,
	DBName,
	Charset string
	ParseTime bool
}

func (s *Setting) ReadSection(k string, v interface{}) error {
	if err := s.vp.UnmarshalKey(k, v); err != nil {
		return err
	}
	return nil
}
