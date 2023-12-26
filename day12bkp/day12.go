package main

import (
	"fmt"
	"strconv"
	"strings"
	openfile "utils"
)

func slicesMap[T, U any](ts []T, f func(T) U) []U {
	us := make([]U, len(ts))
	for i := range ts {
		us[i] = f(ts[i])
	}
	return us
}

// homegrown impl of maps.Clear until it's standardised
func mapsClear[M ~map[K]V, K comparable, V any](m M) {
	for k := range m {
		delete(m, k)
	}
}

func atoi(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}

// pattern matching NFA that runs in O(n)
func countPossible(s []byte, c []int) int {
	pos := 0
	// state is a tuple of 4 values
	cstates := map[[4]int]int{{0, 0, 0, 0}: 1}
	nstates := map[[4]int]int{}
	for len(cstates) > 0 {
		for state, num := range cstates {
			si, ci, cc, expdot := state[0], state[1], state[2], state[3]
			if si == len(s) {
				if ci == len(c) {
					pos += num
				}
				continue
			}
			switch {
			case (s[si] == '#' || s[si] == '?') && ci < len(c) && expdot == 0:
				// we are still looking for broken springs
				if s[si] == '?' && cc == 0 {
					// we are not in a run of broken springs, so ? can be working
					nstates[[4]int{si + 1, ci, cc, expdot}] += num
				}
				cc++
				if cc == c[ci] {
					// we've found the full next contiguous section of broken springs
					ci++
					cc = 0
					expdot = 1 // we only want a working spring next
				}
				nstates[[4]int{si + 1, ci, cc, expdot}] += num
			case (s[si] == '.' || s[si] == '?') && cc == 0:
				// we are not in a contiguous run of broken springs
				expdot = 0
				nstates[[4]int{si + 1, ci, cc, expdot}] += num
			}
		}
		cstates, nstates = nstates, cstates
		mapsClear(nstates)
	}
	return pos
}

func readLines(line string) (string, []int) {
	separated := strings.Split(line, " ")
	springs := separated[0]
	arrangements := strings.Split(separated[1], ",")

	//transform arrangements into int
	arrangementsInt := []int{}
	for _, v := range arrangements {
		vInt, _ := strconv.Atoi(v)
		arrangementsInt = append(arrangementsInt, vInt)
	}

	return springs, arrangementsInt

}
func main() {
	conteudo := openfile.Scan_file("example")
	for i, v := range conteudo {
		fmt.Printf("i: %v, v: %v\n", i, v)
	}
	count := 0
	for i, v := range conteudo {
		springs, arrangements := readLines(v)
		fmt.Printf("springs: %v\n", springs)
		fmt.Printf("count: %v\n", count)
		fmt.Printf("i: %v\n", i)
		p := func() int {
			pos := 0
			cstates := map[[4]int]int{{0, 0, 0, 0}: 1}
			nstates := map[[4]int]int{}
			for len(cstates) > 0 {
				for state, num := range cstates {
					si, ci, cc, expdot := state[0], state[1], state[2], state[3]
					if si == len([]byte(springs)) {
						if ci == len(arrangements) {
							pos += num
						}
						continue
					}
					switch {
					case ([]byte(springs)[si] == '#' || []byte(springs)[si] == '?') && ci < len(arrangements) && expdot == 0:
						if []byte(springs)[si] == '?' && cc == 0 {
							nstates[[4]int{si + 1, ci, cc, expdot}] += num
						}
						cc++
						if cc == arrangements[ci] {
							ci++
							cc = 0
							expdot = 1
						}
						nstates[[4]int{si + 1, ci, cc, expdot}] += num
					case ([]byte(springs)[si] == '.' || []byte(springs)[si] == '?') && cc == 0:
						expdot = 0
						nstates[[4]int{si + 1, ci, cc, expdot}] += num
					}
				}
				cstates, nstates = nstates, cstates
				mapsClear(nstates)
			}
			return pos
		}()
		count += p
	}
	fmt.Println(count)
}
