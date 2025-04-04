package main

import (
	"fmt"
	"strconv"
	"sync"
	"time"

	"github.com/ArturDV-main/Web-App-Advanced/internal/nserver"
)

func main() {
	start := time.Now()
	fmt.Println(start)
	channel1 := make(chan string)
	channel2 := make(chan string)
	channel3 := make(chan string)
	go func() {
		defer func() {
			close(channel1)
			close(channel2)
			close(channel3)
		}()
		// add data to channels
		datasize := 100
		for i := range datasize {
			channel1 <- "channel - 1 " + strconv.Itoa(i)
			channel2 <- "channel - 2 " + strconv.Itoa(i+1)
			channel3 <- "channel - 3 " + strconv.Itoa(i+3)
			time.Sleep(300)
		}
	}()

	// read data from channels
	for value := range MergeChannels(channel1, channel2, channel3) {
		fmt.Println(value)
	}
	fmt.Println(time.Since(start))
	nserver.StartServer()
}

func MergeChannels[T any](channels ...chan T) <-chan T {
	var wg sync.WaitGroup
	wg.Add(len(channels))
	mergedChannel := make(chan T)

	for _, channel := range channels {
		go func() {
			defer wg.Done()
			for value := range channel {
				mergedChannel <- value
			}
		}()

	}
	go func() {
		wg.Wait()
		close(mergedChannel)
	}()

	return mergedChannel
}
