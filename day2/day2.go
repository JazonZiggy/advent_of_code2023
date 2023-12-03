package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Set struct {
	blue  int
	red   int
	green int
}

type Game struct {
	id  int
	set []Set
}

func create_game(s string) Game {

	splitted := strings.Split(s, ": ")
	fmt.Printf("splitted: %v\n", splitted[0])
	strGameID := strings.Split(splitted[0], " ")
	gameId, err := strconv.Atoi(strGameID[1])
	println(gameId)
	if err != nil {
		log.Fatal(err)
	}
	setArray := create_set_array(splitted[1])
	var gameCreated = Game{
		id:  gameId,
		set: setArray,
	}

	checkRules(gameCreated)

	return gameCreated
}

func create_set_array(s string) []Set {

	var setArray []Set
	fmt.Printf("s: %v\n", s)
	splitted := strings.Split(s, "; ")
	for _, v := range splitted {
		fmt.Printf("v: %v\n", v)
		setArray = append(setArray, create_set(v))

	}
	fmt.Printf("setArray: %v\n", setArray)
	return setArray

}
func create_set(s string) Set {
	blue := 0
	red := 0
	green := 0

	splitSet := strings.Split(s, ", ")

	for _, v := range splitSet {
		splitOneMore := strings.Split(v, " ")

		fmt.Printf("splitOneMore: %v\n", splitOneMore)
		stringToInt, err := strconv.Atoi(splitOneMore[0])
		if err != nil {
			log.Fatal(err)
		}
		if strings.Contains(v, "blue") {
			blue = stringToInt
		} else if strings.Contains(v, "red") {
			red = stringToInt
		} else if strings.Contains(v, "green") {
			green = stringToInt
		}
	}
	var setVariable = Set{
		blue:  blue,
		red:   red,
		green: green,
	}

	return setVariable

}

func checkRules(g Game) int {
	blue := 14
	green := 13
	red := 12

	for i, s := range g.set {
		fmt.Printf("i: %v\n", i)
		fmt.Printf("s: %v\n", s)
		if s.blue > blue {
			return 0
		}
		if s.green > green {
			return 0
		}
		if s.red > red {
			return 0
		}

	}

	return g.id

}

func noMoreRules(g Game) int {
	blue := 1
	red := 1
	green := 1

	for _, s := range g.set {
		fmt.Printf("s: %v\n", s)
		if s.green > green {
			green = s.green
		}
		if s.blue > blue {
			blue = s.blue
		}
		if s.red > red {
			red = s.red
		}

	}

	return blue * red * green

}

func main() {
	println("teste")
	f, err := os.Open("example_test_day2")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	sumOfIds := 0

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		fmt.Println(scanner.Text())
		game := create_game(scanner.Text())
		sumOfIds = sumOfIds + noMoreRules(game)
	}

	fmt.Printf("sumOfIds: %v\n", sumOfIds)
}
