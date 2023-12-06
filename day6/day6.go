package main

import (
	"fmt"
	"strconv"
	"strings"
	"utils"
)

type corrida struct {
	time     int
	distance int
}

var speed int = 1
var distances []int
var countMultiply int = 1

func findVelocity(time int, distance int) {
	length := time
	count := 0
	for i := 0; i < length; i++ {
		timeHolded := i * speed
		distanceRunned := timeHolded * (time - i)
		if distanceRunned > distance {
			distances = append(distances, distanceRunned)
			count++
			sumOfRuns := length + 1 - i*2
			fmt.Printf("sumOfRuns: %v\n", sumOfRuns)
			fmt.Printf("i: %v\n", i)
			break
		}
	}
	countMultiply *= count

	fmt.Printf("countMultiply: %v\n", countMultiply)
	fmt.Printf("len(distances): %v\n", len(distances))
}

func main() {
	teste := openfile.Scan_file("example")

	// timers := strings.Fields(teste[0])
	// distance := strings.Fields(teste[1])

	timers := strings.Fields(teste[0])
	distance := strings.Fields(teste[1])

	timersConcat := strings.Join(timers[1:], "")
	distanceConcat := strings.Join(distance[1:], "")

	for i := 1; i < 2; i++ {
		timeInt, _ := strconv.Atoi(timersConcat)
		distanceInt, _ := strconv.Atoi(distanceConcat)
		findVelocity(timeInt, distanceInt)
	}
	//fmt.Printf("distances: %v\n", distances)

}
