package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	start := time.Now()

	apis := []string{
		"https://management.azure.com",
		"https://dev.azure.com",
		"https://api.github.com",
		"https://outlook.office.com",
		"https://api.somewhereintheinternet.com",
		"https://graph.microsoft.com",
	}
	ch := make(chan string) // create a channel of type string
	for _, api := range apis {
		go checkAPI(api, ch) // apply go keyword to run the function in a goroutine (concurrency)
	}

	for i := 0; i < len(apis); i++ {
		fmt.Print(<-ch) // receive the value from the channel
	}

	elapsed := time.Since(start) // time.Since returns a time.Duration type value
	fmt.Printf("Donde in %s \n", elapsed)
}

func checkAPI(api string, ch chan string) {
	_, err := http.Get(api)
	if err != nil {
		ch <- fmt.Sprintf("Error: %s is down \n", api)
		return
	}
	ch <- fmt.Sprintf("Success: %s is up \n", api)
}
