package dbx

var DB Database

type DSN struct {
	Host     string
	Port     int
	User     string
	Password string
	DBName   string
}
