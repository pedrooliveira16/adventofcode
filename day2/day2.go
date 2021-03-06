package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

var twoLetter = 0
var threeLetter = 0
var strings []string

type closeString struct {
	a   string
	b   string
	dif int
}

func dealWithIt(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func part1() {
	f, err := os.Open("input.txt")
	dealWithIt(err)
	strings = make([]string, 0)
	in := bufio.NewScanner(f)

	for in.Scan() {
		letterMap := make(map[rune]int)
		boxID := in.Text()
		strings = append(strings, boxID)
		for _, c := range boxID {
			letterMap[c]++
		}

		hasTwo, hasThree := false, false
		for _, v := range letterMap {
			if hasThree && hasTwo {
				break
			}
			if v == 2 && !hasTwo {
				twoLetter++
				hasTwo = true
			} else {
				if v == 3 && !hasThree {
					threeLetter++
					hasThree = true
				}
			}
		}

	}
	checksum := twoLetter * threeLetter
	fmt.Println(checksum)
}

func compareStrings(s1 string, s2 string) int {
	dif := 0
	for i, c := range s1 {
		if c != rune(s2[i]) {
			dif++
		}
	}
	return dif
}

func compareWithAll(str string, i int, strings []string, comp closeString) closeString {
	for j, candidate := range strings {
		if i == j {
			continue
		}
		dif := compareStrings(candidate, str)
		if comp.dif == 0 || dif < comp.dif {
			comp.a = str
			comp.b = candidate
			comp.dif = dif
		}
	}
	return comp
}

func commonLetters(a string, b string) string {
	common := ""
	for i, c := range a {
		if c == rune(b[i]) {
			common = common + string(c)
		}
	}
	return common
}

func part2() {
	comp := closeString{"", "", 0}
	for i, str := range strings {
		comp = compareWithAll(str, i, strings, comp)
	}
	fmt.Println(commonLetters(comp.a, comp.b))
}

func main() {
	part1()
	part2()
}
