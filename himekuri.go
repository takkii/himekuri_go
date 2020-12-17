package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	
	fmt.Printf("時刻を表示 : %04d年%02d月%02d日 : %02d時%02d分%02d秒",
	t.Year(),
	t.Month(),
	t.Day(),
	t.Hour(),
	t.Minute(),
	t.Second())

switch t.Weekday(){
	case time.Sunday:
	fmt.Printf(" : 日曜日\n")
	case time.Monday:
	fmt.Printf(" : 月曜日\n")
	case time.Tuesday:
	fmt.Printf(" : 火曜日\n")
	case time.Wednesday:
	fmt.Printf(" : 水曜日\n")
	case time.Thursday:
	fmt.Printf(" : 木曜日\n")
	case time.Friday:
	fmt.Printf(" : 金曜日\n")
	case time.Saturday:
	fmt.Printf(" : 土曜日\n")
}

	fmt.Printf("来年の1月1日まであと : %02d日です\n",
		366 - t.YearDay())
	
	fmt.Printf("令和%02d年%02d月%02d日 : R%02d年%02d月%02d日\n",
		t.Year() - 2018,
		t.Month(),
		t.Day(),
	    t.Year() - 2018,
	    t.Month(),
	    t.Day())
	
	const version = "1.0.0"
	const comma = " : "
	fmt.Printf("日めくり数え番号" + comma + version + "\n")
}