package main

import (
	"fmt"
	"strings"
	"utils"
)

type nodeMap struct {
	name  string
	left  string
	right string
}

func nodeMapCreator(s string) nodeMap {
	fmt.Printf("s: %v\n", s)
	data := []string{}
	for _, v := range strings.Fields(s) {
		trimmed := strings.Trim(v, "=(,)")
		if trimmed != "" {
			data = append(data, trimmed)
		}
	}
	fmt.Printf("data: %v\n", data)
	nodeMap := nodeMap{
		name:  data[0],
		left:  data[1],
		right: data[2],
	}
	return nodeMap

}

func gcd(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t

	}
	return a
}

func lcm(a, b int) int {
	return a * b / gcd(a, b)
}

func calculateMMC(numbers []int) int {
	if len(numbers) < 2 {
		panic("numbers must have at least 2 elements")
	}
	result := numbers[0]
	for i := 1; i < len(numbers); i++ {
		result = lcm(result, numbers[i])
	}
	return result
}

func main() {
	conteudo := openfile.Scan_file("day8")
	instructions := conteudo[0]
	fmt.Printf("instructions: %v\n", instructions)
	nodes := map[string]nodeMap{}
	var head []nodeMap
	for _, v := range conteudo {
		if strings.Contains(v, "=") {
			nodeMap := nodeMapCreator(v)
			// Check if nodes empty
			if strings.HasSuffix(nodeMap.name, "A") {
				head = append(head, nodeMap)
			}

			nodes[nodeMap.name] = nodeMap
		}
	}
	fmt.Printf("head: %v\n", head)
	instructionsArray := strings.Split(instructions, "")
	tailFounded := false
	resetcount := 0
	count := 0
	actualNode := head
	countNodes := 0
	listOfCounts := []int{}
	for !tailFounded {
		instruction := instructionsArray[resetcount]
		if instruction == "L" {
			newNode := actualNode[countNodes].left
			actualNode[countNodes] = nodes[newNode]
		}
		if instruction == "R" {
			newNode := actualNode[countNodes].right
			actualNode[countNodes] = nodes[newNode]
		}
		count++
		resetcount++
		if resetcount >= len(instructionsArray) {
			resetcount = 0
		}
		actualNodeName := actualNode[countNodes].name
		if strings.HasSuffix(actualNodeName, "Z") {
			countNodes++
			fmt.Printf("count: %v\n", count)
			listOfCounts = append(listOfCounts, count)
			count = 0

		}
		if countNodes >= len(actualNode) {
			tailFounded = true
		}

	}
	fmt.Printf("listOfCounts: %v\n", listOfCounts)
	fmt.Printf("calculateMMC: %v\n", calculateMMC(listOfCounts))

}
