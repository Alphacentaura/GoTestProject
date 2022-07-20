package driver

import (
	"database/sql"
	"github.com/lib/pq"
	"log"
	"os"
)

var db *sql.DB

func LogFatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func ConnectDB() *sql.DB {
	pgUrl, err := pq.ParseURL(os.Getenv("ELEPHANT_SQL_URL"))
	LogFatal(err)

	db, err = sql.Open("postgres", pgUrl)
	LogFatal(err)

	err = db.Ping()
	LogFatal(err)

	return db
}
