package main

import (
	"reflect"
	"solar/fmt"
)

func main()  {

    //fmt.Println(strconv.ParseFloat("3.1", 64))
	const maxUint64 = (1<<63 - 1)
	t := reflect.TypeOf(uint64(maxUint64))
	fmt.Println(t.String())

	/*
	const uintSize = 32 << (^uint(0) >> 63)
	const maxShift = uintSize - 4
	fmt.Printf("%b \n", 32 << (^uint(0) >> 63))
	fmt.Println(uintSize)
	fmt.Println(maxShift)
	*/
}
