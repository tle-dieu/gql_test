package mysql

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
	migrate "github.com/golang-migrate/migrate/v4"
	mysql_migrate "github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

type clientMySQL struct {
	db *sql.DB
}

func NewMySQLClient() *clientMySQL {
	db, err := sql.Open("mysql", "root:password@tcp(localhost:3306)/local-db")
	if err != nil {
		log.Fatal(err)
	}

	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}
	return &clientMySQL{
		db: db,
	}
}

func (cli *clientMySQL) Migrate() {
	if err := cli.db.Ping(); err != nil {
		log.Fatal(err)
	}
	driver, _ := mysql_migrate.WithInstance(cli.db, &mysql_migrate.Config{})
	m, _ := migrate.NewWithDatabaseInstance(
		"file://internal/pkg/db/migrations/mysql",
		"mysql",
		driver,
	)
	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatal(err)
	}
}
