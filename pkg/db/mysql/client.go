package mysql

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
	migrate "github.com/golang-migrate/migrate/v4"
	mysql_migrate "github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

type ClientMySQL struct {
	db *sql.DB
}

func NewMySQLClient() *ClientMySQL {
	db, err := sql.Open("mysql", "root:password@tcp(localhost:3307)/local-db")
	if err != nil {
		log.Fatal(err)
	}

	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}
	return &ClientMySQL{
		db: db,
	}
}

func (cli *ClientMySQL) Migrate() {
	if err := cli.db.Ping(); err != nil {
		log.Fatal(err)
	}
	driver, _ := mysql_migrate.WithInstance(cli.db, &mysql_migrate.Config{})
	m, err := migrate.NewWithDatabaseInstance(
		"file://pkg/db/migrations/mysql",
		"mysql",
		driver,
	)
	if err != nil {
		log.Fatal(err)
	}
	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatal(err)
	}
}
