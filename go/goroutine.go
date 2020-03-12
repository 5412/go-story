package main

import "fmt"
import "time"

func Add(x, y int) {
    z := x + y
    fmt.Println("z :", z)
}

func main() {
    go Add(1,1)
    time.Sleep(10000)
    fmt.Println("Programing is here.")
}
