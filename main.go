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
	fmt.Printf("Request Method: %s\n", req.Method)
	fmt.Printf("Response Status: %s\n", resp.Status)
	fmt.Printf("Response Status Code: %d\n", resp.StatusCode)
	fmt.Println("---------------------------------------------")

	fmt.Println("Let's test this out:")

	links, err := helpers.CheckLinks(resp)
	if err != nil {
		log.Fatal("Could not check links: ", err)
	}

	for _, link := range links {
		fmt.Println(link)
	}

}
