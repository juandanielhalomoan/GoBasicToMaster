package main

import (
	"flag"
	"fmt"
)

func main() {
	//var nama_variabel *variabel_pointer = flag.typeData("nama_spesific", "default", "deskripsi")

	//nama flag adalah host
	//defaultnya adalah localhost
	//deskripsinya adalah Put your database host
	//var nama_variabel *variabel_pointer = flag.typeData("nama_spesific", "default", "deskripsi")
	/* kode pertama
	host := flag.String("host", "localhost", "Put your database host")
	username := flag.String("username", "root", "Put your database username")
	password := flag.String("password", "root", "Put your database password")
	port := flag.Int("port",80,"Database Port")
	flag.Parse()
	fmt.Println(*host, *username, *password, *port)
	*/

	//kode kedua
	var host *string = flag.String("host", "localhost", "Put your database host")
	var username *string = flag.String("username", "root", "Put your database username")
	var password *string = flag.String("password", "password", "Put your database password")
	var port *int = flag.Int("port", 80, "Database Port")
	//flag.Parse()
	fmt.Println("host", *host)
	fmt.Println("username", *username)
	fmt.Println("password", *password)
	fmt.Println("port", *port)
}

//go run main.go -username=juan -host=donjuan -password=123456 -port=123

//kita bisa gunakan kutip dua dipassword jika password mengandung space atau lebih dari 1 kata, ini bisa diterapkan diusername atau host dll. contohnya seperti berikut ini

//go run main.go -username=juan -host=donjuan -password="123456 rahasia" -port=123
//expected result
/*host donjuan
username juan
password 123456 5
port 123 */

//misalkan dimatikan/comment //flag.Parse() dan dijalankan perintah berikut ini
//go run main.go -username=juan -host=donjuan -password="123456 rahasia" -port=123
//expected result nya adalah default
/*host localhost
username root
password password
port 80*/
