package main

import (
	"errors"
	"fmt"
)

var (
	ValidationError = errors.New("validation error")
	NotFoundError   = errors.New("not found error")
)

func GetById(id string) error {
	if id == "" {
		return ValidationError
	}
	if id != "juan" {
		return NotFoundError
	}
	return nil
}

func main() {
	err := GetById("juan")
	//mengecek jenis error
	//misal kita membuat jenis error sendiri, lalu kita ingin mengecek jenis errornya
	//kita bisa menggunakan error.is() untuk mengecek jenis type errornya
	if err != nil {
		if errors.Is(err, ValidationError) {
			fmt.Println("validation error")
		} else if errors.Is(err, NotFoundError) {
			fmt.Println("not found error")
		} else {
			fmt.Println("uknown error")
		}
	}
}
