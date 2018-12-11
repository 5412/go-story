package main

import (
	"fmt"
	"time"
)

func main() {
	satr := "Hello World, 你好，世界"
	length := len(str)

	for i := 0; i < length; i++ {
		//fmt.Printf("The %d element is : %d \n", i, str[i])
	}

	for i, ch := range str {
		fmt.Println("index", i, "value", ch)
	}

	time.Sleep(300)

}
