package main

import (
    "fmt"
    "sync"
    "runtime"
)

var counter int = 0

func Count(lock *sync.Mutex) {
    lock.Lock()
    counter++
    z := counter
    fmt.Println("Counter now is", z)
    lock.Unlock()
}

func main() {
    lock := &sync.Mutex{} //互斥锁
    fmt.Println("System is ", runtime.GOOS)
    fmt.Println("CPU nums is  ", runtime.NumCPU())
    fmt.Println("Go root is  ", runtime.GOROOT())
    for i := 0; i < 11; i++ {
        go Count(lock)
    }

    for {
        lock.Lock()
        c := counter
        lock.Unlock()
        if c > 10 {
            fmt.Println("finish")
            break
        } else {
            fmt.Println("10 goroutines is not finish yet.", c)
            runtime.Gosched() //让出系统CPU
        }
    }
}
