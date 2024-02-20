package main

import (
	"fmt"
	"os"
	"strings"
)

var scoreMap map[string]int

func init() {
	scoreMap = make(map[string]int)

	scoreMap["X"] = 1
	scoreMap["Y"] = 2
	scoreMap["Z"] = 3

	scoreMap["A"] = 1
	scoreMap["B"] = 2
	scoreMap["C"] = 3
}

func main() {
	b, err := os.ReadFile("2022/day2/input.txt")

	dataStr := string(b)

	if err != nil {
		fmt.Println("error", err)
	}

	matches := strings.Split(dataStr, "\n")

	var myScore int
	for _, match := range matches {
		// A,B,C
		// X,Y,Z
		//Rock,Paper,Scissor

		// hisChoice := string(match[0])
		// myChoice := string(match[2])

		// myScore += scoreMap[myChoice]
		// myScore += matchPoints(myChoice, hisChoice)
		// fmt.Println(match, matchPoints(myChoice, hisChoice))
		myChoice := string(match[2])

		if myChoice == "X" {
			myScore += 1
		}
		if myChoice == "Y" {
			myScore += 4
		}
		if myChoice == "Z" {
			myScore += 7
		}
	}
	fmt.Println(myScore)
}

func matchPoints(myChoice string, hisChoice string) int {
	if scoreMap[myChoice] == scoreMap[hisChoice] {
		return 6
	}

	if scoreMap[myChoice] == 1 && scoreMap[hisChoice] == 2 {
		return 0
	}

	if scoreMap[myChoice] == 1 && scoreMap[hisChoice] == 3 {
		return 12
	}

	if scoreMap[myChoice] == 2 && scoreMap[hisChoice] == 1 {
		return 12
	}

	if scoreMap[myChoice] == 2 && scoreMap[hisChoice] == 3 {
		return 0
	}

	if scoreMap[myChoice] == 3 && scoreMap[hisChoice] == 1 {
		return 0
	}

	if scoreMap[myChoice] == 3 && scoreMap[hisChoice] == 2 {
		return 12
	}

	return 0
}
