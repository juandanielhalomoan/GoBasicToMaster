package main

import (
	"fmt"
	"strconv"
)

func main() {
	/*
		fmt.Println(strconv.parseBool("true"))  //mengubah string ke boolean
		fmt.Println(strconv.parseFloat("14.0")) //mengubah string ke float
		fmt.Println(strconv.parseInt("14"))     //mengubah string ke int
		fmt.Println(strconv.FormatBool(true))   //mengubah bool ke string
		fmt.Println(strconv.FormatFloat(16.0))  //mengubah float ke string
		fmt.Println(strconv.FormatInt(131))     //mengubah int64 ke string*/

	resultBoolean, err := strconv.ParseBool("false") //ini merupakan parsebool
	if err == nil {
		fmt.Println(resultBoolean)
	} else {
		fmt.Println("Error", err.Error())
	}

	resultFloat, err := strconv.ParseFloat("16.58", 64) //ini merupakan parseFloat
	//jika kita milih 32 outpuntya akan menyertakan bilangan desimal
	//jika kita milih 64 outputnya akan dibulatkan ke atas
	if err != nil {
		fmt.Println("Error", err.Error())
	} else {
		fmt.Println(resultFloat)
	}

	resultInt, err := strconv.Atoi("100000000") //ini merupakan parse int menggunakan Atoi
	// atoi sering digunakan untuk konversi string ke int dan tidak menggunakan baseint makanya sering digunakan
	if err == nil {
		fmt.Println(resultInt)
	} else {
		fmt.Println("Error", err.Error())
	}

	//jika menggunakan converstion binary
	binary := strconv.FormatInt(999, 36)
	fmt.Println(binary)

	//mau lebih sederhana menggunakan strconv.Itoa
	//hasilnya adalah string 999
	var itoa string = strconv.Itoa(999)
	fmt.Println(itoa)
}
