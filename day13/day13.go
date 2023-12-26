package main

import (
	"fmt"
	"strings"
	openfile "utils"
)

var smudge int = 0
var checkSmudge bool = true
var newLine []string = []string{}

func compareTwoLines(l1, l2 []string) bool {

	linha1 := strings.Join(l1, "")
	linha2 := strings.Join(l2, "")
	equals := 0
	for i, v := range l1 {
		if v == l2[i] {
			equals++
		}
	}
	if len(l1)-1 == equals {
		fmt.Printf("l1: %v\n", l1)
		fmt.Printf("l2: %v\n", l2)
		if smudge == 0 && checkSmudge {
			println("entrei?")
			smudge++
			newLine = l1
			return true
		}
	}

	return linha1 == linha2

}

func checkHorizontal(matrix [][]string, start int, run int) int {

	line1 := start + run + 0
	line2 := start - run + 1

	if line1 >= len(matrix) || line2 == -1 {
		return 0
	}
	if run == 1 {
		smudge = 0
	}

	if compareTwoLines(matrix[line1], matrix[line2]) {
		if line1 == len(matrix)-1 || line2 == 0 {
			if smudge == 1 {

				checkSmudge = false
				actualScore := start + 1
				matrix[line1] = newLine
				fmt.Printf("matrix: %v\n", matrix)
				return checkVertical(matrix, 0, 1) + actualScore*100
			}
			if !checkSmudge {
				return (start + 1) * 100
			}
		}
		result := checkHorizontal(matrix, start, run+1)
		return result

	}

	return checkHorizontal(matrix, start+1, 1)
}
func createNewMatrixWithRow(matrix [][]string, newLine []string, where int) [][]string {
	fmt.Printf("newLine: %v\n", newLine)
	fmt.Printf("matrix: %v\n", matrix)
	for i, _ := range matrix {
		matrix[i][where] = newLine[i]
	}
	fmt.Printf(" NEW matrix: %v\n", matrix)
	return matrix
}

func checkVertical(matrix [][]string, start int, run int) int {

	line1 := start + run + 0
	line2 := start - run + 1
	if line1 >= len(matrix[0]) || line2 == -1 {
		return 0
	}

	if run == 1 {
		smudge = 0
	}

	lineVertical1 := []string{}
	lineVertical2 := []string{}
	for i, _ := range matrix {
		lineVertical1 = append(lineVertical1, matrix[i][line1])
		lineVertical2 = append(lineVertical2, matrix[i][line2])
	}
	if compareTwoLines(lineVertical1, lineVertical2) {
		if line1 == len(matrix[0])-1 || line2 == 0 {
			if smudge == 1 {

				checkSmudge = false
				actualScore := start + 1
				newMatrix := createNewMatrixWithRow(matrix, newLine, line2)
				return 100*checkHorizontal(newMatrix, 0, 1) + actualScore
			}
			for _, v := range matrix {
				fmt.Printf("v: %v\n", v)
			}
			if !checkSmudge {
				return start + 1
			}
		}
		result := checkVertical(matrix, start, run+1)
		return result

	}

	return checkVertical(matrix, start+1, 1)
}

func main() {

	conteudo := openfile.Scan_file("example")
	matrix := openfile.Convert_matrix(conteudo)
	contador := 0
	score := 0
	for i, v := range matrix {
		if len(v) == 0 {
			checkSmudge = true
			scoreV := checkVertical(matrix[contador:i], 0, 1)
			scoreH := 0
			if scoreV == 0 {
				checkSmudge = true
				scoreH = checkHorizontal(matrix[contador:i], 0, 1)
			}
			fmt.Printf("smudge: %v\n", smudge)
			fmt.Printf("scoreV: %v\n", scoreV)
			fmt.Printf("scoreH: %v\n", scoreH)
			score += scoreV + scoreH
			contador = i + 1
		}
		if len(matrix)-1 == i {
			scoreH := 0
			checkSmudge = true
			scoreV := checkVertical(matrix[contador:], 0, 1)
			if scoreV == 0 {
				checkSmudge = true
				scoreH = checkHorizontal(matrix[contador:], 0, 1)
			}
			score += scoreH + scoreV
			fmt.Printf("scoreH: %v\n", scoreH)
			fmt.Printf("scoreV: %v\n", scoreV)
		}
	}
	fmt.Printf("score: %v\n", score)
}
