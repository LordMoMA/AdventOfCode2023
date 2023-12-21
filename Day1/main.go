package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	// part one
	textArray := extractFile()
	SumArray := formSum(textArray)
	Result := calcSum(SumArray)
	fmt.Println(Result)

	// part two
	file, err := os.Open("day1.txt")
	if err != nil {
		log.Printf("file not found")
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	sum2 := 0
	for scanner.Scan() {
		line := scanner.Text()
		sum2 += firstNumber(line)*10 + lastNumber(line)
	}

	fmt.Println(sum2)
}
