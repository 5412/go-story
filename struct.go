package main

import "fmt"
import "encoding/json"

type User struct {
    Name string `json:"userName"`
    Age  int    `json:"userAge"`
}

func main() {
    var user User
    user.Name = "nick"
    user.Age = 18
    conJson, _ := json.Marshal(user)
    fmt.Println(string(conJson))    //{"userName":"nick","userAge":0}                            
}
