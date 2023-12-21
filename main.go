package main

import "fmt"

func main() {
	// day 1
	textArray := extractFile()
	SumArray := formSum(textArray)
	// fmt.Println(SumArray)
	Result := calcSum(SumArray)
	fmt.Println(Result)
}
