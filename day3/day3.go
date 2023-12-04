package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"unicode"
)

type partNumber struct {
	positionX int
	positionY int
	size      int
	value     int
}

type symbol struct {
	positionX int
	positionY int
	char      string
}

func read_line(s string, positionY int) ([]partNumber, []symbol) {
	size := 0
	value := ""
	fmt.Printf("len(s): %v\n", len(s))
	var partNumbers []partNumber
	var symbols []symbol

	for pos, char := range s {
		//fmt.Printf("Caractere %c at position %d\n", char, pos)
		if _, err := strconv.Atoi(string(char)); err == nil {
			size++
			value = value + string(char)
			continue

		}
		if unicode.IsSymbol(char) || char != '.' {
			newSymbol := symbol{
				positionX: pos,
				positionY: positionY,
				char:      string(char),
			}
			if !unicode.IsDigit(char) {
				symbols = append(symbols, newSymbol)
			}
		}
		if value != "" {
			valueToInt, err := strconv.Atoi(value)
			if err != nil {
				log.Fatal(err)
			}
			newPartNumber := partNumber{
				positionY: positionY,
				positionX: pos - size,
				value:     valueToInt,
				size:      size,
			}
			partNumbers = append(partNumbers, newPartNumber)
			size = 0
			value = ""
		}

	}
	if value != "" {
		valueToInt, err := strconv.Atoi(value)
		if err != nil {
			log.Fatal(err)
		}
		newPartNumber := partNumber{
			positionY: positionY,
			positionX: len(s) - size,
			value:     valueToInt,
			size:      size,
		}
		partNumbers = append(partNumbers, newPartNumber)
		size = 0
		value = ""
		fmt.Printf("value: %v\n", value)
	}

	return partNumbers, symbols
}

func compare_distance(number partNumber, symbol symbol) int {
	size := number.size
	for i := 0; i < size; i++ {

		maiorX := math.Max(float64(number.positionX+i), float64(symbol.positionX))
		maiorY := math.Max(float64(number.positionY), float64(symbol.positionY))
		menorX := math.Min(float64(number.positionX+i), float64(symbol.positionX))
		menorY := math.Min(float64(number.positionY), float64(symbol.positionY))
		if (maiorX-menorX <= 1) && (maiorY-menorY <= 1) {
			fmt.Printf("partNumber: %v\n", number)
			fmt.Printf("symbol: %v\n", symbol)
			return number.value

		}
	}
	return 0

}

func compare_arrays(numbers []partNumber, symbols []symbol) {
	sumOfValues := 0
	for _, s := range symbols {
		for _, pn := range numbers {

			sumOfValues = sumOfValues + compare_distance(pn, s)
		}
	}
	fmt.Printf("sumOfValues: %v\n", sumOfValues)
}

func compare_distance_gears(number partNumber, symbol symbol) int {
	size := number.size
	for i := 0; i < size; i++ {

		maiorX := math.Max(float64(number.positionX+i), float64(symbol.positionX))
		maiorY := math.Max(float64(number.positionY), float64(symbol.positionY))
		menorX := math.Min(float64(number.positionX+i), float64(symbol.positionX))
		menorY := math.Min(float64(number.positionY), float64(symbol.positionY))
		if (maiorX-menorX <= 1) && (maiorY-menorY <= 1) {
			fmt.Printf("partNumber: %v\n", number)
			fmt.Printf("symbol: %v\n", symbol)
			return number.value
		}
	}
	return 0

}

func compare_gears(numbers []partNumber, symbols []symbol) {
	multiplysum := 0
	for _, s := range symbols {
		var total_numbers_in_gear []int
		for _, pn := range numbers {
			closestGear := compare_distance_gears(pn, s)
			if closestGear != 0 && s.char == "*" {
				total_numbers_in_gear = append(total_numbers_in_gear, closestGear)
				if len(total_numbers_in_gear) == 2 {
					multiplysum = multiplysum + (total_numbers_in_gear[0] * total_numbers_in_gear[1])
				}

			}
			fmt.Printf("multiplysum: %v\n", multiplysum)

		}
	}
	fmt.Printf("multiplysum: %v\n", multiplysum)
}

func scan_file(nameOfFile string) {
	f, err := os.Open(nameOfFile)

	if err != nil {
		log.Fatal(err)
	}
	var partNumbers []partNumber
	var symbolsFinal []symbol

	defer f.Close()
	scanner := bufio.NewScanner(f)
	counter_of_lines := 0
	for scanner.Scan() {
		numbers, symbols := read_line(scanner.Text(), counter_of_lines)
		counter_of_lines++
		partNumbers = append(partNumbers, numbers...)
		symbolsFinal = append(symbolsFinal, symbols...)
	}

	compare_gears(partNumbers, symbolsFinal)

}

func main() {
	scan_file("example_test_day3")
}
