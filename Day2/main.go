package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func extractFile(filename string) int {
	file, err := os.Open(filename)
	if err != nil {
		log.Printf("file not found")
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var result = 0

	for scanner.Scan() {
		line := scanner.Text()
		count := calculateGameID(line)
		result += count
	}
	return result
}

func calculateGameID(line string) int {
	bag := map[string]int{
		"red":   12,
		"blue":  14,
		"green": 13,
	}

	var isPossible = true
	var sum = 0
	var gameID int

	// for example:
	// "Game 1: 1 red, 5 blue, 10 green; 5 green, 6 blue, 12 red; 4 red, 10 blue, 4 green"
	parts := strings.Split(line, ":")
	words := strings.Split(parts[0], " ")

	gameID, err := strconv.Atoi(words[1])
	if err != nil {
		log.Printf("error converting gameID to int")
	}

	rounds := strings.Split(parts[1], ";") // Fix delimiter here

	// 1 red, 5 blue, 10 green
	for _, round := range rounds {
		isPossible = true
		colorTotals := map[string]int{"red": 0, "green": 0, "blue": 0}
		colors := strings.Split(round, ",") // 1 red

		for _, color := range colors {
			count := strings.Split(strings.TrimSpace(color), " ")
			num, _ := strconv.Atoi(count[0])
			colorTotals[count[1]] += num // count[1] is the color in a set within a round
			if colorTotals[count[1]] > bag[count[1]] {
				isPossible = false
				break
			}
		}

		if !isPossible {
			break
		}
	}

	if isPossible {
		sum += gameID
	}
	return sum
}

func main() {
	// part one
	result := extractFile("day2.txt")
	fmt.Println(result)
}
