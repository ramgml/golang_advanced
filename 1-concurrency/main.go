package main

import (
	"fmt"
	"math/rand"
	"sync"
)


func main() {
	numCh := make(chan int)
	resultCh := make(chan int)
	var wg sync.WaitGroup
	go getRandomNumbers(numCh)
	go getSquare(numCh, resultCh)
	var results []int
	wg.Add(1)
	go func() {
		defer wg.Done()
		for square := range resultCh {
			results = append(results, square)
		}
	}()
	wg.Wait()
	fmt.Println(results)
}

func getRandomNumbers(numCh chan int) {
	for range 10 {
		numCh <- rand.Intn(101)
	}
	close(numCh)
}

func getSquare(numCh chan int, resultCh chan int) {
	for num := range numCh {
		square := num * num
		resultCh <- square
	}
	close(resultCh)
}