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
	c := make(chan string)

	for _, link := range links {
		go checkLink(link, c)
	}
	//receive a message from channel
	// mystring := <-c
	// fmt.Println(<-c)
	// for i := 0; i < len(links); i++ {
	// 	//receive value
	// 	checkLink(<-c)
	// }

	//infinite loop but when receive only a value in channel
	for l := range c {
		//function literals
		//pass the value l to the function literal to copy the memory address
		go func(link string) {
			time.Sleep(5 * time.Second)
			go checkLink(link, c)
		}(l)
	}

}
func checkLink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "might be down")
		//send to our channel
		// c <- "Might be down I think"
		c <- link
		return
	}
	fmt.Println(link, "is up!")
	c <- link
}

//go routines run parallel
