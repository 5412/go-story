package main

import "fmt"
/******************************************/
/** 自定义结构体 type 结构体名称 struct {}
/******************************************/
type PersonInfo struct {
    ID string
    Name string
    Address string
}

func main() {
    /******************************************/
    /** 定义map类型变量
    /******************************************/
    var PersonDB map[string] PersonInfo
    PersonDB = make(map[string] PersonInfo)
    
    /******************************************/
    /** 字典赋值
    /******************************************/
    PersonDB["12345"] = PersonInfo{"12345", "piaopiao", "beijing huoying"}
    PersonDB["12"] = PersonInfo{"12", "yujindan", "beijinghuoying"}

    /******************************************/
    /** 获取字典中元素的值
    /******************************************/
    Person, ok := PersonDB["12"]
    if ok {
        fmt.Println("The person's name is", Person.Name)
    } else {
        fmt.Println("Did not find person with ID 12")
    }
}
