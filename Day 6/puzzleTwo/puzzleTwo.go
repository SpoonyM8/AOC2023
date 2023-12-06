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
	var arr []string
	var time int
	var distance int

	if err != nil {
		fmt.Println("Awh oh!!")
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		arr = append(arr, scanner.Text())
	}

	time, _ = strconv.Atoi(strings.ReplaceAll(strings.Split(arr[0], ":")[1], " ", ""))
	distance, _ = strconv.Atoi(strings.ReplaceAll(strings.Split(arr[1], ":")[1], " ", ""))

	fmt.Println(getNumPossibleWinnings(time, distance))
}

func getNumPossibleWinnings(time int, distance int) int {
	numWins := 0
	for heldTime := 0; heldTime < time; heldTime++ {
		lengthTravelled := (time - heldTime) * heldTime
		if lengthTravelled > distance {
			numWins += 1
		}
	}
	return numWins
}
