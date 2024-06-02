package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	/*
		//code program algoritma Base64
		var encoded = base64.StdEncoding.EncodeToString([]byte("Juan daniel Halomoan"))
		fmt.Println(encoded)
		//konversi menjadi nilai string semula
		var decoded, err = base64.StdEncoding.DecodeString(encoded)
		if err != nil {
			fmt.Println(err.Error())
		} else {
			fmt.Println(decoded)
			fmt.Println(string(decoded))
		}*/

	//code program algoritma CSV Reader (membaca CSV)
	//CSV reader biasa digunakan di microsoft excel saat kita simpan tabelnya didalam sebuah text file
	//maka kita akan simpan didalam satu baris itu row lalu kita akan koma sebagai pemisahnya
	csvString := "juan, daniel halomoan\n" + "adang, pirma \n" + "naomi\n"
	reader := csv.NewReader(strings.NewReader(csvString))
	fmt.Println(reader)
	for {
		record, err := reader.Read() // reader.Read() membaca 1 baris langsung disimpan ke dalam variabel record
		if err == io.EOF {           //EOF adalah end of file artinya sudah berada dibaris terakhir
			break
		}
		fmt.Println(record)
	}
	// baca programnya:
	//ketika read 1 baris nanti dapat akan disimpan di record dan menjadi slice dan membaca jadi csv
	//=========================================

	//code program algoritma CSV Writer (menulis CSV)
	writer := csv.NewWriter(os.Stdout)
	//1 kali write 1 baris bentuknya csv format
	_ = writer.Write([]string{"juan", "daniel", "halomoan"})
	_ = writer.Write([]string{"ganteng", "banget", "sih"})
	_ = writer.Write([]string{"semoga", "dapat", "pekerjaan"})
	writer.Flush() //Flush() mengirim semua perubahannya

}
