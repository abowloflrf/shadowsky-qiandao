package main

import (
	"regexp"
	"strconv"
)

// 成功： {Msg:获得了 259 MB流量. Ret:1}
// 已签到： {Msg:您似乎已经签到过了... Ret:1}

type CheckinResult struct {
	Msg string `json:"msg"`
	Ret int    `json:"ret"`
}

// DataN parse the checkin response message and get the number of free data in MB just got
func (cr *CheckinResult) DataN() int {
	compRegex := regexp.MustCompile(`^获得了 (\d+) MB流量$`)
	res := compRegex.FindStringSubmatch(cr.Msg)
	if len(res) != 2 {
		return 0
	}
	n, err := strconv.Atoi(res[1])
	if err != nil {
		return 0
	}
	return n
}
