package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func main() {
	runtime.GOMAXPROCS(1)
	fmt.Println("My goroutine testing!")
	fmt.Println("My CPU num is :", runtime.NumCPU())
	fmt.Println("Goroutine num is", runtime.NumGoroutine())
	wg.Add(2)
	fmt.Println("Start Goroutines")
	/*go func() {
		defer wg.Done()
		for count := 0; count < 10; count++ {
			for char := 'a'; char < 'a' + 26; char++ {
				fmt.Printf("%c 1", char)
			}
		}
	}()

	go func() {
		defer wg.Done()
		for count := 0; count < 10; count++  {
			for char := 'A'; char < 'A' + 26; char++ {
				fmt.Printf("%c 2", char)
			}
		}
	}()*/
	go printPrime("A")
	go printPrime("B")
	fmt.Println("Waiting To Finish")
	wg.Wait()
	fmt.Println("\nTerminating Program")
}

func printPrime(prefix string) {
	defer wg.Done()
	next:
		for outer := 2; outer < 100; outer++ {
			for inner := 2; inner < outer; inner++ {
				if outer % inner == 0 {
					continue next
				}
			}
			fmt.Printf("%s:%d\n", prefix, outer)
		}
		fmt.Println("Completed", prefix)
}