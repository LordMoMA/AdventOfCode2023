package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"unicode"
)

var TextArray = []string{}

var SumArray = []int{}
var Result int

func extractFile() []string {
	file, err := os.Open("day1.txt")
	if err != nil {
		log.Printf("file not found")
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		TextArray = append(TextArray, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		fmt.Printf("scan error")
	}
	return TextArray
}

func extractCalibration(line string) int {
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
		fmt.Errorf("first digit not found")
	}

	newDigit, err := strconv.Atoi(string([]rune{firstDigit, lastDigit}))
	if err != nil {
		fmt.Errorf("could not convert to int")
	}

	return newDigit
}

func formSum(textArray []string) []int {
	for _, v := range textArray {
		SumArray = append(SumArray, extractCalibration(v))
	}
	return SumArray
}

func calcSum(sumArray []int) int {
	for _, v := range sumArray {
		Result += v
	}
	return Result
}
