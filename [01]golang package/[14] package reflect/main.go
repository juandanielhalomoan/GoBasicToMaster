//program pertama package reflect

/*
package main

import (

	"fmt"
	"reflect"

)

	type Sample struct {
		Name string
	}

	type Person struct {
		Name, Adress, Email string
		Umur                int64
	}

	func ReadField(value any) {
		var valueType reflect.Type = reflect.TypeOf(value)
		fmt.Println("Type Name", valueType.Name())
		//looping
		for i := 0; i < valueType.NumField(); i++ {
			var valueField reflect.StructField = valueType.Field(i)
			fmt.Println(valueField.Name, "with type", valueField.Type)
		}
	}

	func main() {
		/*
			//untuk membaca sebuah format metadata dari sebuah object kita bisa gunakan reflectioning
				//cara pertama
				sample := Sample{"juan"}
				sampleType := reflect.TypeOf(sample)
				structField := sampleType.Field(0)
				fmt.Println(sample, structField)


		//cara kedua
		//membaca filed sample
		ReadField(Sample{"Juan"})
		ReadField(Person{"daniel", "tamer", "junker@gmail.com", 25})
	}
*/
/*
// kode program kedua StructTag
package main

import (
	"fmt"
	"reflect"
)

type Sample struct {
	Name string `required:"true" max:"10"`
}
type Person struct {
	Name   string `required:"juan" umur:"25"`
	Adress string `required:"tamer" umur:"25"`
	Email  string `required:"junker@gmail.com" umur:"25"`
}

func ReadField(value any) {
	var valueType reflect.Type = reflect.TypeOf(value)
	fmt.Println("Type Name", valueType.Name())
	//looping
	for i := 0; i < valueType.NumField(); i++ {
		var structField reflect.StructField = valueType.Field(i)
		fmt.Println(structField.Name, "with type", structField.Type)
		structField.Tag.Get("required")
		structField.Tag.Get("max")

	}
}

func main() {

		//cara pertama
		sample := Sample{"juan"}
		sampleType := reflect.TypeOf(sample)
		structField := sampleType.Field(0)
		required := structField.Tag.Get("max")

		fmt.Println(required)

	//cara kedua
	ReadField(Sample{"juan"})
	ReadField(Person{"umur", "juan", "juan@gmail.com"})

}*/

// kode program ketiga StructTag Validation Library
package main

import (
	"fmt"
	"reflect"
)

type Sample struct {
	Name string `required:"true" max:"10"`
}
type Person struct {
	Name   string `required:"true" umur:"25"`
	Adress string `required:"true" umur:"25"`
	Email  string `required:"true" umur:"25"`
}

func ReadField(value any) {
	var valueType reflect.Type = reflect.TypeOf(value)
	fmt.Println("Type Name", valueType.Name())
	//looping
	for i := 0; i < valueType.NumField(); i++ {
		var structField reflect.StructField = valueType.Field(i)
		fmt.Println(structField.Name, "with type", structField.Type)
		structField.Tag.Get("required")
		structField.Tag.Get("max")

	}
}

// cara 1 validasi return boolean
/*
func IsValid(value any) bool {
	t := reflect.TypeOf(value)
	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		//lakukan validasi
		if f.Tag.Get("required") == "true" {
			//data yang ada didalam struct
			data := reflect.ValueOf(value).Field(i).Interface() // interface adalah datanya
			//kita cek
			return data != ""
		}
	}
	return true
}*/

// cara 2 validasi return boolean
func IsValid(value any) (result bool) {
	result = true
	t := reflect.TypeOf(value)
	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		//lakukan validasi
		if f.Tag.Get("required") == "true" {
			//data yang ada didalam struct
			data := reflect.ValueOf(value).Field(i).Interface() // interface adalah datanya
			result = data != ""
			//kita cek
			if result == false {
				return result
			}
		}
	}
	return result
}

func main() {
	/*
		//cara pertama
		sample := Sample{"juan"}
		sampleType := reflect.TypeOf(sample)
		structField := sampleType.Field(0)
		required := structField.Tag.Get("max")

		fmt.Println(required) */

	//cara kedua
	ReadField(Sample{"juan"})
	ReadField(Person{"umur", "juan", "juan@gmail.com"})
	person := Person{
		Name:   "Eko",
		Adress: "",
		Email:  "s",
	}
	fmt.Println(IsValid(person))

}
