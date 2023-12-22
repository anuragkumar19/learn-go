package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	b, err := os.ReadFile("2022/day1/input.txt")

	if err != nil {
		fmt.Println("Error reading file", err)
	}

	dataStr := string(b)

	elves := strings.Split(dataStr, "\n\n")
	var totalCal []int

	for _, elf := range elves {
		foods := strings.Split(elf, "\n")

		var total int

		for _, food := range foods {
			cal, err := strconv.Atoi(food)

			if err != nil {
				fmt.Println("Error", err)
			}
			total += cal
		}

		totalCal = append(totalCal, total)
	}

	sort.Slice(totalCal, func(i, j int) bool {
		return totalCal[i] > totalCal[j]
	})

	fmt.Println(totalCal[0])
	fmt.Println(totalCal[0] + totalCal[1] + totalCal[2])
}
