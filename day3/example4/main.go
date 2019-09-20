package main

import (
	"fmt"
	"time"
)

const (
	Nanosecond  time.Duration = 1
	Microsecond               = 1000 * Nanosecond
	Millisecond               = 1000 * Microsecond
	Second                    = 1000 * Millisecond
	Minute                    = 60 * Second
	Hour                      = 60 * Minute
)

func main() {
	now := time.Now()
	fmt.Println("当前时间：", now)
	fmt.Println("年：", now.Year())
	fmt.Println("月：", now.Month(), int(now.Month())) // 默认获取的月份是英文，用int转换后可以得到月份的数字类型
	fmt.Println("日：", now.Day())
	fmt.Println("时：", now.Hour())
	fmt.Println("分：", now.Minute())
	fmt.Println("秒：", now.Second())

	Nanosecond := now.Nanosecond()
	Microsecond := Nanosecond / 1000
	Millisecond := Microsecond / 1000
	fmt.Println("纳秒：", Nanosecond)
	fmt.Println("微秒：", Microsecond)
	fmt.Println("微秒：", Millisecond)

	// 获取当前时间时间戳
	timeStamp := now.Unix()
	fmt.Println("当前时间戳", timeStamp)

	// str格式化时间
	fmt.Println(now.Format("2006/01/02 15:04:05")) // 必须是这个时间点, 记忆方法:6-1-2-3-4-5
	fmt.Println(now.Format("2006/1/02 15:04"))
	fmt.Println(now.Format("02/1/2006 15:04:05"))

	// 时间戳转str格式化时间
	strTime := time.Unix(1568859265, 0).Format("2006-01-02 15:04:05")
	fmt.Println(strTime)

	// str格式化时间转时间戳
	// 方式1
	theTime := time.Date(2019, 9, 9, 9, 9, 9, 0, time.Local)
	unixTime := theTime.Unix()
	fmt.Println("Local timeStamp", unixTime)
	// 方式2
	theUTCTime, err := time.Parse("2006-01-02 15:04:05", "2019-09-09 09:09:09")
	fmt.Println("theUTCTime:", theUTCTime)
	theLocalTime, err := time.ParseInLocation("2006-01-02 15:04:05", "2019-09-09 09:09:09", time.Local)
	fmt.Println("theLocalTime:", theLocalTime)
	if err == nil {
		utcTime := theUTCTime.Unix()
		localTime := theLocalTime.Unix()
		fmt.Println("UTC timeStamp:", utcTime)
		fmt.Println("Local timeStamp:", localTime)
	}
}
