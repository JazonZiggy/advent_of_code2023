package main

import (
	"fmt"
	"math"
	"utils"
)

type Universe struct {
	x int
	y int
}

var expandedXs = []int{}
var expandedYs = []int{}

type HashSet map[Universe]bool

func (set HashSet) Add(value Universe) {
	set[value] = true
}

func (set HashSet) Remove(value Universe) {
	delete(set, value)
}

func (set HashSet) Contains(value Universe) bool {
	_, exists := set[value]
	return exists
}

func (set HashSet) FoundByX(value int) bool {
	for k := range set {
		if k.x == value {
			return true
		}
	}
	return false
}

func (set HashSet) FoundByY(value int) bool {
	for k := range set {
		if k.y == value {
			return true
		}
	}
	return false
}

func universeConstructor(x, y int) Universe {
	universe := Universe{
		x: x,
		y: y,
	}
	return universe
}

func calcDistance(from Universe, to Universe) int {
	maiorX := math.Max(float64(from.x), float64(to.x))
	menorX := math.Min(float64(from.x), float64(to.x))
	maiorY := math.Max(float64(from.y), float64(to.y))
	menorY := math.Min(float64(from.y), float64(to.y))

	expandedUniveseX := 0
	expandedUniveseY := 0

	for _, v := range expandedXs {
		if v > int(menorX) && v < int(maiorX) {
			expandedUniveseX += 1000000 - 1
		}
	}
	for _, v := range expandedYs {
		if v > int(menorY) && v < int(maiorY) {
			expandedUniveseY += 1000000 - 1
		}
	}
	stepsinX := math.Abs(float64(from.x) - float64(to.x))
	stepsinY := math.Abs(float64(from.y) - float64(to.y))
	dist := int(stepsinX+stepsinY) + expandedUniveseX + expandedUniveseY

	return int(dist)
}

func CalculateAllPairsDistances(set HashSet) int {
	distances := 0
	alreadysaw := make(map[Universe]bool)
	for k := range set {
		for k2 := range set {
			if k != k2 {
				if !alreadysaw[k2] {
					distances += calcDistance(k, k2)
				}
			}
		}
		alreadysaw[k] = true
	}

	return distances
}

func main() {
	conteudo := openfile.Scan_file("example")
	fmt.Printf("conteudo: %v\n", conteudo)
	mySet := make(HashSet)

	//fmt.Printf("conteudo: %v\n", conteudo)
	matrix := openfile.Convert_matrix(conteudo)
	lengthXUniverse := len(matrix[0])
	lengthYUniverse := len(matrix)

	for i, v := range matrix {
		for i2, v2 := range v {
			if v2 == "#" {
				universe := universeConstructor(i2, i)
				mySet.Add(universe)

			}
		}
	}
	for i := 0; i < lengthXUniverse-1; i++ {
		if !mySet.FoundByX(i) {
			fmt.Printf("FoundByX: %v\n", i)
			expandedXs = append(expandedXs, i)
		}
	}
	for i := 0; i < lengthYUniverse-1; i++ {
		if !mySet.FoundByY(i) {
			fmt.Printf("FoundByY: %v\n", i)
			expandedYs = append(expandedYs, i)
		}
	}
	fmt.Printf("expandedXs: %v\n", expandedXs)
	fmt.Printf("expandedYs: %v\n", expandedYs)

	fmt.Printf("lengthXUniverse: %v\n", lengthXUniverse)
	fmt.Printf("lengthYUniverse: %v\n", lengthYUniverse)
	fmt.Printf("mySet: %v\n", mySet)
	distance := CalculateAllPairsDistances(mySet)
	fmt.Printf("distance: %v\n", distance)

}
