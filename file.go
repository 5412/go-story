package main

import (
	"os"
	"fmt"
)
func main()  {
	file,err := os.OpenFile("./test.text", 1, 0777)
	if err != nil {
		fmt.Println(err)
	}
	file.WriteAt()
	
}
