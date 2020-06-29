package main

import (
	"log"
	"net/http"

	"./pgdb"
	"./routes"
)

func main() {
	host := "localhost:3000"
	log.Println("Starting...", host)
	// create new router
	mux := routes.Register()

	// connect to postgres database
	cnnStr := pgdb.ConnectionStr(pgdb.Settings{
		Host:     "localhost",
		Port:     5432,
		User:     "postgres",
		Password: "changeme",
		Dbname:   "auth_db",
	})
	// println(cnnStr)
	db := pgdb.Connect(cnnStr)
	//close connection at the end
	defer db.Close()

	//start server and register router
	log.Fatal(http.ListenAndServe(host, mux))
}
