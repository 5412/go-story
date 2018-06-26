package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var (
	wg sync.WaitGroup
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	wg.Add(2)
	fmt.Println("Goroutine with no buffer chan")
	court := make(chan int)
	//启动两个选手
	go player("solar", court)
	go player("yujindan", court)

	//发球
	court <- 1
	wg.Wait()
}

func player(name string, court chan int) {
	//函数调用Done来通知main函数工作完成
	defer wg.Done()

	for {
		ball, ok := <-court
		if !ok {
			fmt.Printf("Player %s Won\n", name)
			return
		}

		//选个随机数然后用这个数来模拟我们是否丢球
		n := rand.Intn(100)
		if n%13 == 0 {
			fmt.Printf("Player %s Missed\n", name)
			close(court)
			return
		}

		//显示击球数并将击球数加1

		fmt.Printf("Player %s Hit %d\n", name, ball)
		ball++

		//将球打向对手
		court <- ball
	}
}
