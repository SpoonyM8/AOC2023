package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Key struct {
	line, idx int
}

type Value struct {
	val      int
	hasRatio bool
}

var starMap = map[Key]Value{}

func main() {
	file, err := os.Open("./puzzleOneInput.txt")
	var schematics []string

	if err != nil {
		fmt.Println("Awh oh!!")
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		schematics = append(schematics, line)
	}

	for i := 0; i < 140; i++ {
		for j := 0; j < 140; j++ {
			starMap[Key{line: i, idx: j}] = Value{val: 0, hasRatio: false}
		}
	}
	addRatiosToStarMap(schematics)
	total := 0
	for _, v := range starMap {
		if v.hasRatio {
			total += v.val
		}
	}
	fmt.Println(total)
}

func addRatiosToStarMap(schematics []string) {
	isNum := false
	firstIdx, lastIdx := 0, 0

	for currLine, line := range schematics {
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
					addToStarMap(schematics, currLine, firstIdx, lastIdx)
					isNum = false
				}
			}
			if currChar == 139 && isNum {
				addToStarMap(schematics, currLine, firstIdx, lastIdx)
				isNum = false
			}
		}
	}
}

func addToStarMap(schematics []string, lineNum int, firstNumIndex int, lastNumIndex int) {
	hasSpecialCharAbove(schematics, lineNum, firstNumIndex, lastNumIndex)
	hasSpecialCharBelow(schematics, lineNum, firstNumIndex, lastNumIndex)
	hasSpecialCharEitherSide(schematics, lineNum, firstNumIndex, lastNumIndex)
}

func hasSpecialCharAbove(schematics []string, lineNum int, firstNumIndex int, lastNumIndex int) {
	if lineNum == 0 {
		return
	} else {
		for idx, char := range schematics[lineNum-1] {
			if idx >= firstNumIndex-1 && idx <= lastNumIndex+1 {
				if isCharSpecial(int(char)) {
					val, _ := strconv.Atoi(schematics[lineNum][firstNumIndex : lastNumIndex+1])
					addRatioToStarMap(lineNum-1, idx, val)
				}
			}
		}
	}
}

func hasSpecialCharBelow(schematics []string, lineNum int, firstNumIndex int, lastNumIndex int) {
	if lineNum == 139 {
		return
	} else {
		for idx, char := range schematics[lineNum+1] {
			if idx >= firstNumIndex-1 && idx <= lastNumIndex+1 {
				if isCharSpecial(int(char)) {
					val, _ := strconv.Atoi(schematics[lineNum][firstNumIndex : lastNumIndex+1])
					addRatioToStarMap(lineNum+1, idx, val)
				}
			}
		}
	}
}

func hasSpecialCharEitherSide(schematics []string, lineNum int, firstNumIndex int, lastNumIndex int) {
	for idx, char := range schematics[lineNum] {
		if idx >= firstNumIndex-1 && idx <= lastNumIndex+1 {
			if isCharSpecial(int(char)) {
				val, _ := strconv.Atoi(schematics[lineNum][firstNumIndex : lastNumIndex+1])
				addRatioToStarMap(lineNum, idx, val)
			}
		}
	}
}

func isCharSpecial(char int) bool {
	return char == 42 // 42 is ascii for *
}

func addRatioToStarMap(line int, idx int, val int) {

	if starMap[Key{line: line, idx: idx}].val != 0 {
		existingVal := starMap[Key{line: line, idx: idx}].val
		starMap[Key{line: line, idx: idx}] = Value{val: existingVal * val, hasRatio: true}
	} else {
		starMap[Key{line: line, idx: idx}] = Value{val: val, hasRatio: false}
	}
}
