package main

import (
	"container/list"
	"fmt"
)

func main() {
	/*//list
	data := list.New()
	data.PushBack("juan")
	data.PushBack("daniel")
	data.PushBack("halomoan")
	// mengambil data paling depan menggunakan Front()
	// kemudian mengambil data selanjutnya menggunakan Next() dari hasil elementnya
	for e := data.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)

	}*/

	var data *list.List = list.New()
	//pushBack() memiliki paramater any yang artinya bisa memasukan data apa saja
	data.PushBack("juan")
	data.PushBack("halomoan")
	data.PushBack("daniel")
	//mengambil data yang pertama menggunakan Front()
	// return Front() adalah pointer list.Element yang paling depan
	var head *list.Element = data.Front()
	//mengambil data selanjutnya menggunakan Next jika ingin mengambil lagi gunakan head.Next().Next() sampai seterusnya
	//sampai nextnya itu adalah pointer dia bisa bernilai nil, jika ia bernilai nil maka ia tidak memiliki nilainya lagi
	fmt.Println(head.Value) //juan

	next := head.Next()
	fmt.Println(next.Value) //halomoan
	//kemudian kita ambil lagi
	next = next.Next() //daniel
	fmt.Println(next.Value)

	//menggunakan looping
	//1.kita ambil dulu data headnya
	//2. elementnya kita ambil dari data.Front()
	//3. lalu kita cek kalo elementnya != nil(selama masih ada elementnya) elementya kita ambil Next()

	for e := data.Front(); e != nil; e = e.Next() {
		fmt.Println("ini merupakan element menggunakan loop: ", e.Value)
	}
}
