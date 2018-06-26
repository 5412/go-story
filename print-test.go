package main

import (
	"fmt"
	"unicode/utf8"
)
// Use simple []byte instead of bytes.Buffer to avoid large dependency.
type buffer []byte

func (b *buffer) Write(p []byte) {
	*b = append(*b, p...)
}

func (b *buffer) WriteString(s string) {
	*b = append(*b, s...)
}

func (b *buffer) WriteByte(c byte) {
	*b = append(*b, c)
}

func (bp *buffer) WriteRune(r rune) {
	if r < utf8.RuneSelf {
		*bp = append(*bp, byte(r))
		return
	}

	b := *bp
	n := len(b)
	for n+utf8.UTFMax > cap(b) {
		b = append(b, 0)
	}
	w := utf8.EncodeRune(b[n:n+utf8.UTFMax], r)
	*bp = b[:n+w]
}

func main()  {
	fmt.Println("Testing slice!")
	slio := make([]int, 2, 4)
	sli1 := slio[0:1]
	sli3 := slio[1:3]
	fmt.Println(slio);
	fmt.Println(sli1);
	sli3[0] = 2
	sli3 = append(sli3, 3)
	fmt.Println(sli3);
	fmt.Println(slio[0:4]);


	/*var ru rune = 345
	var bt buffer = []byte{0,0}
	fmt.Println("byte ", byte(ru))
	//var bt buffer = make([]byte, 2, 4)
	fmt.Println("The rune is :", ru)
	fmt.Println("RunSelf is :", utf8.RuneSelf)

	b := bt

	n := len(bt)
	fmt.Println("The bt is :", bt)

	fmt.Println("The b is :", b)
	fmt.Println("The cap is :", cap(b))
	for n+utf8.UTFMax > cap(b) {
		b = append(b, 0)
	}
	fmt.Println("The cap is :", cap(b))

	fmt.Println("The b is :", b)
	w := utf8.EncodeRune(b[n:n+utf8.UTFMax], ru)
	fmt.Println("The w is :", w)
	fmt.Println("The b is :", b)
	fmt.Println("n+w is :", b[n:n+utf8.UTFMax])
	bt = b[0:8]
	fmt.Println("The b is :", b)
	fmt.Println("The last bt is :", bt)

	b = append(b, byte(2))
	fmt.Println("The b is :", b)
	fmt.Println("The last bt is :", bt)*/



	/*fmt.Println("Testing print package buffer type!")
	var b buffer
	var bt []byte = []byte{1, '1', 2, 3, 99, 'c', 'd'}
	var str string  = "i'm string"
	var be byte  = 'g'
	var ru rune = 's'

	fmt.Println("The default buffer is:", b)
	fmt.Println("The default []byte is:", bt)
	fmt.Println("The default string is:", str)

	//var nbt []byte = []byte{'c', 'r'}
	//var nstr string = "i'm new string"
	b.Write(bt)
	fmt.Println("After b.Write(bt) b's value is:", b)
	b.WriteString(str)
	fmt.Println("After b.WriteString(str) b's value is :", b)
	b.WriteByte(be)
	fmt.Println("After b.WriteByte b's value is :", b)
	b.WriteRune(ru)
	fmt.Println("After b.WriteRune b's value is :",b)

	fmt.Println("the result string of b is :", string(b))
	fmt.Println(string(ru))
	fmt.Println([]byte(string(b)))*/
}
