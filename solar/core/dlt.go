package core

import (
	"fmt"
	"strings"
)

const (
	REDTHREE  = "05 06 11 13 27"
	BLUETHREE = "02 08"
	REDFOUR   = "02 07 12 23 31"
	BLUEFOUR  = "08 10"
)

type Lo struct {
	Lot Lottery `json:"lottery"`
}

type Lottery struct {
	//CodeNumber string `json: "codeNumber"`
	Number   string `json:"number"`
	OpenTime string `json:"openTime_fmt"`
}

func (l Lottery) Valid() {
	result := strings.Split(l.Number, "-")
	red := result[0]
	blue := result[1]
	if red == REDTHREE && blue == BLUETHREE {
		fmt.Println("大乐透中大奖了，你还不赶紧查查！")
	} else if red == REDFOUR && blue == BLUEFOUR {
		fmt.Println("大乐透中大奖了，你还不赶紧查查！")
	} else {
		fmt.Println("大乐透这次没有种大奖，下次继续努力把！")
	}
	fmt.Println("大乐透的", "红球：", red, "蓝球：", blue)
}
