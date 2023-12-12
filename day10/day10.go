package main

import (
	"fmt"
	"strings"
	openfile "utils"
)

func findSLocation(S string, matrix [][]string) (int, int) {
	for i, row := range matrix {
		for j, char := range row {
			if char == S {
				return i, j
			}
		}
	}
	return -1, -1
}

var matrix [][]string
var count int = 0
var matrixCoordinates [][]int
var total int = 0

func walk(x int, y int, whoAmI string, cameFrom string) int {
	downChars := "7|F"
	upChars := "|JL"
	rightChars := "FL-"
	leftChars := "-J7"
	matrixCoordinates = append(matrixCoordinates, []int{x, y})

	if whoAmI == "S" && (cameFrom != "") {
		fmt.Printf("count: %v\n", count)
		resto := count % 2
		divisao := count / 2
		total := resto + divisao
		fmt.Printf("total: %v\n", total)
		return total
	}
	if whoAmI == "S" {
		if y < len(matrix[0])-1 {
			rightChar := matrix[x][y+1]
			if strings.Contains(leftChars, rightChar) {
				return walk(x, y+1, rightChar, "left")
			}
		}
		if x < len(matrix[0])-1 {
			downChar := matrix[x+1][y]
			if strings.Contains(upChars, downChar) {
				return walk(x+1, y, downChar, "up")
			}
		}
		if y > 0 {

			leftChar := matrix[x][y-1]
			if strings.Contains(rightChars, leftChar) {
				return walk(x, y-1, leftChar, "right")
			}
		}
		if x > 0 {

			upChar := matrix[x-1][y]
			if strings.Contains(downChars, upChar) {
				return walk(x-1, y, upChar, "down")
			}
		}
	}
	count++
	newChar := matrix[x][y]
	fmt.Printf("whoAmI: %v\n", whoAmI)

	fmt.Printf("newChar: %v\n", newChar)
	fmt.Printf("x: %v\n", x)
	fmt.Printf("y: %v\n", y)
	fmt.Printf("cameFrom: %v\n", cameFrom)
	fmt.Printf("count: %v\n", count)
	total = count
	fmt.Printf("total: %v\n", total)

	if strings.Contains(upChars, newChar) && (cameFrom != "up") {
		return walk(x-1, y, newChar, "down")
	}
	if strings.Contains(downChars, newChar) && (cameFrom != "down") {
		return walk(x+1, y, newChar, "up")
	}
	if strings.Contains(leftChars, newChar) && (cameFrom != "left") {
		return walk(x, y-1, newChar, "right")
	}
	if strings.Contains(rightChars, newChar) && (cameFrom != "right") {
		return walk(x, y+1, newChar, "left")
	}
	return 0

}

func pickTheorem(perimeter int, area int) int {
	return area - (perimeter / 2) + 1

}
func calculateAreaPolygon(matrixCoordinates [][]int) int {
	var area int = 0
	var x1, y1, x2, y2 int
	for i := 0; i < len(matrixCoordinates)-1; i++ {
		x1 = matrixCoordinates[i][0]
		y1 = matrixCoordinates[i][1]
		x2 = matrixCoordinates[i+1][0]
		y2 = matrixCoordinates[i+1][1]
		area += (x1 * y2) - (x2 * y1)
	}
	return area / 2
}

func main() {
	conteudo := openfile.Scan_file("example")
	const animal = "S"
	const BORDERS = "SPOIV"
	// Create matrix from conteudo
	for _, line := range conteudo {
		var row []string
		for _, char := range line {
			row = append(row, string(char))
		}
		matrix = append(matrix, row)
	}
	x, y := findSLocation(animal, matrix)
	println(x, y)
	walk(x, y, "S", "")
	fmt.Printf("matrixCoordinates: %v\n", matrixCoordinates)
	area := calculateAreaPolygon(matrixCoordinates)
	fmt.Printf("area: %v\n", area)
	fmt.Printf("total: %v\n", total)
	pickvalue := pickTheorem(total, area)
	fmt.Printf("pickvalue: %v\n", pickvalue)

}
