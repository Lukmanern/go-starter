package database

import (
	"context"
	"database/sql"
	"log"
	"sync"

	_ "github.com/go-sql-driver/mysql"
)

type DB interface {
	Query(query string, args ...interface{}) (*sql.Rows, error)
	QueryContext(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error)
	Prepare(query string) (*sql.Stmt, error)
	PrepareContext(ctx context.Context, query string) (*sql.Stmt, error)
	Exec(query string, args ...interface{}) (sql.Result, error)
	ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error)
	QueryRowContext(ctx context.Context, query string, args ...interface{}) *sql.Row
}

// TODO: Refactor
// ADD: more function to check if table is exist or not

var (
	database     *sql.DB
	databaseOnce sync.Once
)

func LoadSQLDatabase() *sql.DB {
	databaseOnce.Do(func() {
		var err error
		database, err = sql.Open("mysql", "config.Configuration().GetdatabaseURI()")
		if err != nil {
			log.Panicf("cannot connect to MySQL %s\n", err)
		}
		err = database.Ping()
		if err != nil {
			log.Panicf("cannot ping to MySQL %s\n", err)
		}
	})

	return database
}
