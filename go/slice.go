package main
import "fmt"

func main() {
    /***********************************************/
    /** 定义数组
    /***********************************************/
    var array [10]int = [10]int{1,2,3,4,5,6,7,8,9,0}
    var slice []int = array[:5]
    /***********************************************/
    /** 定义切片[first:last]
    /***********************************************/
    myslice := array[6:7]
    /***********************************************/
    /** append 向切片中追加 如果第二个参数是切片需要
    /** 多加3个. 
    /***********************************************/
    slice = append(slice, myslice...)
    fmt.Println("The element in array is :")
    for _,v := range array {
        fmt.Print(v, " ")
    }
    fmt.Println("\nThe element of slice is :")
    for i,v := range slice {
        fmt.Println("index :", i, "value:", v)
    }
    fmt.Println("This is the test for slice of array", array[9])
    /***********************************************/
    /** cap()返回切片空间大小，len()返回切片元素个数
    /***********************************************/
    fmt.Println("The cap(slice) is", cap(slice))
    fmt.Println("The len(slice) is", len(slice))

    slice1 := []int{1,2,3,4,5}
    slice2 := []int{4,5,6}

    copy(slice1,slice2)

    for i,v := range slice1 {
        fmt.Println(i,v)
    }
}
