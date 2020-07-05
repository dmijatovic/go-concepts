package main

import (
	"log"
	"time"
)

type actionType struct {
	Type string
	// Payload interface{}
	Sum int32
}

func process1(action *actionType, ch chan string) {
	// sleep 1 second
	time.Sleep(1 * time.Second)
	action.Type = "PROCESS_1"
	action.Sum = action.Sum - 1
	ch <- "DONE"
}

func process2(action *actionType, ch chan string) {
	// sleep 1 second
	time.Sleep(3 * time.Second)
	action.Type = "PROCESS_2"
	action.Sum = action.Sum + 3
	ch <- "DONE"
}

func syncronize() {
	action := actionType{
		Type: "ADD",
		Sum:  0,
	}
	log.Println("Synchornized process 1 & 2")
	ch1 := make(chan string)
	// ch2 := make(chan actionType)
	// synchronising processes
	// process1 takes 1 second
	// process2 takes 3 secdons
	// each loop will wait on process2
	// to return before making next cicle!!!
	round := 0
	for {
		round++
		log.Println("Round:", round)
		go process1(&action, ch1)
		<-ch1
		log.Println(action)
		go process2(&action, ch1)
		<-ch1
		log.Println(action)
		// do just 5 rounds
		if round == 5 {
			break
		}
	}
}

func firstProcess1ThenProcess2() {
	action := actionType{
		Type: "ADD",
		Sum:  0,
	}
	log.Println("In order, FIRST process 1 THEN 2")
	ch1 := make(chan string)
	// ch2 := make(chan actionType)
	round := 0
	for {
		round++
		go process1(&action, ch1)
		<-ch1
		log.Println(action)
		// do just 5 rounds
		if round == 5 {
			break
		}
	}
	round = 0
	for {
		round++
		go process2(&action, ch1)
		<-ch1
		log.Println(action)
		// do just 5 rounds
		if round == 5 {
			break
		}
	}
}

// Running process 1 and 2 independently
func independet() {
	action := actionType{
		Type: "ADD",
		Sum:  0,
	}
	log.Println("Independent RUN process 1 & 2")
	ch1 := make(chan string)
	ch2 := make(chan string)
	// process 1 takes inital action
	// add performs its mutations
	// independently from process 2
	// the action object is not shared
	// it is passed by value at the begining
	// of the loop
	go func(action actionType) {
		r := 0
		for {
			r++
			go process1(&action, ch1)
			<-ch1
			log.Println(action)
			// do just 5 rounds
			if r == 5 {
				break
			}
		}
	}(action)
	round := 0
	for {
		round++
		go process2(&action, ch2)
		<-ch2
		log.Println(action)
		// do just 5 rounds
		if round == 5 {
			break
		}
	}
}

func main() {
	log.Println("Starting app")
	// process 1 and process 2 run synchronized each cycle
	syncronize()
	// process 1 runs first all loops THEN process 2 runs
	firstProcess1ThenProcess2()
	// process 1 and 2 run independent loops and have their own values
	independet()
}
