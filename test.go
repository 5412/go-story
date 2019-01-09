package main

import (
	"fmt"
	"sync"
)

func main() {
	var (
		wg      sync.WaitGroup
		numbers []int
	)

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			numbers = append(numbers, i)
			wg.Done()
		}(i)
	}

	wg.Wait()

	fmt.Println("The numbers is", numbers)
}
