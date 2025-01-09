package helpers

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"regexp"
)

func PrintBody(response *http.Response) {
	body, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatal("Error reading response body: ", err)
	}

	// fmt.Println(string(body[:]))
	links, err := CollectLinks(body)
	if err != nil {
		log.Fatal("Could not find links? ", err)
	}

	for _, link := range links {
		fmt.Println(link)
	}
}

func CollectLinks(responseBody []byte) ([]string, error) {
	body := string(responseBody)

	pattern := `href=['"](.*?)['"]`

	// Compile regex
	re, err := regexp.Compile(pattern)
	if err != nil {
		return nil, err
	}

	// Find all matches
	matches := re.FindAllStringSubmatch(body, -1)

	// Extract all links from matches
	result := make([]string, 0, len(matches))

	for _, match := range matches {
		if len(match) > 1 {
			// match[0] is the full match (href="..."), match[1] is the captured group (the URL)
			result = append(result, match[1])
		}
	}

	return result, nil
}
