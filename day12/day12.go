package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	openfile "utils"
)

func arranjo(s string, n []int) int {
	contador := 0
	length := len(s)
	somaN := 0
	for _, v := range n {
		somaN += v
	}
	for i := 0; i < length-somaN; i++ {
		contador++
	}

	return contador
}

func findArrangements(spring string, arrangements []int) int {
	length := len(spring)
	sumAllArrangements := 0
	separators := len(arrangements) - 1
	fmt.Printf("spring: %v\n", spring)
	fmt.Printf("arrangements: %v\n", arrangements)
	countHashes := strings.Count(spring, "#")
	fmt.Printf("countHashes: %v\n", countHashes)

	for i := 0; i < len(arrangements); i++ {
		sumAllArrangements += arrangements[i]
	}

	if countHashes == sumAllArrangements {
		return 1
	}
	fmt.Printf("sumAllArrangements: %v\n", sumAllArrangements)
	if length-1 == sumAllArrangements+separators {
		return 1
	}
	if len(arrangements) == 0 {
		if length == 0 {
			return 0
		}
		return 0
	}

	if sumAllArrangements+separators == length {
		return 1
	}
	if spring[0] == '.' {
		return findArrangements(spring[1:], arrangements)
	}
	if spring[length-1] == '.' {
		return findArrangements(spring[:length-1], arrangements)
	}
	if spring[0] == '#' {
		println("SOCORROOO")
		firstResult := findArrangements(spring[0:arrangements[0]], arrangements[:1])
		secondResult := findArrangements(spring[arrangements[0]+1:], arrangements[1:])
		return firstResult * secondResult
	}
	if spring[length-1] == '#' {
		lastArrangement := arrangements[len(arrangements)-1]
		firstResult := findArrangements(spring[0:length-lastArrangement], arrangements[:len(arrangements)-1])
		secondResult := findArrangements(spring[length-lastArrangement:], arrangements[len(arrangements)-1:])
		return firstResult * secondResult
	}
	//Estado final que resolve o arranjo
	// if strings.Count(spring, "?") == length {
	// 	resultadoArranjo := arranjo(spring, arrangements)
	// 	return resultadoArranjo
	// }
	if spring[0] == '?' {
		newString1 := strings.Replace(spring, "?", ".", 1)
		newArrangement := findArrangements(newString1, arrangements)
		newString2 := strings.Replace(spring, "?", "#", 1)
		newArrangement2 := findArrangements(newString2, arrangements)
		return newArrangement + newArrangement2
	}

	fmt.Printf("spring: %v\n", spring)
	fmt.Printf("arrangements: %v\n", arrangements)
	totalHashes := strings.Count(spring, "#")
	totalQuestionMarks := strings.Count(spring, "?")
	totalDots := strings.Count(spring, ".")
	fmt.Printf("totalHashes: %v\n", totalHashes)
	fmt.Printf("totalQuestionMarks: %v\n", totalQuestionMarks)
	fmt.Printf("totalDots: %v\n", totalDots)
	springLength := len(spring)
	fmt.Printf("springLength: %v\n", springLength)
	fmt.Printf("arrangements: %v\n", arrangements)

	log.Fatal("Not Resolved yet")
	println("Not Resolved yet")
	return 0

}

func readlines(line string) (string, []int) {
	separated := strings.split(line, " ")
	springs := separated[0]
	arrangements := strings.split(separated[1], ",")

	//transform arrangements into int
	arrangementsint := []int{}
	for _, v := range arrangements {
		vint, _ := strconv.atoi(v)
		arrangementsint = append(arrangementsint, vint)
	}

	return springs, arrangementsint

}

func main() {
	conteudo := openfile.Scan_file("example")
	for i, v := range conteudo {
		fmt.Printf("i: %v, v: %v\n", i, v)
	}
	count := 0
	for i, v := range conteudo {
		springs, arrangements := readLines(v)
		count += findArrangements(springs, arrangements)
		fmt.Printf("springs: %v\n", springs)
		fmt.Printf("count: %v\n", count)
		fmt.Printf("i: %v\n", i)
	}
	fmt.Printf("count: %v\n", count)
	//findArrangements(springs, arrangements)
}
