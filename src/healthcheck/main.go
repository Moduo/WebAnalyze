package main

import (
	"time"
	"log"
	"net/http"
	"fmt"
)

func main() {
	url := "https://www.swiftdevelopment.nl/"
	poll(url)
}

func isHealthy (url string) bool {
	resp, err := http.Get("https://golangcode.com")
    if err != nil {
        log.Fatal(err)
	}
	
    if resp.StatusCode == 200 {
        return true
    }
	return false
}

func poll (url string) {
	for true {
		if isHealthy(url){
			fmt.Printf("%s is healthy\n", url)
		} else {
			fmt.Printf("%s is not healthy\n", url)
		}
		time.Sleep(1000 * time.Millisecond)
	}
}