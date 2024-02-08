package main

import (
	"fmt"
	"log"
	"sync"
)

func makeThumbnails(filenames <-chan string) int64 {
	sizes := make(chan int64)
	var wg sync.WaitGroup
	for f := range filenames {
		wg.Add(1)
		// worker
		go func(f string) {
			defer wg.Done()
			fileSize, err := getImageFile(f)
			if err != nil {
				log.Println(err)
				return
			}
			sizes <- fileSize
		}(f)
	}
	// closer
	go func() {
		wg.Wait()
		close(sizes)
	}()
	var total int64
	for size := range sizes {
		total += size
	}
	return total
}

func getImageFile(filename string) (int64, error) {
	fileStats := map[string]int64{
		"main.go":    10,
		"channel.go": 15,
	}
	return fileStats[filename], nil
}

func main() {
	filenames := make(chan string)
	filenames <- "main.go"
	filenames <- "channel.go"
	totalFilesSize := makeThumbnails(filenames)
	fmt.Println("Total files size is ", totalFilesSize)
}
