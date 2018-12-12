package main

import (
	"fmt"
	"time"
)

func main() {
	// i := make(chan int, 10)
	// j := 0
	// for {

	//     fmt.Println("Loop index is ", j)
	//     select {
	//         case i <- 0:
	//             fmt.Println("0 value is <-")
	//         case i <- 1:
	//             fmt.Println("1 value is <-")
	//         default:
	//             fmt.Println("Nothing done")
	//     }
	//     j++

	//     if j > 1024 {
	//         break
	//     }
	// }

	//for c:=0; c < 11; c++ {
	// v, ok :=  <-i
	//if ok {
	//   fmt.Println("Value recieve is ", v)
	//} else {
	//   close(i)
	// fmt.Println("Closing chan")
	//}
	//fmt.Println("Value now is ", v)
	//}
	// for k := range i {
	//      fmt.Println("Recieve chan ", k)
	// }

	// chStr := make(chan string, 1)
	// chInt := make(chan int, 1)

	// go strWorker(chStr)
	// go intWorker(chInt)

	// for i := 0; i < 2; i++ {
	// 	select {
	// 	case <-chStr:
	// 		fmt.Println("get value from strWorker")

	// 	case <-chInt:
	// 		fmt.Println("get value from intWorker")

	// 	}
	// }

	done := make(chan bool, 1)

	go worker(done)

	//<-done
}

func worker(done chan bool) {
	fmt.Println("start working...")
	done <- true
	fmt.Println("end working...")
}

func strWorker(ch chan string) {
	time.Sleep(1 * time.Second)
	fmt.Println("do something with strWorker...")
	ch <- "str"
}

func intWorker(ch chan int) {
	time.Sleep(2 * time.Second)
	fmt.Println("do something with intWorker...")
	ch <- 1
}
