package main

import "fmt"

func main () {
    i := make(chan int, 10)
    j := 0
    for {

        fmt.Println("Loop index is ", j)
        select {
            case i <- 0:
                fmt.Println("0 value is <-")
            case i <- 1:
                fmt.Println("1 value is <-")
            default:
                fmt.Println("Nothing done")
        }
        j++

        if j > 1024 {
            break
        }
    }

    //for c:=0; c < 11; c++ {
      // v, ok :=  <-i
       //if ok {
         //   fmt.Println("Value recieve is ", v)
       //} else {
         //   close(i)
           // fmt.Println("Closing chan")
       //}
       //fmt.Println("Value now is ", v)
    //}
    for k := range i {
         fmt.Println("Recieve chan ", k)
    }
}
