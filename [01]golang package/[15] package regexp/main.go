package main

import (
	"fmt"
	"regexp"
)

func main() {
	/* contoh kode pertama
	//[a-z] artinya karakter boleh dari a-z diakhiri dengan l
	regex := regexp.MustCompile(`j([a-z])n`)
	fmt.Println(regex.MatchString("juan"))
	fmt.Println(regex.MatchString("daniel"))
	fmt.Println(regex.MatchString("halomoan"))

	fmt.Println(regex.FindAllString("emang ea dani danie daniel juan", 10))
	*/

	//contoh kode kedua
	//var regex *regexp.Regexp = regexp.MustCompile(`e([a-z])o`)
	var regex *regexp.Regexp = regexp.MustCompile(`s([a-z][a-z])o`)

	fmt.Println(regex.MatchString("sko"))
	fmt.Println(regex.MatchString("seeko"))
	fmt.Println(regex.MatchString("seXo"))
	fmt.Println(regex.FindAllString("seko ekopak ekochan ekodher sko seto etoo sekos", 10))
}
