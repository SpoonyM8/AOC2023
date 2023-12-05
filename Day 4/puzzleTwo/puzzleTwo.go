package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var m = map[int]int{}

func main() {
	file, err := os.Open("./puzzleTwoInput.txt")
	var arr []string

	//var runningTotal int = 0
	if err != nil {
		fmt.Println("Awh oh!!")
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		gameNum, _ := strconv.Atoi(strings.Fields(strings.Split(line, ":")[0])[1])
		m[gameNum] = 1
		arr = append(arr, line)
	}
	output := doThing(arr)
	/*
		for k, v := range m {
			fmt.Println(k, v)
		}
	*/
	fmt.Println(output)

}

func doThing(arr []string) int {
	total := 0
	for _, line := range arr {
		gameNum, _ := strconv.Atoi(strings.Fields(strings.Split(line, ":")[0])[1])
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
		for i := 1; i < numMatchedNums+1; i++ {
			m[gameNum+i] = m[gameNum+i] + (1 * m[gameNum])
		}
		fmt.Println(gameNum, numMatchedNums)
		total = total + (1 * m[gameNum])
	}
	return total
}
