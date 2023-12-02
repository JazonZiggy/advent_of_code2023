package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func replaceString(t string) string {

	numbers := map[string]string{
		"1": "one",
		"2": "two",
		"3": "three",
		"4": "four",
		"5": "five",
		"6": "six",
		"7": "seven",
		"8": "eight",
		"9": "nine",
	}

	for k, v := range numbers {
		valueSize := len(v)
		if len(t) < valueSize {
			continue
		}
		if strings.Contains(t[:valueSize], v) {
			fmt.Printf("Contem o numero %s\n", v)
			return k
		}
	}
	return "-1"
}

func GetNumbers(t string) int {

	var firstNumber string = "-1"
	var secondNumber string

	for pos, char := range t {
		//fmt.Printf("Caractere %c at position %d\n", char, pos)
		if _, err := strconv.Atoi(string(char)); err == nil {
			fmt.Printf("%q looks like a number in position %d. \n", char, pos)
			if firstNumber == "-1" {
				firstNumber = string(char)
			}
			secondNumber = string(char)
			continue
		}
		stringNumberinPos := replaceString(t[pos:])
		if stringNumberinPos != "-1" {
			if firstNumber == "-1" {
				firstNumber = stringNumberinPos
			}
			secondNumber = stringNumberinPos
		}
	}
	println(firstNumber)
	println(secondNumber)
	finalNumber := firstNumber + secondNumber
	finalNumberToInt, err := strconv.ParseInt(finalNumber, 10, 16)
	if (err) != nil {
		log.Fatal(err)
	}
	println(finalNumberToInt)
	return int(finalNumberToInt)
}

func main() {
	fmt.Println("Hello World!")

	f, err := os.Open("example")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	sum := 0

	for scanner.Scan() {
		//fmt.Println(scanner.Text())
		number := GetNumbers(scanner.Text())
		sum = sum + number

	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	println(sum)
}
