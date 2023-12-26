package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	// Horário entre 08:00 e 09:00
	firstTime := generateRandomTime(8, 9)

	// Horário entre 11:30 e 13:00
	secondTime := generateRandomTime(11, 13)
	for !(secondTime.Hour() >= 11 && secondTime.Minute() >= 30) {
		secondTime = generateRandomTime(11, 13)
	}

	// Horário entre 14:00 e 14:59
	thirdTime := generateRandomTime(14, 14)
	for thirdTime.Before(secondTime.Add(time.Hour)) {
		thirdTime = generateRandomTime(14, 14)
	}

	// Horário após 17:30
	fourthTime := generateRandomTime(17, 23)
	for fourthTime.Before(thirdTime.Add(time.Hour * 8)) {
		fourthTime = generateRandomTime(17, 23)
	}

	fmt.Println("Primeiro horário:", firstTime.Format("15:04"))
	fmt.Println("Segundo horário:", secondTime.Format("15:04"))
	fmt.Println("Terceiro horário:", thirdTime.Format("15:04"))
	fmt.Println("Quarto horário:", fourthTime.Format("15:04"))

	diff1 := secondTime.Sub(firstTime)
	diff2 := fourthTime.Sub(thirdTime)
	totalDiff := diff1 + diff2

	fmt.Println("Diferença entre o segundo e primeiro horário:", diff1)
	fmt.Println("Diferença entre o quarto e terceiro horário:", diff2)
	fmt.Println("Total da diferença:", totalDiff)
}

func generateRandomTime(startHour, endHour int) time.Time {
	now := time.Now()
	minute := rand.Intn(60)
	second := rand.Intn(60)
	return time.Date(now.Year(), now.Month(), now.Day(), rand.Intn(endHour-startHour)+startHour, minute, second, 0, now.Location())
}
