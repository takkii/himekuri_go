package main

import (
	"fmt"
	"time"
)

func main() {
	// string keyword
	const bal = "\n"
	const comma = " : "
	
	// time
	t := time.Now()
	
	fmt.Printf("時刻を表示" + comma + "%04d年%02d月%02d日" + comma +"%02d時%02d分%02d秒",
	t.Year(),
	t.Month(),
	t.Day(),
	t.Hour(),
	t.Minute(),
	t.Second())

switch t.Weekday(){
	case time.Sunday:
	fmt.Printf(comma + "日曜日" + bal)
	case time.Monday:
	fmt.Printf(comma + "月曜日" + bal)
	case time.Tuesday:
	fmt.Printf(comma + "火曜日" + bal)
	case time.Wednesday:
	fmt.Printf(comma + "水曜日" + bal)
	case time.Thursday:
	fmt.Printf(comma + "木曜日" + bal)
	case time.Friday:
	fmt.Printf(comma + "金曜日" + bal)
	case time.Saturday:
	fmt.Printf(comma + "土曜日" + bal)
}
	// OneYear = 365 + 1 (uru calc) | Oneday = 366 
	const Oneday = 366
	
	fmt.Printf("来年の1月1日まであと" + comma + "%02d日です" + bal,
		Oneday - t.YearDay())
	
	fmt.Printf("令和%1d年%02d月%02d日" + comma + "R%02d年%02d月%02d日" + bal,
		t.Year() - 2018,
		t.Month(),
		t.Day(),
		t.Year() - 2018,
		t.Month(),
		t.Day())
	
	// version info
	const version = "1.0.0.1"
	
	fmt.Printf("日めくり数え番号" + comma + version + bal)
}