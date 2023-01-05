package app

import "time"

var cstSh, _ = time.LoadLocation("Asia/Shanghai")

func getLocalTime() (t string) {
	return time.Now().In(cstSh).Format("2006-01-02T15:04:05+08:00")

}

func getUnixLocalTime() (t int64) {
	return time.Now().In(cstSh).Unix()
}
