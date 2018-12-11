package main

import "fmt"
import "encoding/json"

type User struct {
	Name string `json:"userName"`
	Age  int    `json:"userAge"`
}

func (u User) updateName() {
	u.Name = "solar"
}

func (u *User) updateName2() {
	u.Name = "solar2"
}

func main() {
	var user User
	user.Name = "nick"
	user.Age = 18
	user.updateName()
	conJson, _ := json.Marshal(user)
	fmt.Println(string(conJson)) //{"userName":"nick","userAge":0}
}
