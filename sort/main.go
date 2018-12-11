package main

import (
	"solar/fmt"
)


func maxDepth(n int) int {
	var depth int
	for i := n; i > 0; i >>= 1 {
		depth++
	}
	return depth * 2
}

func main()  {
	fmt.Println("Testing Sort Package begin")
	lo := 0
	hi := 7
	m := int(uint(lo+hi) / 2)
	fmt.Println(m)

	/*
	var s string
	fmt.Printf("Pick an integer from 0 to 100.\n")
	answer := sort.Search(100, func(i int) bool {
		fmt.Printf("Is your number <= %d? ", i)
		fmt.Scanf("%s", &s)
		return s != "" && s[0] == 'y'
	})
	fmt.Printf("Your number is %d.\n", answer)
	/*
	var slice []int = []int{
		9,2,3,5,7,8,3,0,11,29,32,13,1,
	}

	sort.Ints(slice)

	fmt.Println(slice)
	*/
}
