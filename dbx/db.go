package dbx

var DB *DBX

type DSN struct {
	Host     string
	Port     int
	User     string
	Password string
	DBName   string
}
