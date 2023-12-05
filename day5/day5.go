package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type seedsMap struct {
	name        string
	destination []int
	source      []int
	rangeSize   []int
}

func scan_file(nameOfFile string) string {
	f, err := os.Open(nameOfFile)

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()
	scanner := bufio.NewScanner(f)

	var conteudo string

	for scanner.Scan() {
		conteudo += scanner.Text() + "\n"
	}

	return conteudo

}

func getSeeds(conteudo string) []int {
	firstLine := strings.Split(conteudo, "\n")[0]
	numbersWithSpaces := strings.Split(firstLine, ":")[1]
	numbersWithTrim := strings.Split(numbersWithSpaces, " ")
	var numbers []int
	for _, v := range numbersWithTrim {
		number := strings.Trim(v, " ")
		numberInt, err := strconv.Atoi(number)
		if err != nil {
			log.Println(err)
			continue
		}
		numbers = append(numbers, numberInt)
	}
	return numbers

}

func createTheMap(title string, numbers []int) seedsMap {
	firstRow := []int{}
	secondRow := []int{}
	thirdRow := []int{}
	for i := 0; i < len(numbers)-1; i += 3 {
		firstRow = append(firstRow, numbers[i])
		secondRow = append(secondRow, numbers[i+1])
		thirdRow = append(thirdRow, numbers[i+2])
	}
	seedsMap := seedsMap{
		name:        title,
		destination: firstRow,
		source:      secondRow,
		rangeSize:   thirdRow,
	}
	return seedsMap
}

func createMaps(conteudo string) []seedsMap {
	valueStr := strings.Fields(strings.TrimSpace(conteudo))
	title := ""
	var numbers []int
	var locationMaps []seedsMap
	for i, v := range valueStr {
		if title != "" {
			tempNumber, err := strconv.Atoi(v)
			if err != nil && v != "map:" {
				log.Println(err)
				seedsMap := createTheMap(title, numbers)
				locationMaps = append(locationMaps, seedsMap)
				numbers = numbers[:0]
				title = ""
			} else {
				numbers = append(numbers, tempNumber)
			}
		}
		if v == "map:" {
			title = valueStr[i-1]
		}
	}
	if title != "" {
		seedsMap := createTheMap(title, numbers)
		locationMaps = append(locationMaps, seedsMap)
	}
	return locationMaps

}

var minimumValue int = 9999999999999

func checkMinorLocation(minimum int) {
	if minimumValue > minimum {
		minimumValue = minimum
		fmt.Printf("minimumValue: %v\n", minimumValue)
	}
}

func compareSeedWithLocation(s int, sml []seedsMap, count int) int {
	sm := sml[count]
	size := len(sm.rangeSize)
	for i := 0; i < size; i++ {
		if s >= sm.source[i] && s < sm.source[i]+sm.rangeSize[i] {
			newCounter := count + 1
			newSeed := s + sm.destination[i] - sm.source[i]
			if sm.name == "humidity-to-location" {
				checkMinorLocation(newSeed)
				return newSeed
			}
			return compareSeedWithLocation(newSeed, sml, newCounter)
		}
	}
	if sm.name == "humidity-to-location" {
		checkMinorLocation(s)
		return s
	}
	return compareSeedWithLocation(s, sml, count+1)

}

func findSeedLocation(s []int, sm []seedsMap) int {
	// for _, v := range s {
	// 	compareSeedWithLocation(v, sm, 0)
	// }
	for i := 0; i < len(s); i += 2 {
		for j := s[i]; j < s[i]+s[i+1]; j++ {
			compareSeedWithLocation(j, sm, 0)
		}
	}

	return 0
}

func main() {
	conteudo := scan_file("example_test_day5")
	seeds := getSeeds(conteudo)
	seedMaps := createMaps(conteudo)
	findSeedLocation(seeds, seedMaps)
	fmt.Printf("seeds: %v\n", seeds)
}
