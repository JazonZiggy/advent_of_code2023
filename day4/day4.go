package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

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
	fmt.Printf("%v\n", conteudo)

	return conteudo

}

type winningCard struct {
	id      int
	numbers []int
}

type playerCard struct {
	id      int
	numbers []int
}

func compareWinningCard(wc winningCard, pc playerCard) int {
	count := 0
	for _, v := range wc.numbers {
		for _, v2 := range pc.numbers {
			if v == v2 {
				if count == 0 {
					count++
				} else {
					count *= 2
				}
			}
		}
	}
	return count
}
func compareWinningCardCreateNew(wc winningCard, pc playerCard) int {
	count := 0
	for _, v := range wc.numbers {
		for _, v2 := range pc.numbers {
			if v == v2 {
				count++
			}
		}
	}

	return count
}

func createWinningCard(id int, numbersInString string) winningCard {
	numbersSplitted := strings.Split(numbersInString, " ")
	var numbers []int
	for _, v := range numbersSplitted {
		number, _ := strconv.Atoi(v)
		if number != 0 {
			numbers = append(numbers, number)
		}
	}
	wc := winningCard{
		id:      id,
		numbers: numbers,
	}
	return wc
}

func createPlayerCard(id int, numbersInString string) playerCard {
	numbersSplitted := strings.Split(numbersInString, " ")
	var numbers []int
	for _, v := range numbersSplitted {
		number, _ := strconv.Atoi(v)
		if number != 0 {
			numbers = append(numbers, number)
		}
	}
	pc := playerCard{
		id:      id,
		numbers: numbers,
	}
	return pc
}

func returnplayerCardbyId(playerCards []playerCard, id int) playerCard {
	for _, v := range playerCards {
		if v.id == id {
			return v
		}
	}
	return playerCard{}
}

func createCards(s string) {
	lines := strings.Split(s, "\n")
	var cardNumber []string
	for _, v := range lines {
		if v != "" {
			cardNumber = append(cardNumber, v)
		}
	}
	var playerCards []playerCard
	var winningCards []winningCard

	for i, v := range cardNumber {
		cards := strings.Split(v, ":")
		fmt.Printf("cards: %v\n", cards)
		winningCard, playerCard, _ := strings.Cut(cards[1], "|")
		winningCards = append(winningCards, createWinningCard(i+1, winningCard))
		playerCards = append(playerCards, createPlayerCard(i+1, playerCard))
	}
	fmt.Printf("playerCards: %v\n", playerCards)
	length := len(playerCards)
	for i := 0; i < length; i++ {
		pc := playerCards[i]
		count := compareWinningCardCreateNew(winningCards[pc.id-1], pc)
		for j := 1; j <= count; j++ {
			playerCard := returnplayerCardbyId(playerCards, pc.id+j)
			fmt.Printf("playerCard: %v created playerCard %v ", pc.id, playerCard.id)
			playerCards = append(playerCards, playerCard)
			length++
			fmt.Printf("count: %v\n", count)
		}
	}

}

func main() {
	texto := scan_file("example_test_day4")
	fmt.Printf("texto: %v\n", texto)
	createCards(texto)
}
