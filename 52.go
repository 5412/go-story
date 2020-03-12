package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"
)

func main() {
	for i:=0;i<10;i++ {
		r := rand.Int63n(10)
		fmt.Println(r)
	}
	count := 0
	map2 := make(map[string]int64, 100000)
	startIndexTime := time.Now()
	for i := int64(50000000000000000); i < 50000000000100000; i = i + 2 {
		code := Transform((i), 14)
		if ok := map2[code]; ok > 0 {
			fmt.Println(code, i, map2[code])
			break
		}
		map2[code] = i
		count++
	}
	endIndexTime := time.Now()
	duration := float64(1.0*(endIndexTime.UnixNano()-startIndexTime.UnixNano())) / float64(time.Second)
	log.Println("create index duration :", duration)
	fmt.Println(count, len(map2))

}

func Transform(num int64, length int) string {
	var alphabet = []byte{65, 66, 67, 68, 69, 70, 71, 72, 73, 74, 75, 76, 77, 78, 79, 80, 81, 82, 83, 84, 85, 86, 87, 88, 89, 90, 97, 98, 99, 100, 101, 102}
	in := num
	res := make([]byte, 0)
	for num >= 0 {
		bt := alphabet[num&31]
		res = append(res, bt)
		num >>= 5
		if num == 0 {
			break
		}
	}
	log.Println(string(res), in)
	for len(res) < length {
		res = rangeNum(res)
	}
	log.Println(string(res), in)
	return string(res)
}

func rangeNum(res []byte) []byte {
	r := rand.Int63n(10)
	temp := res[r]
	res[r] = byte(r) + '0'
	res = append(res, temp)
	return res
}

func init() {
	file := "./" + "record" + ".txt"

	logFile, err := os.OpenFile(file, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0766)
	if err != nil {
		panic(err)
	}

	log.SetOutput(logFile) // 将文件设置为log输出的文件
}