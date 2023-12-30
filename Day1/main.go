package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"unicode"
)

// part one

func extractFile(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("failed to open file: %v", err)
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("failed to scan file: %v", err)
	}
	return lines, nil
}

func extractCalibration(line string) (int, error) {
	var firstDigit, lastDigit rune
	for _, v := range line {
		if unicode.IsDigit(v) {
			firstDigit = v
			break
		}
	}
	for i := len(line) - 1; i >= 0; i-- {
		if unicode.IsDigit(rune(line[i])) {
			lastDigit = rune(line[i])
			break
		}
	}
	if firstDigit == 0 {
		return 0, fmt.Errorf("first digit not found")
	}
	newDigit, err := strconv.Atoi(string([]rune{firstDigit, lastDigit}))
	if err != nil {
		return 0, fmt.Errorf("failed to convert to int: %v", err)
	}
	return newDigit, nil
}

func formSum(lines []string) ([]int, error) {
	var sumArray []int
	for _, line := range lines {
		calibration, err := extractCalibration(line)
		if err != nil {
			return nil, err
		}
		sumArray = append(sumArray, calibration)
	}
	return sumArray, nil
}

func calcSum(sumArray []int) int {
	var result int
	for _, v := range sumArray {
		result += v
	}
	return result
}

// part two

var nums = []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}

func firstNumber(s string) (int, error) {
	acc := ""

	for _, r := range s {
		if unicode.IsDigit(r) {
			return int(r - '0'), nil
		}

		acc += string(r)

		for i, d := range nums {
			if strings.HasSuffix(acc, d) {
				return i + 1, nil
			}
		}
	}
	return 0, errors.New("no number found in string")
}

func lastNumber(s string) (int, error) {
	acc := ""

	for i := len(s) - 1; i >= 0; i-- {
		if unicode.IsDigit(rune(s[i])) {
			return int(s[i] - '0'), nil
		}

		acc = string(s[i]) + acc

		for i, d := range nums {
			if strings.HasPrefix(acc, d) {
				return i + 1, nil
			}
		}
	}
	return 0, errors.New("no number found in string")
}

func processFile(filename string) (int, error) {
	file, err := os.Open(filename)
	if err != nil {
		return 0, fmt.Errorf("failed to open file: %v", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	sum := 0
	for scanner.Scan() {
		line := scanner.Text()
		firstNum, err := firstNumber(line)
		if err != nil {
			return 0, err
		}
		lastNum, err := lastNumber(line)
		if err != nil {
			return 0, err
		}
		sum += firstNum*10 + lastNum
	}
	return sum, nil
}

func main() {

	// part one
	lines, err := extractFile("day1.txt")
	if err != nil {
		log.Fatalf("failed to extract file: %v", err)
	}
	sumArray, err := formSum(lines)
	if err != nil {
		log.Fatalf("failed to form sum: %v", err)
	}
	result := calcSum(sumArray)
	fmt.Println(result)

	// part two

	sum2, err := processFile("day1.txt")
	if err != nil {
		log.Fatalf("failed to process file: %v", err)
	}
	fmt.Println(sum2)
}

/*
An alternative implementation of firstNumber and lastNumber
func firstNumber(s string) int {
	acc := ""

	for i := 0; i < len(s); i++ {
		if s[i] >= '0' && s[i] <= '9' {
			return int(s[i] - '0')
		}

		acc += string(s[i])

			for i, d := range nums {
				if strings.HasSuffix(acc, d) {
					return i + 1
				}
			}
		}
		log.Fatal("not found")
		return 0
	}

func lastNumber(s string) int {
	acc := ""

	for i := len(s) - 1; i >= 0; i-- {
		if s[i] >= '0' && s[i] <= '9' {
			return int(s[i] - '0')
		}

		acc = string(s[i]) + acc

		for i, d := range nums {
			if strings.HasPrefix(acc, d) {
				return i + 1
			}
		}
	}
	log.Fatal("not found")
	return 0
}
*/
