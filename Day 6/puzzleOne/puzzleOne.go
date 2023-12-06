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
	var arr []string
	var times []string
	var distances []string
	total := 1

	if err != nil {
		fmt.Println("Awh oh!!")
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		arr = append(arr, scanner.Text())
	}

	times = strings.Fields(strings.Split(arr[0], ":")[1])
	distances = strings.Fields(strings.Split(arr[1], ":")[1])

	for idx, val := range times {
		time, _ := strconv.Atoi(val)
		distance, _ := strconv.Atoi(distances[idx])
		total = total * getNumPossibleWinnings(time, distance)
	}

	fmt.Println(total)
}

func getNumPossibleWinnings(time int, distance int) int {
	numWins := 0
	for heldTime := 0; heldTime < time; heldTime++ {
		lengthTravelled := (time - heldTime) * heldTime
		fmt.Println(lengthTravelled, distance)
		if lengthTravelled > distance {
			numWins += 1
		}
	}
	return numWins
}
