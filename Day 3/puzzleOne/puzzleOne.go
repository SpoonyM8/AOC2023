package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("./puzzleOneInput.txt")
	var schematics []string

	//var runningTotal int = 0
	if err != nil {
		fmt.Println("Awh oh!!")
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		schematics = append(schematics, line)
	}

	fmt.Println(getTotalPartNumbers(schematics))
}

func getTotalPartNumbers(schematics []string) int {
	isNum := false
	firstIdx, lastIdx := 0, 0
	runningTotal := 0

	for currLine, line := range schematics {
		lineTotal := 0
		for currChar, char := range line {
			num := char - 48
			if num <= 9 && num >= 0 {
				if !isNum {
					firstIdx = currChar
					lastIdx = currChar
				} else {
					lastIdx = currChar
				}
				isNum = true
			} else {
				if isNum {
					if hasAdjacentSpecialChar(schematics, currLine, firstIdx, lastIdx) {
						val, _ := strconv.Atoi(schematics[currLine][firstIdx : lastIdx+1])
						lineTotal += val
						runningTotal += val
					}
					isNum = false
				}
			}
			if currChar == 139 && isNum {
				if hasAdjacentSpecialChar(schematics, currLine, firstIdx, lastIdx) {
					val, _ := strconv.Atoi(schematics[currLine][firstIdx : lastIdx+1])
					lineTotal += val
					runningTotal += val
				}
				isNum = false
			}
		}
	}

	return runningTotal
}

func hasAdjacentSpecialChar(schematics []string, lineNum int, firstNumIndex int, lastNumIndex int) bool {
	if hasSpecialCharAbove(schematics, lineNum, firstNumIndex, lastNumIndex) {
		return true
	} else if hasSpecialCharBelow(schematics, lineNum, firstNumIndex, lastNumIndex) {
		return true
	} else if hasSpecialCharEitherSide(schematics, lineNum, firstNumIndex, lastNumIndex) {
		return true
	}
	return false
}

func hasSpecialCharAbove(schematics []string, lineNum int, firstNumIndex int, lastNumIndex int) bool {
	if lineNum == 0 {
		return false
	} else {
		for idx, char := range schematics[lineNum-1] {
			if idx >= firstNumIndex-1 && idx <= lastNumIndex+1 {
				if isCharSpecial(int(char)) {
					return true
				}
			}
		}
	}
	return false
}

func hasSpecialCharBelow(schematics []string, lineNum int, firstNumIndex int, lastNumIndex int) bool {
	if lineNum == 139 {
		return false
	} else {
		for idx, char := range schematics[lineNum+1] {
			if idx >= firstNumIndex-1 && idx <= lastNumIndex+1 {
				if isCharSpecial(int(char)) {
					return true
				}
			}
		}
	}
	return false
}

func hasSpecialCharEitherSide(schematics []string, lineNum int, firstNumIndex int, lastNumIndex int) bool {
	for idx, char := range schematics[lineNum] {
		if idx >= firstNumIndex-1 && idx <= lastNumIndex+1 {
			if isCharSpecial(int(char)) {
				return true
			}
		}
	}
	return false
}

func isCharSpecial(char int) bool {
	return char > 57 || (char < 48 && char != 46) // 46 is ascii for decimal point
}
