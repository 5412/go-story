package main

import (
	"study/fmt"
)

func main()  {
	/*var ru []rune = []rune("我是字符串哈哈哈12a")
	var bt []byte = []byte("我是字符串哈哈哈12a")

	fmt.Println("rune len:", len(ru))
	fmt.Println("byte len:", len(bt))


	for i, v := range ru  {
		fmt.Println("rune index:", i, "value:", v)
	}

	for i, v := range bt  {
		fmt.Println("byte index:", i, "value:", v)
	}*/

	//fmt.Print("asdas","sda",1,2)
	//fmt.Printf("The format string is : %#v", "i'm string")

	//var num int = '朴' - '伟'
	//fmt.Println('朴')

	//fmt.Printf("The format string is : %#v", "i'm string")
	//fmt.Printf("The format result is: %[2]*.[3]*[1]f", 12.3456789, 1 , 2)
	//fmt.Printf("The format result is: %[2]2d", 128,2)
	//fmt.Printf("The format result is: %d", -10)
	//fmt.Printf("The format result is : %X \n", "雄")
	//fmt.Printf("汉字雄hex %x \n", "雄")
	//fmt.Printf("汉字雄hex %x \n", '雄')
	//fmt.Printf("%d \n", '雄')
	//fmt.Printf("The format result is : %#f \n", 12.345)
	var name string
	var name2 string

	fmt.Println("Please tell me your name.")

	fmt.Scanln(&name, name2)


	fmt.Println("Your name is", name)
	fmt.Println("Your name is", name2)

}