package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
	"utils"
)

type game struct {
	hand  string
	score int
}

type games []game

func isFiveOfAKind(a string) string {
	charList := strings.Split(a, "")
	for _, char := range charList {
		if strings.Count(a, char) == 5 {
			return "FiveOfKind"
		}
	}

	return isFourOfAKind(a)
}

func isFourOfAKind(a string) string {
	charList := strings.Split(a, "")
	for _, char := range charList {
		if strings.Count(a, char) == 4 {
			return "FourOfKind"
		}
	}
	return isFullHouse(a)
}

func isThreeOfAKind(a string) string {
	charList := strings.Split(a, "")
	//validate if not four or five
	for _, char := range charList {
		if strings.Count(a, char) == 3 {
			return "ThreeOfKind"
		}
	}
	return twoPairs(a)
}

func isFullHouse(a string) string {
	charList := strings.Split(a, "")
	//check fo FullHouse
	saveChar := ""
	for _, char := range charList {
		if strings.Count(a, char) == 3 {
			saveChar = char
			for _, char2 := range charList {
				if strings.Count(a, char2) == 2 && char2 != saveChar {
					return "FullHouse"
				}
			}
		}
	}

	return isThreeOfAKind(a)
}

func twoPairs(a string) string {
	charList := strings.Split(a, "")
	pairs := 0
	charPair := ""
	fmt.Printf("a: %v\n", a)
	for _, char := range charList {
		if strings.Count(a, char) == 2 && char != charPair {
			println("entrou")
			pairs++
			charPair = char
		}
		if pairs == 2 {
			fmt.Printf("a: %v\n", a)
			return "TwoPairs"
		}

	}
	fmt.Printf("pairs: %v\n", pairs)

	return onePair(a)
}

func onePair(a string) string {
	charList := strings.Split(a, "")
	for _, char := range charList {
		if strings.Count(a, char) == 2 {
			return "OnePair"
		}
	}
	return "HighCard"
}

func higherCardLeftToRight(a string, i int) int {
	charList := strings.Split(a, "")
	if charList[i] == "T" {
		return 10
	}
	if charList[i] == "J" {
		return 1
	}
	if charList[i] == "Q" {
		return 12
	}
	if charList[i] == "K" {
		return 13
	}
	if charList[i] == "A" {
		return 14
	}
	card, _ := strconv.Atoi(charList[i])
	return card
}

func (a games) Len() int {
	return len(a)
}

//CreateKey for the types of games from 7 types of games from 1 to 7 7 beeing the better

func (a games) Less(i, j int) bool {
	aHand := a[i].hand
	bHand := a[j].hand
	aHandType := ""
	bHandType := ""
	if strings.Contains(aHand, "J") {
		count := 0
		charToSub := ""
		copyOfHand := strings.Split(aHand, "")
		for _, v := range copyOfHand {
			temp := strings.Count(aHand, v)
			fmt.Printf("temp: %v\n", temp)
			fmt.Printf("count: %v\n", count)
			if temp > count && v != "J" {
				count = temp
				charToSub = v
			}
		}
		if charToSub == "" {
			charToSub = "J"
		}
		tempString := strings.ReplaceAll(aHand, "J", charToSub)
		aHandType = isFiveOfAKind(tempString)
	} else {
		aHandType = isFiveOfAKind(aHand)
	}

	if strings.Contains(bHand, "J") {
		count := 0
		charToSub := ""
		copyOfHand := strings.Split(bHand, "")
		for _, v := range copyOfHand {
			temp := strings.Count(bHand, v)
			if temp > count && v != "J" {
				count = temp
				charToSub = v
			}
		}
		if charToSub == "" {
			charToSub = "J"
		}
		tempString := strings.ReplaceAll(bHand, "J", charToSub)
		bHandType = isFiveOfAKind(tempString)
	} else {
		bHandType = isFiveOfAKind(bHand)
	}

	mapKey := map[string]int{
		"FiveOfKind":  7,
		"FourOfKind":  6,
		"FullHouse":   5,
		"ThreeOfKind": 4,
		"TwoPairs":    3,
		"OnePair":     2,
		"HighCard":    1,
	}
	length := len(aHand)
	if mapKey[aHandType] < mapKey[bHandType] {
		return true
	}
	if aHandType == bHandType {
		for i := 0; i < length; i++ {
			aCard := higherCardLeftToRight(aHand, i)
			bCard := higherCardLeftToRight(bHand, i)
			if aCard < bCard {
				return true
			}
			if aCard > bCard {
				return false
			}
		}
	}

	return false
}

func (a games) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]

}

func gameFactory(g string) game {
	hand := strings.Fields(g)[0]
	score, _ := strconv.Atoi(strings.Fields(g)[1])

	game := game{
		hand:  hand,
		score: score,
	}
	return game
}

func compareFunction(a game, b game) int {

	return 0
}

func main() {
	conteudo := openfile.Scan_file("example")
	teste_games := []game{}
	for _, v := range conteudo {
		game := gameFactory(v)
		teste_games = append(teste_games, game)
	}
	sumBeforeSort := 0
	for _, g := range teste_games {
		sumBeforeSort += g.score
	}
	sort.Sort(games(teste_games))
	all_score := 0
	length := len(teste_games)
	for i := 0; i < length; i++ {
		fmt.Println(teste_games[i])
		all_score += teste_games[i].score * (i + 1)

	}
	fmt.Printf("all_score: %v\n", all_score)
}
