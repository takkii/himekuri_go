package main

import (
	"fmt"
	"time"
)

func main() {
	const bal = "\n"
	const comma = " : "
	
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

	fmt.Printf("来年の1月1日まであと" + comma + "%02d日です" + bal,
		366 - t.YearDay())
	
	fmt.Printf("令和%1d年%02d月%02d日" + comma + "R%02d年%02d月%02d日" + bal,
		t.Year() - 2018,
		t.Month(),
		t.Day(),
	    t.Year() - 2018,
	    t.Month(),
	    t.Day())
	
	const version = "1.0.0"
	fmt.Printf("日めくり数え番号" + comma + version + "\n")
}