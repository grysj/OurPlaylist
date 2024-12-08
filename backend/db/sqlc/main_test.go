package db

import (
	"database/sql"
	"log"
	"os"
	"ourplaylist/util"
	"testing"

	_ "github.com/lib/pq"
)

var testQueries *Queries

// const (
// 	dbDriver = "postgres"
// 	dbSource = "postgresql://root:secret@localhost/ourplaylist_db?sslmode=disable"
// )

func TestMain(m *testing.M) {

	config, err := util.LoadConfig("../..")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}
	conn, err := sql.Open(config.DBDriver, config.DBSource)

	if err != nil {
		log.Fatal("cannot connect to  db:", err)
	}

	testQueries = New(conn)

	os.Exit(m.Run())

}
