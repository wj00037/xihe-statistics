package app

import "time"

var (
	cstSh, _   = time.LoadLocation("Asia/Shanghai")
	timeFormat = "2006-01-02T15:04:05+08:00"
)

func getLocalTime() (t string) {
	return time.Now().In(cstSh).Format(timeFormat)

}

func getUnixLocalTime() (t int64) {
	return time.Now().In(cstSh).Unix()
}

func toTimeStamp(t string) (stamp time.Time, err error) {
	stamp, err = time.ParseInLocation(timeFormat, t, cstSh)
	if err != nil {
		return
	}

	return
}

func toUnixTime(t string) (tt int64, err error) {
	stamp, err := toTimeStamp(t)
	if err != nil {
		return
	}

	tt = stamp.Unix()

	return
}
