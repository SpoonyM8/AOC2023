package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("./puzzleOneInput.txt")
	var runningTotal int = 0
	if err != nil {
		fmt.Println("Awh oh!!")
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		value, isValid := checkForValidGame(line)
		if isValid {
			runningTotal += value
		}
	}
	fmt.Println(runningTotal)
}

func checkForValidGame(line string) (int, bool) {
	m := map[string]int{
		"red":   12,
		"green": 13,
		"blue":  14,
	}

	splitLine := strings.Split(line, ":")
	gameId, _ := strconv.Atoi(strings.Split(splitLine[0], " ")[1])
	rounds := splitLine[1]

	for _, round := range strings.Split(rounds, ";") {
		for _, pick := range strings.Split(round, ",") {
			splitPick := strings.Split(strings.TrimSpace(pick), " ")
			numPicked, _ := strconv.Atoi(splitPick[0])
			colourPicked := splitPick[1]

			if numPicked > m[colourPicked] {
				return gameId, false
			}
		}
	}

	return gameId, true
}
