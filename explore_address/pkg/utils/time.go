package utils

import "time"

func TimeNowVietNamString() string {
	location, _ := time.LoadLocation("Asia/Bangkok")
	time := time.Now().In(location)
	return time.Format("2006-01-02T15:04:05.000Z")
}

func TimeNowVietNam() time.Time {
	location, _ := time.LoadLocation("Asia/Bangkok")
	time := time.Now().In(location)
	return time
}
