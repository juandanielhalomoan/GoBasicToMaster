package main

import (
	"fmt"
	"time"
)

func main() {
	/*
		//function time()
			now := time.Now()
			fmt.Println(now)
			fmt.Println(now.Local())
			//membuat waktu menggunakan date
			utc := time.Date(2022, time.June, 14, 23, 0, 0, 0, time.UTC)
			fmt.Println(utc)

			//misalnya punya string waktu bisa menggunakan parsing
			//RFC3339 adalah sebuah format baku
			parse, _ := time.Parse(time.RFC3339, "2024-01-02T15:04:05Z")
			fmt.Println(parse)
	*/

	//cara kedua
	//Now() menggunakan time zone waktu komputer yang kita gunakan
	var now time.Time = time.Now()
	//sesuai waktu kita
	fmt.Println(now.Local())

	//membuat waktu UTC
	var utc time.Time = time.Date(2009, time.April, 16, 0, 0, 0, 0, time.UTC)
	fmt.Println(utc)
	//utc menggunakan waktu local
	fmt.Println(utc.Local())

	//membuat time menggunakan format RFC33339
	// formatnya : "2006-01-02T15:04:05Z07:00"
	//formating digolang sedikit membingungkan

	//formater := "yyyy-MM-dd HH:mm:ss"
	formater := "2006-01-02 15:04:05"

	value := "2022-10-10 10:10:10"
	//conversi value yang formatnya time
	valueTime, err := time.Parse(formater, value)
	if err != nil {
		fmt.Println("Error", err.Error())
	} else {
		fmt.Println("success ", valueTime)
	}

	fmt.Println(valueTime.Year())
	fmt.Println(valueTime.Month())
	fmt.Println(valueTime.Day())
	fmt.Println(valueTime.Hour())
	fmt.Println(valueTime.Minute())

	//function Duration
	//duration bentuknya adalah nanosecond
	var duration1 time.Duration = time.Second * 100
	var duration2 time.Duration = time.Millisecond * 10
	var duration3 time.Duration = duration1 - duration2
	fmt.Println(duration1, duration2, duration3) //bentuknya dalam nanosecond
	fmt.Printf("duration3: %d\n", duration3)
	//duration banyak banget tipe contohnya seperti koneksi kedalam database,berapa lama terhubungnya/menunggunya

}
