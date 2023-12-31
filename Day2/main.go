package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

// part one

// Bag represents the available colors and their counts.
var Bag = map[string]int{
	"red":   12,
	"blue":  14,
	"green": 13,
}

// OpenFile opens a file and returns a scanner.
func OpenFile(filename string) (*bufio.Scanner, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	return bufio.NewScanner(file), nil
}

// ExtractFile reads a file and calculates the total game ID.
func ExtractFileAndCalcTotalGameID(filename string) (int, error) {
	scanner, err := OpenFile(filename)
	if err != nil {
		return 0, err
	}

	var total = 0
	for scanner.Scan() {
		line := scanner.Text()
		gameID, err := CalculateGameID(line)
		if err != nil {
			log.Printf("error calculating gameID: %v", err)
			continue
		}
		total += gameID
	}
	return total, nil
}

// CalculateGameID calculates the game ID from a line of text.
func CalculateGameID(line string) (int, error) {
	parts := strings.Split(line, ":")
	gameID, err := strconv.Atoi(strings.Split(parts[0], " ")[1])
	if err != nil {
		return 0, fmt.Errorf("error converting gameID to int: %w", err)
	}

	rounds := strings.Split(parts[1], ";")
	for _, round := range rounds {
		if !IsRoundPossible(round) {
			return 0, nil
		}
	}
	return gameID, nil
}

// IsRoundPossible checks if a round is possible given the bag of colors.
func IsRoundPossible(round string) bool {
	colorTotals := map[string]int{"red": 0, "green": 0, "blue": 0}
	colors := strings.Split(round, ",")
	for _, color := range colors {
		count := strings.Split(strings.TrimSpace(color), " ")
		num, _ := strconv.Atoi(count[0])
		colorTotals[count[1]] += num
		if colorTotals[count[1]] > Bag[count[1]] {
			return false
		}
	}
	return true
}

// part two
func ExtractFileAndCalcPowerSum(filename string) (int, error) {
	scanner, err := OpenFile(filename)
	if err != nil {
		return 0, err
	}

	var total = 0
	for scanner.Scan() {
		line := scanner.Text()
		power, err := CalculatePower(line)
		if err != nil {
			log.Printf("error calculating power: %v", err)
			continue
		}
		total += power
	}
	return total, nil
}

// "Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green"
func CalculatePower(line string) (int, error) {
	parts := strings.Split(line, ": ")
	rounds := strings.Split(parts[1], "; ")
	maxColor := map[string]int{"red": 0, "green": 0, "blue": 0}
	for _, round := range rounds {
		colors := strings.Split(round, ",")
		for _, colorset := range colors {
			colorset := strings.Split(strings.TrimSpace(colorset), " ")
			count, _ := strconv.Atoi(colorset[0])
			color := colorset[1]
			if maxColor[color] < count {
				maxColor[color] = count
			}
		}
	}
	return maxColor["red"] * maxColor["green"] * maxColor["blue"], nil
}

func main() {
	// part one
	total, err := ExtractFileAndCalcTotalGameID("day2.txt")
	if err != nil {
		log.Fatalf("error extracting file: %v", err)
	}
	fmt.Println(total)

	// part two
	power, err := ExtractFileAndCalcPowerSum("day2.txt")
	if err != nil {
		log.Fatalf("error extracting file: %v", err)
	}
	fmt.Println(power)
}
