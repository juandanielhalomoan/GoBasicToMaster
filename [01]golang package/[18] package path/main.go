package main

import (
	"fmt"
	"path"
	"path/filepath"
)

func main() {
	//code program pertama : package path
	fmt.Println(path.Dir("hello/hai/word.go"))          //mengambil direktori(folder)
	fmt.Println(path.Base("halo/word.go"))              // mengambil nama file akhirnya
	fmt.Println(path.Ext("helo.gs/world.gs"))           //mengambil akhir ekstension
	fmt.Println(path.Join("hello", "itsme", "main.go")) // join beberapa folder
	fmt.Println("")

	//code program kedua : path/filepath
	//kelebihan dari menggunakan filepath adalah nanti dia bisa tau apakah file ini sistem operasinya windows atau bukan
	fmt.Println(filepath.Dir("hello/world.go"))             //nama direktori
	fmt.Println(filepath.Base("hello/world.go"))            //nama filenya
	fmt.Println(filepath.Ext("hello/world.go"))             //nama ektension
	fmt.Println(filepath.IsAbs("hello/world.go"))           //absolutnya: artinya dimulai paling depan, kalo diwindows dari drive C
	fmt.Println(filepath.Join("hello", "world", "main.go")) // bisa baca sistem operasi macam apa saja

}
