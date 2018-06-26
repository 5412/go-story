package main

import "fmt"

func main() {
    /********************************/
    /** label 标签先定义后使用
    /********************************/
    JLoop:
        for i:=0;i<5;i++ {
            for j:=0;j<10;j++{
                if i == 4 {
                    break JLoop
                }
                fmt.Println("i is %d,j is %d", i, j)
            }
        }
    //
}
