package db

import (
	"log"
	"os"

	"github.com/go-pg/pg"
)

// Connect to local postgres database
func Connect() *pg.DB {
	opt := &pg.Options{
		Addr:     "localhost:5432",
		User:     "postgres",
		Password: "changeme",
		Database: "auth_db",
	}
	//create db pointer
	var db *pg.DB = pg.Connect(opt)
	if db == nil {
		log.Printf("Failed to connect to db\n")
		os.Exit(401)
	}
	log.Printf("Connected to db\n")
	return db
}

// Close postgres database connection for passed pointer
func Close(db *pg.DB) {
	closeErr := db.Close()
	if closeErr != nil {
		log.Printf("Failed to close db connection %v\n", closeErr)
	}
	log.Println("Closing db connection")
}
