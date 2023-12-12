package main

import (
	"fmt"
	"strconv"
	"strings"
	"utils"
)

func subtract(a, b int) int {
	return b - a
}

var nextValue int

func findNextValue(history []int) int {
	subArray := []int{}
	nextValue += history[len(history)-1]
	if len(history) == 1 {
		fmt.Printf("nextValue: %v\n", nextValue)
		return 0
	}
	for i := 1; i < len(history); i++ {
		newValue := history[i] - history[i-1]
		subArray = append(subArray, newValue)
	}
	fmt.Printf("subArray: %v\n", subArray)
	return findNextValue(subArray)

}

func stringToIntArray(history string) []int {
	fields := strings.Fields(history)
	intHistory := []int{}
	for _, v := range fields {
		newType, _ := strconv.Atoi(v)
		intHistory = append(intHistory, newType)
	}
	return intHistory
}

func main() {
	conteudo := openfile.Scan_file("example")
	for _, v := range conteudo {
		intHistory := stringToIntArray(v)
		for i, j := 0, len(intHistory)-1; i < j; i, j = i+1, j-1 {
			intHistory[i], intHistory[j] = intHistory[j], intHistory[i]
		}

		findNextValue(intHistory)
	}
	fmt.Printf("nextValue: %v\n", nextValue)
}
