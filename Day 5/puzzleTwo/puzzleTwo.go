package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type DestAndRange struct {
	d, r int
}

func main() {
	file, err := os.Open("./puzzleTwoInput.txt")
	seeds := [20]int{}
	seedToSoil := map[int]DestAndRange{}
	soilToFertilizer := map[int]DestAndRange{}
	fertilizerToWater := map[int]DestAndRange{}
	waterToLight := map[int]DestAndRange{}
	lightToTemperature := map[int]DestAndRange{}
	temperatureToHumidity := map[int]DestAndRange{}
	humidityTolocation := map[int]DestAndRange{}

	if err != nil {
		fmt.Println("Awh oh!!")
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	firstLine := true
	currMapName := "seed-to-soil"
	for scanner.Scan() {
		line := scanner.Text()

		if firstLine {
			strSeeds := strings.Fields(strings.Split(line, ": ")[1])
			for idx, str := range strSeeds {
				intSeed, _ := strconv.Atoi(str)
				seeds[idx] = intSeed
			}
			firstLine = false
			continue
		}

		if len(line) == 0 {
			currMapName = ""
			continue
		}

		if currMapName == "" {
			currMapName = strings.Split(line, " ")[0]
		} else {
			var m map[int]DestAndRange
			switch currMapName {
			case "seed-to-soil":
				m = seedToSoil
			case "soil-to-fertilizer":
				m = soilToFertilizer
			case "fertilizer-to-water":
				m = fertilizerToWater
			case "water-to-light":
				m = waterToLight
			case "light-to-temperature":
				m = lightToTemperature
			case "temperature-to-humidity":
				m = temperatureToHumidity
			case "humidity-to-location":
				m = humidityTolocation
			default:
				m = seedToSoil
			}
			addToMap(m, line)
		}
	}
	smallestLocation := 9999999999
	for i := 0; i < 20; i += 2 {
		for j := 0; j < seeds[i+1]; j++ {
			val := getValueFromMap(seedToSoil, seeds[i]+j)
			val = getValueFromMap(soilToFertilizer, val)
			val = getValueFromMap(fertilizerToWater, val)
			val = getValueFromMap(waterToLight, val)
			val = getValueFromMap(lightToTemperature, val)
			val = getValueFromMap(temperatureToHumidity, val)
			val = getValueFromMap(humidityTolocation, val)
			if val < smallestLocation {
				smallestLocation = val
			}
		}
	}

	fmt.Println(smallestLocation)
}

func addToMap(m map[int]DestAndRange, line string) {
	slicedLine := strings.Split(line, " ")
	destRangeStart, _ := strconv.Atoi(slicedLine[0])
	sourceRangeStart, _ := strconv.Atoi(slicedLine[1])
	rangeLength, _ := strconv.Atoi(slicedLine[2])
	m[sourceRangeStart] = DestAndRange{d: destRangeStart, r: rangeLength}
}

func getValueFromMap(m map[int]DestAndRange, key int) int {
	for src, val := range m {
		if key >= src && key <= src+val.r {
			return val.d + (key - src)
		}
	}
	return key
}
