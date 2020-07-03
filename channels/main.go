package main

import (
	"log"
	"time"
)

type actionType struct {
	Type    string
	Payload interface{}
}

func process1(ch chan actionType) {
	// sleep 1 second
	time.Sleep(1 * time.Second)
	ch <- actionType{
		Type:    "PROCESS_1",
		Payload: []string{"Value1", "Value2", "Value3"},
	}
}

func process2(ch chan actionType) {
	// sleep 1 second
	time.Sleep(3 * time.Second)
	ch <- actionType{
		Type:    "PROCESS_2",
		Payload: []int32{1, 2, 3, 4, 5, 6},
	}
}

func counter(ch chan actionType) {
	switch ch.Type {
	case "INC":
		ch.Payload = ch.Payload + 1
	case "DEC":
		ch.Payload = ch.Payload - 1
	}
}

func main() {
	log.Println("Starting app")

	ch1 := make(chan actionType)
	// ch2 := make(chan actionType)
	// go process1(ch1)
	for {
		go process1(actionType)
		r := <-ch1
		log.Println(<-ch1)
	}
	// go process1(ch1)
	// go process2(ch2)
	// log.Println("Clogin app")
}
