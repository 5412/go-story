package main

import "fmt"

func myPrintf(args ...interface{}) {
    for _, arg := range args {
        switch arg.(type) {
            case int:
                fmt.Println(arg, "is a int value")
            case string:
                fmt.Println(arg, "is a string value")
            default:
                fmt.Println(arg, "I don't know is type")
        }
    }
}

func main() {
    var v1 = 123
    var v2 = "hello"
    myPrintf(v1)
    myPrintf(v2)
}
