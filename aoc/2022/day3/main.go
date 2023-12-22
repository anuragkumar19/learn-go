package main

import (
	"fmt"
	"os"
	"strings"
)

type Comp struct {
	comp1 bool
	comp2 bool
}

func main() {
	b, err := os.ReadFile("2022/day3/input.txt")

	m := make(map[rune]Comp)

	if err != nil {
		fmt.Println("error")
	}

	dataStr := string(b)

	rucksacks := strings.Split(dataStr, "\n")

	var total int

	for _, rucksack := range rucksacks {
		comp1 := rucksack[:len(rucksack)/2]
		comp2 := rucksack[len(rucksack)/2:]

		var common byte

		for _, b := range comp1 {
			m[b] = Comp{
				comp1: true,
				comp2: m[b].comp2,
			}
		}

		for _, b := range comp2 {
			if m[b].comp1 {
				common = byte(b)
			}
		}

		fmt.Println(string(common), common)

		if common > 96 && common < 123 {
			total += int(common) - 96
		} else {
			total += int(common) - 38
		}
	}

	fmt.Println(total)
}
