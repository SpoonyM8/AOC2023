package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("./puzzleTwoInput.txt")
	var runningTotal int = 0
	if err != nil {
		fmt.Println("Awh oh!!")
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		value := checkForValidGame(line)
		runningTotal += value
	}
	fmt.Println(runningTotal)
}

func checkForValidGame(line string) int {
	m := map[string]int{
		"red":   0,
		"green": 0,
		"blue":  0,
	}

	rounds := strings.Split(line, ":")[1]

	for _, round := range strings.Split(rounds, ";") {
		for _, pick := range strings.Split(round, ",") {
			splitPick := strings.Split(strings.TrimSpace(pick), " ")
			numPicked, _ := strconv.Atoi(splitPick[0])
			colourPicked := splitPick[1]

			if numPicked > m[colourPicked] {
				m[colourPicked] = numPicked
			}
		}
	}

	return m["red"] * m["green"] * m["blue"]
}
