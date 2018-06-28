// Copyright (c) 2018 The Solar. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"solar/fmt"
	"strings"
)

func main() {

	//r := makeGenericReplacer([]string{"bc", "MNV", "bc", "ABC"})
	//fmt.Println(r.root.table[0].next.value)
	str := "piao朴考"

	/*
	fmt.Println(strings.Title("hEllo woRld"))
	fmt.Println(strings.Fields("hello world pwx"))
	fmt.Println(strings.Split("hello word pwx", "w"))
	fmt.Println(strings.SplitAfter("1h2h3h4h5h", "h"))
	fmt.Println(strings.Split("h1h2h3h4h5", "h"))

	//str = "sss"

	fmt.Println([]byte(str))
	fmt.Println([]rune(str))
	for i,v := range(str) {
		fmt.Println(i, v)
	}
	fmt.Println(str[4])
	fmt.Println(string(unicode.SimpleFold('m')))
	fmt.Println(func(a ...string) []string {
		return  a
	}(str,str))
	*/

	str = strings.NewReplacer("piao", "朴", "朴", "piao").Replace(str)
	fmt.Println(str)

}