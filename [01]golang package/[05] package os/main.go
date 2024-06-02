/*
//kode pertama
package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args
	//interaksi menggunakan for loop
	// variabel pertama di for merupakan sebuah index
	// variabel kedua di for merupakan isi data dari foor
	// kita hiraukan indexnya terlebih dahulu i menjadi _
	for _, arg := range args {
		fmt.Println(arg)
	}
}

//kita bisa jalankan dengan go run main.go paramater-argumen(any)
//contoh go run main.go juan daniel halomoan
//expected output
//C:\Users\junke\AppData\Local\Temp\go-build3075285787\b001\exe\main.exe
//juan
//daniel
//halomoan
//
// hasil compile baris pertama adalah hasil lokasi binary filenya, jadi sebelum dirun dicompile dulu oleh golang secara otomatis lalu disimpan di temporary file
// hasil baris kedua dan seterusnya adalah hasil data dari parameter yang kita inputkan diterminal
//terus bagaimana kita satukan datanya? kita bisa gunakan kutip dua didalam paramater
//contoh go run main.go "juan daniel halomoan"
//expected output
//C:\Users\junke\AppData\Local\Temp\go-build2725401644\b001\exe\main.exe
//juan daniel halomoan

*/

// kode kedua
package main

import (
	"fmt"
	"os"
)

func main() {
	//os.Hostname() memiliki 2 return yaitu string dan error
	hostname, err := os.Hostname()
	if err == nil {
		fmt.Println(hostname)
	} else {
		fmt.Println("Error", err.Error())
	}
}
