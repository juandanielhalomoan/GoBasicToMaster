package main

import (
	"fmt"

	"slices"
)

func main() {
	names := []string{"juan", "daniel", "halomoan", "ganteng", "banget"}
	values := []int{100, 24, 56, 55}
	fmt.Println(names)
	fmt.Println(slices.Min(names))
	fmt.Println(slices.max(values))
}
