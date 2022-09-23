package db

import (
	"database/sql"
	"log"
	"os"
	"testing"
	_ "github.com/lib/pq"
)

const (
	driverName     = "postgres"
	dataSourceName = "postgres://root:secret@localhost:5432/simple_bank?sslmode=disable"
)

var testQueries *Queries
var testDBConn *sql.DB
func TestMain(m *testing.M) {
	var err error
	testDBConn, err = sql.Open(driverName, dataSourceName)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}
	testQueries = New(testDBConn)
	os.Exit(m.Run())
}
