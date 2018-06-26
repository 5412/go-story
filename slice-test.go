package main

import "fmt"

func main() {
    slice := [][]int{{10}, {100,200}}
    var array [5][5]int
    fmt.Println("This is testing multidimensional sclice")
    fmt.Println(slice)
    fmt.Println("This is testing multidimensional array")
    fmt.Println(array)

    slice2 := make([]int, 1e6)
    fmt.Println(slice2)
}
