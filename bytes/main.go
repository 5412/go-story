package main

func main()  {

	//fmt.Println(bytes.IndexAny([]byte{'a', 'b'}, "ab"))



	/*
	var num uint8  = 15
	fmt.Println(num, num * num, num * num * num)
	*/

	/*
	// Unicode-graphic.go

	const (
		pC     = 1 << iota // a control 控制 character.
		pP                 // a punctuation 标点 character.
		pN                 // a numeral.
		pS                 // a symbolic 符号的 character.
		pZ                 // a spacing 间隔 character.
		pLu                // an upper-case letter.
		pLl                // a lower-case letter.
		pp                 // a printable character according to Go's definition.
		pg     = pp | pZ   // a graphical 图解的 character according to the Unicode definition.
		pLo    = pLl | pLu // a letter that is neither upper nor lower case.
		pLmask = pLo
	)

	fmt.Printf("%b \n", pC)       // 1
	fmt.Printf("%b \n", pP)       // 10
	fmt.Printf("%b \n", pN)       // 100
	fmt.Printf("%b \n", pS)       // 1000
	fmt.Printf("%b \n", pZ)       // 10000
	fmt.Printf("%b \n", pLu)      // 100000
	fmt.Printf("%b \n", pLl)      // 1000000
	fmt.Printf("%b \n", pp)       // 10000000
	fmt.Printf("%b \n", pg)       // 10010000
	fmt.Printf("%b \n", pLo)      // 1100000
	fmt.Printf("%b \n", pLmask)   // 1100000

	str := "1a哈哈哈"
	fmt.Println(str[0:2])

	r := []byte{233,1}
	fmt.Println(utf8.FullRune(r))

	a := make([][]byte, 10)
	fmt.Println(a[0])
	fmt.Println(hashStr([]byte{2,3,4,5,6,7,3,9,4,1,1,2,1,1,1}))

	//fmt.Println(Index([]byte{1,2,3,4,5,6,7,8,9,9,0,8}, []byte{2,3,4,5,6,7,8,9,9}))
	str := "1于大傻子"
	for k,v := range str {
		fmt.Println(k, ":", string(v), string(str[k]))
	}
	*/
}
