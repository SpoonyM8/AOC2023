package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Node struct {
	left, right string
}

func main() {
	file, err := os.Open("./puzzleOneInput.txt")

	sequence := ""

	if err != nil {
		fmt.Println("Awh oh!!")
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	var m = map[string]Node{}

	firstLine := true
	
	for scanner.Scan() {
		line := strings.ReplaceAll(scanner.Text(), " ", "")
		if firstLine {
			sequence = line
			firstLine = false
		} else if line == "" {
			continue
		} else {

			strippedLine := strings.Split(line, "=")
			pos := strippedLine[0]
			links := strippedLine[1]
			links = links[1 : len(links)-1]
			left := strings.Split(links, ",")[0]
			right := strings.Split(links, ",")[1]

			m[pos] = Node{left: left, right: right}
		}
	}

	currIdx := 0
	numSteps := 0
	for dest := "AAA"; dest != "ZZZ"; {
		if currIdx == len(sequence) {
			currIdx = 0
		}

		move := string(sequence[currIdx])

		if move == "L" {
			dest = m[dest].left
		} else {
			dest = m[dest].right
		}

		currIdx += 1
		numSteps += 1
	}

	fmt.Println(numSteps)
}
