package openfile

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Scan_file(nameOfFile string) []string {
	file, err := os.Open(nameOfFile)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	fileText := []string{}
	for scanner.Scan() {
		fileText = append(fileText, scanner.Text())
	}
	return fileText
}

func Convert_matrix(conteudo []string) [][]string {
	matriz := make([][]string, 0)
	fmt.Printf("conteudo: %v\n", conteudo)
	for _, linha := range conteudo {
		colunas := strings.Split(linha, "")
		fmt.Printf("colunas: %v\n", colunas)
		matriz = append(matriz, colunas)
	}
	return matriz

}
