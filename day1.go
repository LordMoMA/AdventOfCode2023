package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"unicode"
)

// PART ONE

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

/////////////////////////////////////////////////////////
/////////////////////////////////////////////////////////
/////////////////////////////////////////////////////////

// PART TWO

var SumArray2 = []int{}

func extractCalibration2(line string) int {
	nums := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
		"five":  5,
		"six":   6,
		"seven": 7,
		"eight": 8,
		"nine":  9,
	}
	var firstDigit, lastDigit int
	firstIndex, lastIndex := 0, len(line)

	for word, num := range nums {
		index := strings.Index(line, word)
		if index != -1 && index > firstIndex {
			firstIndex = index
			firstDigit = num
		}
		if index != -1 && index < lastIndex {
			lastIndex = index
			lastDigit = num
		}
	}

	for i := 1; i <= 9; i++ {
		index := strings.Index(line, strconv.Itoa(i))
		if index != -1 && index < firstIndex {
			firstIndex = index
			firstDigit = i
		}
		if index != -1 && index > lastIndex {
			lastIndex = index
			lastDigit = i
		}
	}
	return firstDigit*10 + lastDigit

}

func formSum2(textArray []string) []int {
	for _, v := range textArray {
		SumArray2 = append(SumArray2, extractCalibration2(v))
	}
	return SumArray2
}
