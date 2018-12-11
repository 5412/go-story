package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"net/http"
	"solar/core"
	"strconv"
	"io/ioutil"
)

var code string

func main() {
	flag.StringVar(&code, "c", "2018030", "Usage: -c 双色球的期数,2018030是默认期数")
	flag.Parse()
	doSsq(code)
	doDlt(code)
}

func doDlt(code string) {

	num, _ := strconv.Atoi(code)
	code = strconv.Itoa(num - 2000000)
	fmt.Println("大乐透第", code, "期开奖结果")
	url := "http://www.lottery.gov.cn/api/lottery_kj_detail_new.jspx?_ltype=4&_term=" + code
	client := &http.Client{}
	request, err := http.NewRequest("GET", url, nil)
	request.Close = true
	resp, err := client.Do(request)
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
	}
	var data []core.Lo
	json.Unmarshal(body, &data)
	data[0].Lot.Valid()
}

func doSsq(code string) {
	fmt.Println("双色球第", code, "期开奖结果")
	client := &http.Client{}
	url := "http://www.cwl.gov.cn/cwl_admin/kjxx/findKjxx/forIssue?name=ssq&code=" + code
	reqest, err := http.NewRequest("GET", url, nil)
	reqest.Header.Set("Referer", "http://www.cwl.gov.cn/kjxx/ssq/")

	resp, err := client.Do(reqest)

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
	}

	//str := string(buf[0:n])
	var data core.Ssq
	json.Unmarshal(body, &data)
	data.Valid()
}
