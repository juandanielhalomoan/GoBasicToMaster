package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {

	//code program Reader
	input := strings.NewReader("ini merupakan string panjang\n dengan baris yang barus\n")
	reader := bufio.NewReader(input)

	for {
		line, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		}
		fmt.Println(string(line))
	}

	//code program Writer
	writer := bufio.NewWriter(os.Stdout)
	_, _ = writer.WriteString("hello ini adalah program writer \n")
	_, _ = writer.WriteString("selamat belajar")
	writer.Flush()
}
