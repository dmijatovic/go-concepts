package main

/*
Adapter patter means making a intermediate method that will connect two not compatible interfaces.
This is very simple example where we have one service providing us information that cannot be directly consumed by our internal interface. Then we make an adaptor to convert reveived info into a format our code can work with.
If this information is intensive we should implement caching. This example implements local caching using md5 hash algo to create unique map keys from received object

*/

import (
	"crypto/md5"
	"encoding/json"
	"log"
	"math/rand"
	"time"
)

// we cache converted celcsius objects
var tinyCache = map[[16]byte]tempInC{}

type tempInF struct {
	name    string
	value   int
	prefix  string
	postfix string
}

type tempInC struct {
	name  string
	value int
}

func getFahrenheitTemp(seed int64) tempInF {
	rand.Seed(seed)

	return tempInF{
		name:    "Fahrenheit",
		value:   rand.Intn(100),
		prefix:  "F",
		postfix: " fahrenheit",
	}

}

func getMd5Hash(obj interface{}) [16]byte {
	// convert object to json string
	bytes, err := json.Marshal(obj)
	if err != nil {
		log.Panic(err)
	}
	// hash json string using md5 algo
	return md5.Sum(bytes)
}

func getFromCache(f tempInF) (tempInC, bool) {
	//create hash key for this object
	hash := getMd5Hash(f)
	// check cache map
	val, ok := tinyCache[hash]
	if ok == true {
		return val, true
	}
	return tempInC{}, false
}

func saveToCache(f tempInF, c tempInC) {
	hash := getMd5Hash(f)
	// log.Println("Hash to use...", hash)
	// save to map cache
	tinyCache[hash] = c
}

func fahrenheidToCelsius(f tempInF) tempInC {

	c, ok := getFromCache(f)
	if ok == true {
		log.Println("Returning from cache...")
		return c
	}

	log.Println("Performing calculation...")
	time.Sleep(2 * time.Second)

	nc := tempInC{
		name:  "Celsius",
		value: (f.value - 32) * 5 / 9,
	}

	saveToCache(f, nc)

	return nc
}

func main() {

	f := getFahrenheitTemp(25)
	// firts call
	c := fahrenheidToCelsius(f)
	// second call
	c2 := fahrenheidToCelsius(f)

	log.Printf("Convert %v to %v", f.name, c.name)

	log.Println("Temp in F: ", f.value)
	log.Println("Temp in C: ", c.value)
	log.Println("Temp in C: ", c2.value)

}
