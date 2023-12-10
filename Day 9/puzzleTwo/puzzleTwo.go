package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

type Node struct {
	left, right string
}

func main() {
	file, err := os.Open("./puzzleTwoInput.txt")

	if err != nil {
		fmt.Println("Awh oh!!")
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	total := 0
	for scanner.Scan() {
		line := strings.Fields(scanner.Text())
		total += resolveValue(strToIntSlice(line))
	}
	fmt.Println(total)
}

func resolveValue(ints []int) int {
	currRow := ints
	var newRow = []int{}
	var allRows = [][]int{}
	allRows = append(allRows, currRow)
	for {
		isAllZeros := true

		for i := 0; i < len(currRow)-1; i++ {
			diff := currRow[i+1] - currRow[i]
			newRow = append(newRow, diff)
			if diff != 0 {
				isAllZeros = false
			}
		}

		allRows = append(allRows, newRow)

		if !isAllZeros {
			currRow = newRow
			newRow = []int{}

			continue
		}

		for idx := len(allRows) - 1; idx >= 0; idx-- {
			if allRows[idx][len(allRows[idx])-1] == 0 {
				allRows[idx] = append(allRows[idx], 0)
			} else {
				allRows[idx] = append(allRows[idx], allRows[idx][0]-allRows[idx+1][0])
			}
			slices.Reverse(allRows[idx])
		}

		return allRows[0][0]
	}
}

func strToIntSlice(strOfNums []string) []int {
	var ints = []int{}
	for _, numAsStr := range strOfNums {
		num, _ := strconv.Atoi(string(numAsStr))
		ints = append(ints, num)
	}
	return ints
}
