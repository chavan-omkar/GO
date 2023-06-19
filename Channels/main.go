package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://amazon.com",
	}

	channel := make(chan string)

	for _, link := range links {
		go checkLinks(link, channel)
	}
	/*
		we should wait for all the responses going to be published on channel the below code won't work it will
		print the single response as soon as the response is published on that channel and will terminate the program

		so the main thing here to keep in mind that receiving a msg on channel will be a blocking call
	*/
	// fmt.Println(<-channel)
	// fmt.Println(<-channel)
	// fmt.Println(<-channel)
	// fmt.Println(<-channel)
	// fmt.Println(<-channel)
	/* we have 5 websites for which the msg will be sent on channel if we add more no of listen from channel in that case
	the main thread will be kept ideal and waiting for the msg to receive on that channel for infinite time.
	*/
	// fmt.Println(<-channel)

	// for i := 0; i < len(links); i++ {
	// 	fmt.Println(<-channel)
	// }

	// This for loop is infinite loop
	// for {
	// 	go checkLinks(<-channel, channel)
	// }

	// for l := range channel {
	// 	go checkLinks(l, channel)
	// }

	for l := range channel {
		/*
			Here we have passed the value in l variable to go rotine and then we have passed that value to checkLink function as it is
			better way to not use single variable in different go routines directly

			instead of that pass the value received by main routine to another go routine in the form of argument and then use that value,
			as Go is pass by value or try to access it by channels

			Never access same variable from different child routine
		*/
		go func(link string) {
			time.Sleep(5 * time.Second)
			checkLinks(link, channel)
		}(l)
	}

}

func checkLinks(link string, c chan string) {
	_, err := http.Get(link)

	if err != nil {
		fmt.Println(link, "might be facing some issues!")
		c <- link
		return
	}

	fmt.Println(link, "is up and working fine!")
	c <- link

}
