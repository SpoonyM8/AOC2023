package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Player struct {
	hand string
	bid  int
}

func main() {
	file, err := os.Open("./puzzleOneInput.txt")

	players := []Player{}

	if err != nil {
		fmt.Println("Awh oh!!")
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		strippedLine := strings.Split(scanner.Text(), " ")
		playersHand := strippedLine[0]
		playersBid, _ := strconv.Atoi(strippedLine[1])
		players = append(players, Player{hand: playersHand, bid: playersBid})
	}

	sort.Slice(players, func(i, j int) bool {
		player1Hand := players[i].hand
		player2Hand := players[j].hand
		player1Score := getScore(player1Hand)
		player2Score := getScore(player2Hand)
		if player1Score != player2Score {
			return player1Score < player2Score
		}
		return isHand2Highest(player1Hand, player2Hand)
	})

	maxRank := len(players)
	total := 0
	for i := 1; i < maxRank+1; i++ {
		total = total + (i * players[i-1].bid)
	}
	fmt.Println(total)
}

func getScore(hand string) int {
	// 5 of a kind = 20
	// 4 of a kind = 19
	// full house = 18
	// 3 of a kind = 17
	// 2 pair = 16
	// 1 pair = 15
	var handCountsMap = map[string]int{}
	for _, card := range hand {
		if val, exists := handCountsMap[string(card)]; !exists {
			handCountsMap[string(card)] = 1
		} else {
			handCountsMap[string(card)] = val + 1
		}
	}

	// check for highest number in map
	// if its 5 or 4, then score is known
	// if its 3, need to check for a 2 for full house, otherwise its 3OAK
	// if its 2, need to check for another for another 2, otherwise its 1Pair
	maxNum := 0
	numPairs := 0
	for _, v := range handCountsMap {
		if v > maxNum {
			maxNum = v
		}
		if v == 2 {
			numPairs += 1
		}
	}

	if maxNum == 5 || maxNum == 4 {
		return maxNum + 15 //5OAK or 4OAK
	} else if maxNum < 2 {
		return 0 // highest card comparisons later
	} else if maxNum == 3 {
		if numPairs == 1 {
			return 18 // full house
		} else {
			return 17 // 3OAK
		}
	} else if maxNum == 2 {
		if numPairs == 2 {
			return 16 // 2Pair
		} else {
			return 15 // 1Pair
		}
	} else {
		return 0 // nothing, compare highest card later
	}

}

func isHand2Highest(hand1, hand2 string) bool {
	var cardMap = map[string]int{
		"J": 1,
		"2": 2,
		"3": 3,
		"4": 4,
		"5": 5,
		"6": 6,
		"7": 7,
		"8": 8,
		"9": 9,
		"T": 10,
		"Q": 12,
		"K": 13,
		"A": 14,
	}
	for i := 0; i < 5; i++ {
		hand1CardValue := cardMap[string(hand1[i])]
		hand2CardValue := cardMap[string(hand2[i])]
		if hand1CardValue == hand2CardValue {
			continue
		} else {
			return hand1CardValue < hand2CardValue
		}
	}
	return true
}
