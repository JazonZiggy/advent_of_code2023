package main

import (
	"fmt"
	openfile "utils"
)

func main() {

	conteudo := openfile.Scan_file("example")
	for i, v := range conteudo {
		fmt.Printf("i: %v, v: %v\n", i, v)
	}
}
