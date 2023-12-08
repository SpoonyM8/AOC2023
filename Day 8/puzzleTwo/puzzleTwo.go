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
	file, err := os.Open("./puzzleTwoInput.txt")

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
	var arr = []string{}
	var individualSteps = []int{}

	for val := range m {
		if strings.HasSuffix(val, "A") {
			arr = append(arr, val)
		}
	}

	for _, dest := range arr {
		for str := dest; !strings.HasSuffix(str, "Z"); {
			if currIdx == len(sequence) {
				currIdx = 0
			}

			move := string(sequence[currIdx])

			if move == "L" {
				str = m[str].left
			} else {
				str = m[str].right
			}

			currIdx += 1
			numSteps += 1
		}
		individualSteps = append(individualSteps, numSteps)
		currIdx = 0
		numSteps = 0
	}

	fmt.Println(LCM(individualSteps[0], individualSteps[1], individualSteps[2:]...))
}

// greatest common divisor (GCD) via Euclidean algorithm, from https://siongui.github.io/2017/06/03/go-find-lcm-by-gcd/
func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

// find Least Common Multiple (LCM) via GCD, from https://siongui.github.io/2017/06/03/go-find-lcm-by-gcd/
func LCM(a, b int, integers ...int) int {
	result := a * b / GCD(a, b)

	for i := 0; i < len(integers); i++ {
		result = LCM(result, integers[i])
	}

	return result
}
