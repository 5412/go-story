package main

import "fmt"

func main() {
	for i := 0; i < 5; i++ {
		defer func() {
			fmt.Println(i)
		}()
	}

	for i := 0; i < 5; i++ {
		defer func(i int) {
			fmt.Println(i)
		}(i)
	}

	fmt.Println(deferTest())

}

func deferTest() (i int) {
	defer func() {
		fmt.Println(i)
		i = 4
	}()
	return 2
}
