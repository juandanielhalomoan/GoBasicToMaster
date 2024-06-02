package main

import (
	"container/ring"
	"fmt"
	"strconv"
)

func main() {
	/*
		//contohnya kita punya 5 data, berarti len()/panjang ada 5
		data := ring.New(5)
		for i := 0; i <= data.Len(); i++ {
			data.Value = "Value " + strconv.FormatInt(int64(i), 10)
			//ambil data menggunakan Next()
			data = data.Next()
		}
		//melakukan sesuatu dari datanya dengan menggunakan Do()
		//memiliki sebuah function yang datanya adalah any
		data.Do(func(value interface{}) {
			fmt.Println(value)
		})

		//cara manual
		// ring.new() return pointer ring
		//data default
		var dataRing *ring.Ring = ring.New(4)
		dataRing.Value = "Value 1"
		dataRing = dataRing.Next()
		dataRing.Value = "Value 2"
		dataRing = dataRing.Next()
		dataRing.Value = "Value 3"
		dataRing = dataRing.Next()
		dataRing.Value = "Value 4"
		dataRing = dataRing.Next()
		//fmt.Println(dataRing.Value)
		//lalu kita bisa bikin iterasi menggunakan Do()
		//dan buat sebuah function
		dataRing.Do(func(value any) {
			fmt.Println(value)
		})

	*/
	//cara otomatis
	var data *ring.Ring = ring.New(5)
	for i := 0; i < data.Len(); i++ {
		data.Value = "Value " + strconv.Itoa(i) //kita gak bisa tambah i karena integer jadi kita ubah ke string
		data = data.Next()
	}
	data.Do(func(value any) {
		fmt.Println(value)
	})
}
