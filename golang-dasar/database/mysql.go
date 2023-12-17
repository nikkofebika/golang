package database

var connection string

func init() {
	connection = "MySikil"
}

func GetConnection() string {
	return connection
}
