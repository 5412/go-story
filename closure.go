package main

import "fmt"

func main() {
    i := 10

    a := func()(func()){
        j:=i
        return func(){
            fmt.Println("i is", i, "j is ", j)
        }
    }()
    b := a
    a()

    i= 20
    j:= 20
    j++

    b()
    b()
}
