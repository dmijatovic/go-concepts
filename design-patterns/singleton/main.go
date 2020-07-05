package main

import (
	"log"
	"sync"
)

type populationDB struct {
	cities map[string]int
}

var once sync.Once
var singleton *populationDB

func (db *populationDB) GetPopulation(name string) int {
	return db.cities[name]
}

func loadDataFromDB() (populationDB, error) {
	c := map[string]int{
		"A": 2345,
		"B": 123123,
		"C": 23423,
	}
	return populationDB{
		cities: c,
	}, nil
}

// GetPopulationDB returns data
func GetPopulationDB() *populationDB {
	//this function will called only once
	once.Do(func() {
		log.Println("Init DB...ONCE")
		db, err := loadDataFromDB()
		if err != nil {
			log.Panic(err)
		}
		singleton = &db
	})
	// we will return instance (singleton)
	log.Println("Return db singleton")
	return singleton
}

func main() {
	log.Println("Singleton example")
	db := GetPopulationDB()
	log.Println("db1", db)
	db2 := GetPopulationDB()
	log.Println("db2", db2)
	db3 := GetPopulationDB()
	log.Println("db3", db3)
}
