package main

import (
	"fmt"
	"go-async/common"
	_ "io/ioutil"
	"net/http"
	_ "reflect"
	_ "regexp"
	"time"
)

// HTTPResponse -> shut up golint
type HTTPResponse struct {
	url      string
	response *http.Response
	err      error
}

func asyncHTTPGets(urls []string) []*HTTPResponse {
	ch := make(chan *HTTPResponse, len(urls)) // buffered
	responses := []*HTTPResponse{}
	for _, url := range urls {
		go func(url string) {
			fmt.Printf("Fetching %s \n", url)
			resp, err := http.Get(url)
			if err == nil {
				resp.Body.Close()
			}
			ch <- &HTTPResponse{url, resp, err}
		}(url)
	}

	for {
		select {
		case r := <-ch:
			fmt.Printf("%s was fetched\n", r.url)
			responses = append(responses, r)
			if len(responses) == len(urls) {
				return responses
			}
		case <-time.After(50 * time.Millisecond):
			fmt.Printf(".")
		}
	}
}

func main() {
	ids := []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10", "11", "12", "13", "14", "15"}
	urlBase := "https://jsonplaceholder.typicode.com/todos/"
	var urls []string
	for _, id := range ids {
		url := urlBase + id
		urls = append(urls, url)
	} 
	urlSlices := common.SplitToChunks(urls, 5)
	start := time.Now()
	for _, urlSlice := range urlSlices {
		results := asyncHTTPGets(urlSlice)
		for _, result := range results {
			if result.err != nil {
				fmt.Printf("%s error: %v\n", result.url,
				result.err)
				continue
			}
			fmt.Printf("%s status: %s\n", result.url,
				result.response.Status)
		}
	}
	end := time.Now()
	duration := end.Sub(start)
	fmt.Printf("Duration: %vs\n", duration.Seconds())
	
}


