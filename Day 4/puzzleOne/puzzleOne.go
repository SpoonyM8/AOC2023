package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("./puzzleOneInput.txt")
	var arr []string

	//var runningTotal int = 0
	if err != nil {
		fmt.Println("Awh oh!!")
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		arr = append(arr, line)
	}

	fmt.Println(doThing(arr))
}

func doThing(arr []string) int {
	total := 0
	for _, line := range arr {
		game := strings.Split(strings.Split(line, ":")[1], "|")
		winningNumbers := strings.Fields(strings.TrimSpace(game[0]))
		myNumbers := strings.Fields(strings.TrimSpace(game[1]))
		numMatchedNums := 0

		for _, winningNumber := range winningNumbers {
			for _, num := range myNumbers {
				if winningNumber == num {
					numMatchedNums += 1
					continue
				}
			}
		}
		if numMatchedNums != 0 {
			total += int(math.Pow(2, float64(numMatchedNums-1)))
		}
	}

	return total
}
