package main

import (
	"fmt"
	"math/rand"
)

func main() {
	randInt := rand.Int63n(10e7)
	fmt.Println(randInt)
}
