package configs

var (
	ServerSetting   *Server
	JWTSetting      *JWTConfig
	DatabaseSetting *Database
)

type Server struct {
	RunMode  string
	HttpPort string
}

type JWTConfig struct {
	Secret string
	Issuer string
	Expire int
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
