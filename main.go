package main

import "fmt"

func main() {
	// day1 part one
	textArray := extractFile()
	SumArray := formSum(textArray)
	Result := calcSum(SumArray)
	fmt.Println(Result)

	// day1 part two
	SumArray2 := formSum2(textArray)
	// fmt.Println(SumArray2)
	Result2 := calcSum(SumArray2)
	fmt.Println(Result2)
}
