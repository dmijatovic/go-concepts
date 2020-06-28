package pgdb

import (
	"database/sql"
	"log"
	"os"

	// postgres specific driver
	_ "github.com/lib/pq"
)

var sqlDB *sql.DB

// Options for postgres connection
type Options struct {
	host     string
	port     int
	user     string
	password string
	dbname   string
	sslmode  string
}

func failedToConnect(err error) {
	log.Printf("Failed to connect to db. %v\n", err)
	os.Exit(401)
}

// Connect to local postgres database
func Connect() *sql.DB {
	//"host=localhost port=5432 user=postgres password=changeme dbname=auth_db"
	cnnStr := "host=localhost port=5432 user=postgres password=changeme dbname=auth_db sslmode=disable"
	// open connection
	db, err := sql.Open("postgres", cnnStr)
	if err != nil {
		failedToConnect(err)
	}
	// ping database
	err = db.Ping()
	if err != nil {
		failedToConnect(err)
	}

	log.Printf("Connected to db\n")
	sqlDB = db
	return db
}

// Close postgres database connection for passed pointer
func Close(db *sql.DB) {
	if db == nil && sqlDB != nil {
		db = sqlDB
	}
	closeErr := db.Close()
	if closeErr != nil {
		log.Printf("Failed to close db connection %v\n", closeErr)
	}
	log.Println("Closing db connection")
}

// AllUsers returns slice of users.
// db pointer is optional.
func AllUsers(db *sql.DB) ([]User, error) {
	if db == nil && sqlDB != nil {
		db = sqlDB
	}
	rs, err := db.Query("SELECT * FROM users;")
	if err != nil {
		log.Printf("pgdb.AllUsers query failed: %v", err)
		return nil, err
	}

	var users []User

	for rs.Next() {
		u := User{}
		rs.Scan(&u.ID, &u.Roles, &u.FirstName,
			&u.LastName, &u.Email, &u.password,
			&u.BirthDate, &u.CreateDate)

		users = append(users, u)
		// comma ok idiom -> check for error and if
		// not nil panic
		if err = rs.Err(); err != nil {
			panic(err)
		}
		// log.Printf("User...%v", u.Email)
	}

	return users, nil
}
