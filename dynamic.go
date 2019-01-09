package main

import (
	"fmt"
	"math"
)

type Task struct {
	begin int
	end int
	weight float64
}

type Tasks []Task

var tasks = Tasks{
	1:{1,4,5},
	2:{3,5,1},
	3:{0,6,8},
	4:{4,7,4},
	5:{3,8,6},
	6:{5,9,3},
	7:{6,10,2},
	8:{8,11,4},
}

func prev(i int) int {
	var r int
	t := tasks[i]
	for k,v := range tasks{
		//fmt.Println(k)
		if v.end <= t.begin {
			r = k
		}
	}
	return r
}

func opt(i int) float64  {
	if i == 0 {
		return 0
	}
	t := tasks[i]
	return math.Max(t.weight + opt(prev(i)), opt(i-1))
}

func main()  {
	fmt.Println(opt(6))
}

