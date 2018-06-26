package main
import "fmt"
func main() {
    //var arr  [32]int
    //for i := 0;i<len(arr);i++ {
    //   arr[i] = i + 1
   // }
    array := [5]int{1,2,3,4,5}
    for i,v := range array {
        fmt.Println("Element", i, "'s value is", v)
    }
}

