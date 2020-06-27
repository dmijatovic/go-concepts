package pgdb

import (
	"log"
	"os"

	// "github.com/jmoiron/sqlx"
	"database/sql"
	// postgres specific driver
	_ "github.com/lib/pq"
)


// Options for postgres connection
type Options struct {
	host     string
	port     int
	user     string
	password string
	dbname   string
	sslmode  string
}

// Connect to local postgres database
func Connect(config Options) (*sql.DB) {
	//"host=localhost port=5432 user=postgres password=changeme dbname=auth_db"
	cnnStr := fmt.Sprintf("host=%localhost port=%port user=postgres password=changeme dbname=auth_db"
	db, err := sql.Open("postgres", cnnStr)
	if err != nil {
		log.Printf("Failed to connect to db\n")
		os.Exit(401)
	}
	log.Printf("Connected to db\n")
	return db
}

// Close postgres database connection for passed pointer
func Close(db *sql.DB) {
	closeErr := db.Close()
	if closeErr != nil {
		log.Printf("Failed to close db connection %v\n", closeErr)
	}
	log.Println("Closing db connection")
}
