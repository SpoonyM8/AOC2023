package main

import (
	"fmt"
	"bufio"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("./puzzleOneInput.txt")
	runningTotal := 0
	if err != nil {
		fmt.Println("Awh oh!!")
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		value, _ := strconv.Atoi(calibrationValue(line))
		runningTotal += value
	}

	fmt.Println(runningTotal)
}

func calibrationValue(str string) string {
	var firstNumFound bool = false
	var firstNum, lastNum string

	for _, char := range str {
		num := char - 48
		if num <= 9 && num >= 0 {
			if !firstNumFound {
				firstNum = string(char)
				firstNumFound = true
				lastNum = firstNum
			} else {
				lastNum = string(char)
			}
		}
	}

	return firstNum + lastNum
}
