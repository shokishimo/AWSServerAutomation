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
	_, err := http.Get("https://wbk.onrender.com/")
	if err != nil {
		fmt.Println("Error making request to WBK website:", err)
	} else {
		fmt.Println("Request to WBK website Succeeded.")
	}
}