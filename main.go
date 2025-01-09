package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/UberChili/scrapper/helpers"
)

func main() {
	client := &http.Client{}

	// req, err := http.NewRequest("GET", "https://scrape-me.dreamsofcode.io/in-utero", nil)
	req, err := http.NewRequest("GET", "https://scrape-me.dreamsofcode.io/nirvana", nil)
	if err != nil {
		log.Fatal("Error creating request: ", err)
	}

	// Send the request
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("Error sending request: ", err)
	}
	defer resp.Body.Close() // Don't forget to close the response body

	// Print request and response information
	fmt.Printf("Request Method: %s", req.Method)
	fmt.Printf("Response Status: %s", resp.Status)
	fmt.Printf("Response Status Code: %d", resp.StatusCode)

	fmt.Println("Body:")

	helpers.PrintBody(resp)
}
