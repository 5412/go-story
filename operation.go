package main

import "fmt"

func main()  {
	fmt.Println("Print my test fo abs arithmetic")
	var (
		x int64
		y int64
		)
	x = -1324234234
	y = x >> 63

	fmt.Printf("x value is %d \n", x)
	fmt.Printf("x > 63 value is %d \n", y)

	c := x ^ y

	fmt.Printf("x ^ y value is %d \n", c)

	d := c - y
	fmt.Printf("(x ^ y) - y value is %d \n", d)

	f := ^x
	fmt.Printf("(^x value is %d \n", f)

	g := ^f
	fmt.Printf("(^^x value is %d \n", g)

}
