package main

import (
	"fmt"
	"sync"
	"time"
	"math/rand"

)

const (
	numberGoroutines = 4 //要使用的goroutine 的数量
	taskLoad         = 10 //要处理的工作的数量
)

// wg用来等待程序执行完成

var wg sync.WaitGroup

func init()  {
	// 初始化随机数种子
	rand.Seed(time.Now().Unix())
}

func main()  {
	// 创建一个有缓冲的通道来管理工作
	tasks := make(chan string, taskLoad)

	wg.Add(numberGoroutines)
	// 启动goroutine 来处理工作
	for gr := 1; gr <= numberGoroutines; gr++ {
		go worker(tasks, gr)
	}

	// 增加一组需要完成的任务
	for post := 1; post <= taskLoad; post++{
		tasks <- fmt.Sprintf("Task : %d", post)
	}

	// 当所有工作都完成时关闭通道
	// 以便所有goroutine 退出
	close(tasks)

	// 等待所有工作完成
	wg.Wait()
}

func worker(tasks chan string, worker int)  {
	defer wg.Done()

	for {
		task, ok := <-tasks
		if !ok {
			// 这里意味着通道里面没有数据，并且已经被关闭
			fmt.Printf("Worker: %d : Shutting Down\n", worker)
			return
		}

		// 显示我们开始工作了
		fmt.Printf("Worker: %d : Started %s\n", worker, task)

		//随机等一段时间模拟工作
		sleep := rand.Int63n(100)
		time.Sleep(time.Duration(sleep) * time.Millisecond)
		fmt.Printf("Worker : %d : Completed %s\n", worker, task)
	}
}
