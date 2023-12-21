package main

import (
	"log"
	"strings"
)

var nums = []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}

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
