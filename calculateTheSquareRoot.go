package main

import (
	"fmt"
	"math"
)

func main()  {
	fmt.Println(sqrt(255 ))
	fmt.Println(math.Sqrt(255))
}

func sqrt(num float64) float64  {
	iteration := find(num, 2)
	for math.Abs(num - iteration*iteration) > 0.0000000000001 {
		iteration = find(num, iteration)
	}
	return iteration
}

func find(num, iteration float64) float64  {
	iteration = (iteration + num/iteration)/2
	return  iteration
}


