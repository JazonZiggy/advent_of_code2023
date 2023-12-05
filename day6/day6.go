package main

import (
	"fmt"
	"utils"
)

func main() {
	teste := openfile.Scan_file("example")

	fmt.Printf("teste: %v\n", teste)
}
