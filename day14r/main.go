package main

import (
	"fmt"
	"utils"
)

type rock struct {
	x     int
	y     int
	shape string
}

type rockSet map[rock]struct{}

func (r rockSet) add(rock rock) {
	r[rock] = struct{}{}
}

func (r rockSet) remove(rock rock) {
	delete(r, rock)
}

func (r rockSet) contains(rock rock) bool {
	_, c := r[rock]
	return c
}

func (r rockSet) equals(rockSet rockSet) bool {
	if len(r) != len(rockSet) {
		return false
	}
	for r, _ := range r {
		if !rockSet.contains(r) {
			return false
		}
	}
	return true
}

func (r rockSet) findAllRocksByX(x int) []rock {
	var rocks []rock
	for r, _ := range rockSetMap {
		if r.x == x {
			rocks = append(rocks, r)
		}
	}
	return rocks
}

func (r rockSet) findAllRocksByY(y int) []rock {
	var rocks []rock
	for r, _ := range rockSetMap {
		if r.y == y {
			rocks = append(rocks, r)
		}
	}
	return rocks
}

var rockSetMap = make(rockSet)

func createRock(x, y int, shape string) rock {
	return rock{x, y, shape}
}

var rockSetArray []rockSet

func moveRockUp(rock rock, x, y int) {
	if x < 0 {
		rockSetMap.add(rock)
		return
	}
	upRockHash := createRock(x, y, "#")
	if rockSetMap.contains(upRockHash) {
		rockSetMap.add(rock)
		return
	}
	upRockO := createRock(x, y, "O")
	if rockSetMap.contains(upRockO) {
		moveRockUp(rock, x-1, y)

	}
	rock.x = x
	rock.y = y
	moveRockUp(rock, x-1, y)
}
func moveRockDown(rock rock, x, y int, length int) {
	if x > length {
		rockSetMap.add(rock)
		return
	}
	upRockHash := createRock(x, y, "#")
	if rockSetMap.contains(upRockHash) {
		rockSetMap.add(rock)
		return
	}
	upRockO := createRock(x, y, "O")
	if rockSetMap.contains(upRockO) {
		moveRockDown(rock, x+1, y, length)

	}
	rock.x = x
	rock.y = y
	moveRockDown(rock, x+1, y, length)
}

func moveRockRight(rock rock, x, y int, length int) {
	if y > length {
		rockSetMap.add(rock)
		return
	}
	upRockHash := createRock(x, y, "#")
	if rockSetMap.contains(upRockHash) {
		rockSetMap.add(rock)
		return
	}
	upRockO := createRock(x, y, "O")
	if rockSetMap.contains(upRockO) {
		moveRockRight(rock, x, y+1, length)

	}
	rock.x = x
	rock.y = y
	moveRockRight(rock, x, y+1, length)
}

func moveRockLeft(rock rock, x, y int) {
	if y < 0 {
		rockSetMap.add(rock)
		return
	}
	upRockHash := createRock(x, y, "#")
	if rockSetMap.contains(upRockHash) {
		rockSetMap.add(rock)
		return
	}
	upRockO := createRock(x, y, "O")
	if rockSetMap.contains(upRockO) {
		moveRockLeft(rock, x, y-1)

	}
	rock.x = x
	rock.y = y
	moveRockLeft(rock, x, y-1)
}

func tiltUp(length int) {
	for i := 0; i < length; i++ {
		rocks := rockSetMap.findAllRocksByX(i)
		for _, r := range rocks {
			if r.shape == "O" {
				rockSetMap.remove(r)
				moveRockUp(r, r.x-1, r.y)
			}
		}
	}

}

func tiltDown(length int) {
	for i := length - 1; i >= 0; i-- {
		rocks := rockSetMap.findAllRocksByX(i)
		for _, r := range rocks {
			if r.shape == "O" {
				rockSetMap.remove(r)
				moveRockDown(r, r.x+1, r.y, length-1)
			}
		}
	}
}
func tiltRight(length int) {
	for i := length - 1; i >= 0; i-- {
		rocks := rockSetMap.findAllRocksByY(i)
		for _, r := range rocks {
			if r.shape == "O" {
				rockSetMap.remove(r)
				moveRockRight(r, r.x, r.y+1, length-1)
			}
		}
	}
}

func tiltLeft(length int) {
	for i := 0; i < length; i++ {
		rocks := rockSetMap.findAllRocksByY(i)
		for _, r := range rocks {
			if r.shape == "O" {
				rockSetMap.remove(r)
				moveRockLeft(r, r.x, r.y-1)
			}
		}
	}
}

var score int = 0
var prevScore int = -1
var firstCycle int = 0
var lastCycle int = 0

func checkScore(lengthx int) bool {
	score = 0
	for r, _ := range rockSetMap {
		multiplier := lengthx - r.x
		if r.shape == "O" {
			score += multiplier
		}
	}
	fmt.Printf("score: %v\n", score)
	for i, rs := range rockSetArray {
		if rs.equals(rockSetMap) {
			fmt.Printf("cycle: %v\n", i)
			fmt.Printf("score: %v\n", score)
			firstCycle = i
			return true
		}
	}
	newRockSetMap := make(rockSet)
	for r, _ := range rockSetMap {
		newRockSetMap.add(r)
	}
	rockSetArray = append(rockSetArray, newRockSetMap)

	return false
}

func main() {
	conteudo := openfile.Scan_file("example")

	matrix := openfile.Convert_matrix(conteudo)
	const CYCLES = 1000000000
	for i, v := range matrix {
		for i2, v2 := range v {
			if v2 != "." {
				rockSetMap.add(createRock(i, i2, v2))
			}

		}
	}
	lengthx := len(conteudo)
	lengthy := len(conteudo[0])
	for i := 0; i < CYCLES; i++ {
		tiltUp(lengthx)
		tiltLeft(lengthy)
		tiltDown(lengthx)
		tiltRight(lengthy)
		if checkScore(lengthx) {
			fmt.Printf("i: %v\n", i)
			lastCycle = i
			break
		}
	}
	sizeOfCycle := lastCycle - firstCycle
	fmt.Printf("sizeOfCycle: %v\n", sizeOfCycle)
	index := (CYCLES - firstCycle) % sizeOfCycle
	rockSetMapFinal := rockSetArray[index+firstCycle-1]
	fmt.Printf("rockSetMapFinal: %v\n", rockSetMapFinal)
	lastScore := 0
	for r, _ := range rockSetMapFinal {
		multiplier := lengthx - r.x
		if r.shape == "O" {
			lastScore += multiplier
		}
	}
	fmt.Printf("lastScore: %v\n", lastScore)

}
