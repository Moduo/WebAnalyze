package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

var(
	quit = false
)

func main() {
	urls := []string{
		"https://www.swiftdevelopment.nl/",
		"https://www.google.nl"}
	pollSet(urls...)
}

func isHealthy(url string) bool {
	resp, err := http.Get("https://golangcode.com")
	if err != nil {
		log.Fatal(err)
	}

	if resp.StatusCode == 200 {
		return true
	}
	return false
}

func poll(url string) {
	output := ""
	for {
		if isHealthy(url) {
			output = "%s is healthy"
		} else {
			output = "%s is not healthy"
		}
		fmt.Printf(output + "\n", url)
		time.Sleep(1000 * time.Millisecond)
	}
}

func pollSet(urls ...string){
	for _, url := range urls {
		go poll(url)
	}
	for !quit{
		time.Sleep(1000 * time.Millisecond)
	}
}

func stopPolling(){
	quit = true
}