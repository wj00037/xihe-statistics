package utils

import "time"

func TimeStampToUnixTime(stamp string) (t int64, err error) {
	local, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		return
	}

	tt, err := time.ParseInLocation("2006-01-02T15:04:05+08:00", stamp, local)
	if err != nil {
		return
	}

	return tt.Unix(), nil
}
