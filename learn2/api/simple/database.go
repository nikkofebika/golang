package simple

// type Database struct {
// 	Name string
// }

// func NewDatabaseMysql() *Database {
// 	return &Database{Name: "mysql"}
// }

// func NewDatabasePostgres() *Database {
// 	return &Database{Name: "postgres"}
// }

// type DatabaseRepository struct {
// 	Mysql    *Database
// 	Postgres *Database
// }

// func NewDatabaseRepository(mysql *Database, postgres *Database) *DatabaseRepository {
// 	return &DatabaseRepository{
// 		Mysql:    mysql,
// 		Postgres: postgres,
// 	}
// }

type Database struct {
	Name string
}

type DatabaseMysql Database
type DatabasePostgres Database

func NewDatabaseMysql() *DatabaseMysql {
	return (*DatabaseMysql)(&Database{Name: "mysql"})
}

func NewDatabasePostgres() *DatabasePostgres {
	return (*DatabasePostgres)(&Database{Name: "postgres"})
}

type DatabaseRepository struct {
	Mysql    *DatabaseMysql
	Postgres *DatabasePostgres
}

func NewDatabaseRepository(mysql *DatabaseMysql, postgres *DatabasePostgres) *DatabaseRepository {
	return &DatabaseRepository{
		Mysql:    mysql,
		Postgres: postgres,
	}
}
