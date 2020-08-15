package common

import (      
        "fmt" 
        "sync"
)
// MergeChannels -> merges a slice of channels into a single channel
func MergeChannels(cs []chan string) <-chan string {
	out := make(chan string)
	var wg sync.WaitGroup
	wg.Add(len(cs))
	for _, c := range cs {
		go func(c <-chan string) {
			for v := range c {
				out <- v
			}
			wg.Done()
		}(c)
	}
	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}

// SplitToChunks -> splits slice to n number of slices
func SplitToChunks(slice []string, chunkNum int) [][]string {
        var chunked [][]string
        chunkSize := (len(slice) + chunkNum - 1) / chunkNum
        for i := 0; i < len(slice); i += chunkSize {
                end := i + chunkSize
                if end > len(slice) {
                        end = len(slice)
                }
                chunked = append(chunked, slice[i:end])
		}
	return chunked
}

// HandleErrors -> log error
func HandleErrors(err error) {
        if err != nil {
                fmt.Println("An error occurred", err)
        }
}