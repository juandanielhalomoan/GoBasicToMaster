package main

import (
	"fmt"
	"sort"
)

// membuat sebuah type user
type User struct {
	Name string
	Age  int
}

// alias dari type decoration User
// lalu kita akan implementasikan userSlice untuk sebuah kontrak dari sort.Interface
// agar userslice bisa diurutkan dengan gampang
type UserSlice []User

// untuk mau menggunakan package sort kita harus buat sebuah kontrak untuk userSlice
// tidak butuh pointer karena default slice adalah pointer
func (userSlice UserSlice) Len() int {
	return len(userSlice)
}

// buat function kedua yang harus ada adalah less()
// memiliki index i dan index j int dan returnya adalah boolean
// lalu di cek apakah i lebih kecil dari j atau tidak
func (userSlice UserSlice) Less(i, j int) bool {
	return userSlice[i].Age < userSlice[j].Age
}

// melakukan swap pemindahan posisi dari i ke j
// sudah mengikuti interface kontrak
func (userSlice UserSlice) Swap(i, j int) {
	/*//temporary
	temp := userSlice[i]
	userSlice[i] = userSlice[j]
	userSlice[j] = temp */
	//digolang ada cara yang lebih gampang kalo mau menukar posisi jadi tidak butuh temporary
	userSlice[i], userSlice[j] = userSlice[j], userSlice[i]

}

func main() {
	users := []User{
		{"Juan", 25},
		{"daniel", 15},
		{"halomoan", 20},
	}
	//task harapannya mengurutkan umur dari terkecil hingga terbesar
	sort.Sort(UserSlice(users))
	fmt.Println(users)
	//jadi kita harus buat sebuah kontrak terlebih dahulu
}
