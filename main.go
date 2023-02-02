package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	for {
		makeRequest()
		time.Sleep(14*time.Minute)
	}
}

func makeRequest() {
	_, err := http.Get("https://distinctionality.onrender.com/")
	if err != nil {
		fmt.Println("Error making request to Distinctionality:", err)
	} else {
		fmt.Println("Request to Distinctionality Succeeded.")
	}

	_, err = http.Get("https://one-tap.onrender.com")
	if err != nil {
		fmt.Println("Error making request to One Tap:", err)
	} else {
		fmt.Println("Request to One Tap Succeeded.")
	}
}