package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/UberChili/scrapper/helpers"
)

func main() {
	urlPtr := flag.String("url", "", "The url to scrap")
	flag.Parse()

	client := &http.Client{}

	if *urlPtr == "" {
		log.Fatal("No url provided")
	}

	req, err := http.NewRequest("GET", *urlPtr, nil)
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

	links, err := helpers.CheckLinks(resp)
	if err != nil {
		log.Fatal("Could not check links: ", err)
	}

	for _, link := range links {
		fmt.Println(link)
	}

}
