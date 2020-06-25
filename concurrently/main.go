package main

import (
	"net/http"
	"time"
)

func checkURL(url string, ch chan string) {

	_, err := http.Get(url)
	// wait 3 seconds
	time.Sleep(3 * time.Second)
	// provide end result
	if err != nil {
		println("Site down on GET", url)
		// ch <- fmt.Sprintf("Site is down on GET...%v", url)
		ch <- url
	} else {
		println("Active...", url)
		// ch <- fmt.Sprintf("Active...%v", url)
		ch <- url
	}
}

func main() {
	println("It works")

	sites := []string{
		"http://google.com",
		"http://microsoft.com",
		"http://facebook.com",
	}

	//create new string channel
	ch := make(chan string)

	// issue task to child routines
	for _, url := range sites {
		go checkURL(url, ch)
	}
	// geather responses
	// note! this loop is executing one itteraction
	// on each response. Waiting on message from channel
	// is blocking call
	// for i := 0; i < len(sites); i++ {
	// 	// this will wait for response
	// 	// from the channel
	// 	// blocking - will wait
	// 	fmt.Println(<-ch)
	// }

	// use endles loop
	// for {
	// 	//blocking call
	// 	//it will wait untill
	// 	go checkURL(<-ch, ch)
	// }

	// alterantive endless loop
	// waiting for channel to respond
	// for l := range ch {
	// 	go checkURL(l, ch)
	// }

	// endles loop using function literal (lambda)
	for l := range ch {
		// function literal (ano fn)
		// NOTE! make sure you pass params
		// into anonymous fn because it is
		// in different scope
		go func(url string) {
			time.Sleep(3 * time.Second)
			checkURL(url, ch)
		}(l)
	}
}
