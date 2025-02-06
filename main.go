package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"
	"sync"
	"time"
)

func Pipeline(ctx context.Context, input []int) (int, error) {
	inputChan := make(chan int, len(input))
	processChan := make(chan int, len(input))

	var wg sync.WaitGroup
	var result int
	var mu sync.Mutex

	// Input channel population
	wg.Add(1)
	go func() {
		defer wg.Done()
		defer close(inputChan)
		for _, num := range input {
			select {
			case <-ctx.Done():
				return
			case inputChan <- num:
			}
		}
	}()

	// Processing goroutine
	wg.Add(1)
	go func() {
		defer wg.Done()
		defer close(processChan)

		for num := range inputChan {
			select {
			case <-ctx.Done():
				return
			default:
				if num < 0 {
					return
				}

				time.Sleep(100 * time.Millisecond)
				squared := num * num

				if squared <= 1000 {
					processChan <- squared
				}
			}
		}
	}()

	// Result aggregation
	wg.Add(1)
	go func() {
		defer wg.Done()
		for num := range processChan {
			mu.Lock()
			result += num
			mu.Unlock()
		}
	}()

	wg.Wait()

	// Check context status
	if err := ctx.Err(); err != nil {
		return 0, err
	}

	return result, nil
}

func main() {
	input := make([]int, 0)
	for _, arg := range os.Args[1:] {
		num, err := strconv.Atoi(arg)
		if err != nil {
			log.Fatalf("Invalid input: %v", err)
		}
		input = append(input, num)
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	result, err := Pipeline(ctx, input)
	if err != nil {
		log.Fatalf("Pipeline error: %v", err)
	}

	fmt.Printf("Result: %d\n", result)
}
