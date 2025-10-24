package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"strings"
	
	"golang.org/x/net/html"
)

func getHTML(rawURL string) (string, error) {
	client := &http.Client{}
	
	req, err := http.NewRequest(http.MethodGet, rawURL, nil)
	if err != nil {
		return "", fmt.Errorf("Error '%v' occured while making new request.", err)
	}
	
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/71.0.3578.98 Safari/537.36")
	req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8")
	req.Header.Set("Accept-Language", "en-US,en;q=0.5")
	req.Header.Set("DNT", "1")
	req.Header.Set("Connection", "close")
	
	res, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("Error '%v' occured while executing request.", err)
	}
	defer res.Body.Close()
	
	if res.StatusCode >= 400 {
		log.Println(res.Header)
		return "", fmt.Errorf("Bad response code: '%v'", res.StatusCode)
	}
	
	contentType := res.Header.Get("Content-Type")
	if !strings.Contains(contentType, "text/html") {
		return "", fmt.Errorf("Error: '%v'", err)
	}
	
	body, _ := io.ReadAll(res.Body)
//	log.Println(string(body))
	return string(body), nil
}

func normalizeURL(rawURL string) (string, error) {
	parsedURL, err := url.Parse(rawURL)
	if err != nil {
		return "", fmt.Errorf("Error: '%v'", err)
	}
	
	fullPath := parsedURL.Host + parsedURL.Path
	fullPath = strings.ToLower(fullPath)
	fullPath = strings.TrimSuffix(fullPath, "/")
	
	return fullPath, nil
}

// Needs configured to work with yahoo
func getURLsFromHTML(htmlBody string, baseURL *url.URL) ([]string, error) {
	body, err := html.Parse(strings.NewReader(htmlBody))
	if err != nil {
		return []string{}, fmt.Errorf("Error parsing HTML Body")
	}
	
	var URLs []string
	for n := range body.Descendants() {
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, a := range n.Attr {
				if a.Key == "href" {
//					"Need to include the query alterations; to change pages on screeners"
					basePath, _ := url.Parse(a.Val)
					resolvedURL := baseURL.ResolveReference(basePath)
					URLs = append(URLs, resolvedURL.String())
				}
			}	
		}
	}
		
	return URLs, nil
}
