package main

import (
	"fmt"
	"bufio"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("./puzzleTwoInput.txt")
	runningTotal := 0
	if err != nil {
		fmt.Println("Awh oh!!")
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		value, _ := strconv.Atoi(calibrationValue(convertWordsToNums(line)))
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

func convertWordsToNums(str string) string {
	m := map[string]string {
		"one": "on1e",
		"two": "tw2o",
		"three": "thre3e",
		"four": "fou4r",
		"five": "fiv5e",
		"six": "si6x",
		"seven": "seve7n",
		"eight": "eigh8t",
		"nine": "nin9e",
	}

	for key, value := range m {
		str = strings.Replace(str, key, value, -1) 
	}

	return str
}

