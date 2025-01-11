package helpers

import (
	"errors"
	"io"
	"log"
	"net/http"
	"regexp"
	"strings"
)

const host = "https://scrape-me.dreamsofcode.io/"

// This can probably be just a slice, since I am not really using the values
var statuses = map[int]string{
	400: "Bad Request (Server couldn't understand the request)",
	401: "Unauthorized (Authentication required)",
	403: "Forbidden (Client doesn't have access rights)",
	404: "Not Found (Resource doesn't exist)",
	405: "Method Not Allowed (HTTP method not supported)",
	429: "Too Many Requests (Rate limit exceeded)",
}

func CheckLinks(response *http.Response) ([]string, error) {
	sites := []string{}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatal("Error reading response body: ", err)
	}

	// fmt.Println(string(body[:]))
	links, err := CollectLinks(body)
	if err != nil {
		log.Fatal("Could not find links? ", err)
	}

	// for _, link := range links {
	// 	fmt.Println(link)
	// }

	for _, link := range links {
		result, err := CheckStatus(link)
		if err != nil {
			continue
		}
		sites = append(sites, result)
	}
	return sites, nil
}

func CheckStatus(link string) (string, error) {
	temp_client := &http.Client{}

	site := host + strings.TrimPrefix(link, "/")
	req, err := http.NewRequest("GET", site, nil)
	if err != nil {
		log.Fatal("Error creating request: ", err)
		return "", err
	}

	resp, err := temp_client.Do(req)
	if err != nil {
		log.Fatal("Error sending request: ", err)
	}
	defer resp.Body.Close()

	_, exists := statuses[resp.StatusCode]
	if exists {
		return "", errors.New("Link is dead.")
	}
	return site, nil
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
			if strings.HasSuffix(match[1], ".css") {
				continue
			}
			// match[0] is the full match (href="..."), match[1] is the captured group (the URL)
			result = append(result, match[1])
		}
	}

	return result, nil
}
