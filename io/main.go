package main

import "fmt"
import (
	"study/io/ioutil"
)

func main() {
	fmt.Println("Test io package")
	filename := "/Users/piaoweixiong/phptest/test.php"
	buf, _ := ioutil.ReadFile(filename)
	fmt.Println(string(buf))

	dirname := "/Users/piaoweixiong/"

	list, _ := ioutil.ReadDir(dirname)

	fmt.Println(list[0].IsDir())
	/*var wd []byte
	var wg sync.WaitGroup

	wg.Add(2)

	wd = []byte("你是我的眼")

	pr, pw := io.Pipe()
	rd := make([]byte, 11)

	go pw.Write(wd)
	pr.Read(rd)

	fmt.Println(rd, pr)*/
}
