package main

import (
	"fmt"
	"net/http"

	"golang.org/x/net/html"
)

func main() {
	resp, err := http.Get("https://kolesa.kz/cars/audi/")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	doc, err := html.Parse(resp.Body)
	if err != nil {
		panic(err)
	}
	carLinks := extractCarLinks(doc)

	for _, link := range carLinks {
		carResp, err := http.Get(link)
		if err != nil {
			continue
		}
		defer carResp.Body.Close()

		carDoc, err := html.Parse(carResp.Body)
		if err != nil {
			continue
		}
		carDetails := extractCarDetails(carDoc)

		// Do something with the car details
		fmt.Println(carDetails)
	}

	//Manually set the number of pages

	for i := 2; i <= 351; i++ {
		pageURL := fmt.Sprintf("https://kolesa.kz/cars/audi/?page=%d", i)
		pageResp, err := http.Get(pageURL)
		if err != nil {
			continue
		}
		defer pageResp.Body.Close()

		// Extract car URLs from page and repeat Step 3
		pageDoc, err := html.Parse(pageResp.Body)
		if err != nil {
			continue
		}
		carLinks := extractCarLinks(pageDoc)
		for _, link := range carLinks {
			// Repeat Step 3
		}
	}
}

// Helper functions to extract car URLs and details using HTML parsing
func extractCarLinks(doc *html.Node) []string {
	// Implement logic to extract car links from HTML document
	return nil
}

func extractCarDetails(doc *html.Node) map[string]string {
	// Implement logic to extract car details from HTML document
	return nil
}

func step3(carLinks []string) {

	for _, link := range carLinks {
		carResp, err := http.Get(link)
		if err != nil {
			continue
		}
		defer carResp.Body.Close()

		carDoc, err := html.Parse(carResp.Body)
		if err != nil {
			continue
		}
		carDetails := extractCarDetails(carDoc)

		// Do something with the car details
		fmt.Println(carDetails)
	}
}
