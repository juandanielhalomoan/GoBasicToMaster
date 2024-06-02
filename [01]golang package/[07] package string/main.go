package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("    juan daniel halomoan    ")                                                     //sebelum trim
	fmt.Println(strings.Trim("    juan daniel halomoan    ", " "))                                  //memotong cutset diawal dan akhir string //contohnya ingin menghapus space dikiri dan kanan
	fmt.Println(strings.ToLower("juan daniel halomoan"))                                            //membuat semua karakter string menjadi lower case
	fmt.Println(strings.ToUpper("juan daniel halomoan"))                                            //membuat semua karakter string menjadi upper case
	fmt.Println(strings.Split("juan daniel halomoan", "da"))                                        //memotong string berdasarkan separator dan return slice string
	fmt.Println(strings.Contains("juan daniel halomoan", "halo"))                                   //mengecek apakah string mengandung string lain
	fmt.Println(strings.ReplaceAll("juan daniel halomoan", "juan daniel halomoan", "juan ganteng")) //mengubah semua string dari from - ke to -
}
