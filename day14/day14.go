package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
	"utils"
)

type rock struct {
	x int
	y int
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

const STARTX = 500
const STARTY = 0

var rockSetMap = make(rockSet)

var maiorY int = 0

var sands int = 0

func createSand(startX, startY int) int {
	if rockSetMap.contains(rock{x: STARTX, y: STARTY}) {
		return sands
	}

	if startY == maiorY+2 {
		rockSetMap.add(rock{x: startX, y: startY - 1})
		sands++
		return createSand(STARTX, STARTY)
	}
	//check collider with rock
	if rockSetMap.contains(rock{x: startX, y: startY}) {
		//check diagonal left
		if !rockSetMap.contains(rock{x: startX - 1, y: startY}) {
			return createSand(startX-1, startY)
		}
		//check diagonal right
		if !rockSetMap.contains(rock{x: startX + 1, y: startY}) {
			return createSand(startX+1, startY)
		}
		rockSetMap.add(rock{x: startX, y: startY - 1})
		sands++
		return createSand(STARTX, STARTY)
	}
	return createSand(startX, startY+1)
}

func createRockPath(rocks []rock) {
	for i, v := range rocks {
		if i == len(rocks)-1 {
			break
		}
		xDifference := math.Abs(float64(rocks[i].x) - float64(rocks[i+1].x))
		yDifference := math.Abs(float64(rocks[i].y) - float64(rocks[i+1].y))
		menorX := int(math.Min(float64(rocks[i].x), float64(rocks[i+1].x)))
		menorY := int(math.Min(float64(rocks[i].y), float64(rocks[i+1].y)))

		for i2 := menorX; i2 <= menorX+int(xDifference); i2++ {
			rockSetMap.add(rock{x: int(i2), y: v.y})
		}
		for i2 := menorY; i2 <= menorY+int(yDifference); i2++ {
			rockSetMap.add(rock{x: v.x, y: int(i2)})
		}
	}
}

func splitXYs(rocks []string) {
	rockArray := []rock{}
	for _, v := range rocks {
		splited := strings.Split(v, ",")
		x, _ := strconv.Atoi(splited[0])
		y, _ := strconv.Atoi(splited[1])
		if y > maiorY {
			maiorY = y
		}
		fmt.Printf("maiorY: %v\n", maiorY)
		rockArray = append(rockArray, rock{x: x, y: y})
	}
	createRockPath(rockArray)
}

func main() {
	conteudo := openfile.Scan_file("example")
	for _, v := range conteudo {
		rocks := strings.ReplaceAll(v, "->", "")
		newRocks := strings.Fields(rocks)
		splitXYs(newRocks)
	}
	createSand(STARTX, STARTY)
	fmt.Printf("sands: %v\n", sands)
}
