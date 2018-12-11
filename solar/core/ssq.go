package core

import "fmt"

const (
	REDONE  = "02,07,08,10,12,23"
	BLUEONE = "08"

	REDTWO  = "04,10,17,18,27,31"
	BLUETWO = "02"
)

type Ssq struct {
	State   int      `json:"state"`
	Message string   `json:"message"`
	Result  []result `json:"result"`
}

type result struct {
	Name string `json:"name"`
	Code string `json:"code"`
	Date string `json:"date"`
	Red  string `json:"red"`
	Blue string `json:"blue"`
}

func (s Ssq) Valid() {
	red := s.Result[0].Red
	blue := s.Result[0].Blue

	if red == REDONE && blue == BLUEONE {
		fmt.Println("双色球中大奖了，你还不赶紧查查！")
	} else if red == REDTWO && blue == BLUETWO {
		fmt.Println("双色球中大奖了，你还不赶紧查查！")
	} else {
		fmt.Println("双色球这次没有种大奖，下次继续努力把！")
	}

	fmt.Println("双色球的", "红球：", red, "蓝球：", blue)
}
