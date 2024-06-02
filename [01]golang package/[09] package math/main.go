package main

import (
	"fmt"
	"math"
	"reflect"
)

func main() {

	var nilai_float float64 = 44.44
	fmt.Println(math.Abs(nilai_float))
	fmt.Println(math.Floor(nilai_float))

	// cara konversi dari int ke float
	var nilaiInt int = 55
	var nilaiFloat float64 = float64(nilaiInt)
	fmt.Println(nilaiFloat)
	nilaiFloat += 0.5
	fmt.Println(nilaiFloat)
	fmt.Println(math.Floor(nilaiFloat))
	fmt.Println(reflect.TypeOf(nilaiFloat)) //check type data menggunakan reflect.typeOf

	//kerjakan dengan task berikut
	fmt.Println(math.Ceil(15.6))    //ceil
	fmt.Println(math.Floor(15.6))   //floor
	fmt.Println(math.Round(15.6))   //round
	fmt.Println(math.Max(15.6, 16)) //max
	fmt.Println(math.Min(15.6, 16)) //min

}
