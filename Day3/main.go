package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func ExtrctFileAndCalcSumNearSymbol(filename string) (int, error) {
	file, err := os.Open(filename)
	if err != nil {
		return 0, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return 0, err
	}
	sum := CalcSumNearSymbol(lines)
	return sum, nil
}

/*
(-1,-1) | (-1,0) | (-1,1)
--------------------------
(0,-1)  | (0,0)  | (0,1)
--------------------------
(1,-1)  | (1,0)  | (1,1)

dx[0] and dy[0] (-1, -1) represent moving up-left.
dx[1] and dy[1] (-1, 0) represent moving up.
dx[2] and dy[2] (-1, 1) represent moving up-right.
dx[3] and dy[3] (0, -1) represent moving left.
dx[4] and dy[4] (0, 1) represent moving right.
dx[5] and dy[5] (1, -1) represent moving down-left.
dx[6] and dy[6] (1, 0) represent moving down.
dx[7] and dy[7] (1, 1) represent moving down-right.

467..114..
...*......
..35..633.
......#...
617*......
.....+.58.
..592.....
......755.
...$.*....
.664.598..

*/

func CalcSumNearSymbol(lines []string) int {
	dx := []int{-1, -1, -1, 0, 0, 1, 1, 1}
	dy := []int{-1, 0, 1, -1, 1, -1, 0, 1}
	total := 0
	seen := make(map[string]bool)

	for i := range lines {
		for j := range lines[i] {
			if lines[i][j] >= '0' && lines[i][j] <= '9' {
				// Find the start of the number
				start := j
				for start > 0 && lines[i][start-1] >= '0' && lines[i][start-1] <= '9' {
					start--
				}
				// Find the end of the number
				end := j
				for end < len(lines[i])-1 && lines[i][end+1] >= '0' && lines[i][end+1] <= '9' {
					end++
				}
				numStr := lines[i][start : end+1]
				num, _ := strconv.Atoi(numStr)
				// Check all eight directions for a symbol
				for k := range dx {
					ni, nj := i+dx[k], start+dy[k] // check the start coordinate of the number
					// if the character at the new coordinates is not a dot and not a digit (i.e., it's a symbol).
					//In ASCII, the characters '0' to '9' have consecutive values, so any character with a value less than '0' or greater than '9' is not a digit.
					if ni >= 0 && ni < len(lines) && nj >= 0 && nj < len(lines[i]) && lines[ni][nj] != '.' && (lines[ni][nj] < '0' || lines[ni][nj] > '9') {
						key := fmt.Sprintf("%d,%d", i, start)
						if _, ok := seen[key]; !ok {
							total += num
							seen[key] = true
						}
						break
					}
					ni, nj = i+dx[k], end+dy[k] // check the end of the number
					if ni >= 0 && ni < len(lines) && nj >= 0 && nj < len(lines[i]) && lines[ni][nj] != '.' && (lines[ni][nj] < '0' || lines[ni][nj] > '9') {
						key := fmt.Sprintf("%d,%d", i, end)
						if _, ok := seen[key]; !ok {
							total += num
							seen[key] = true
						}
						break
					}
				}
			}
		}
	}
	return total
}

// part two

func ExtrctFileAndCalcGearRatio(filename string) (int, error) {
	file, err := os.Open(filename)
	if err != nil {
		return 0, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return 0, err
	}
	sum := CalcGearRatio(lines)
	return sum, nil
}

func CalcGearRatio(lines []string) int {
	dx := []int{-1, 0, 1, 0, -1, -1, 1, 1}
	dy := []int{0, 1, 0, -1, -1, 1, -1, 1}
	total := 0

	for i := range lines {
		for j := range lines[i] {
			if lines[i][j] == '*' {
				nums := []int{}
				seen := make(map[string]bool)
				for k := 0; k < 8; k++ {
					ni, nj := i+dx[k], j+dy[k]
					if ni >= 0 && ni < len(lines) && nj >= 0 && nj < len(lines[ni]) && lines[ni][nj] >= '0' && lines[ni][nj] <= '9' {
						start, end := nj, nj
						// Find the start of the number
						for start-1 >= 0 && lines[ni][start-1] >= '0' && lines[ni][start-1] <= '9' {
							start--
						}
						// Find the end of the number
						for end+1 < len(lines[ni]) && lines[ni][end+1] >= '0' && lines[ni][end+1] <= '9' {
							end++
						}

						key := fmt.Sprintf("%d,%d", ni, end)
						if _, exists := seen[key]; !exists {
							num, _ := strconv.Atoi(strings.TrimSpace(lines[ni][start : end+1]))
							nums = append(nums, num)
							seen[key] = true
						}
					}
				}
				if len(nums) == 2 {
					total += nums[0] * nums[1]
				}

			}
		}
	}

	return total
}

func main() {
	// part one
	sum, err := ExtrctFileAndCalcSumNearSymbol("day3.txt")
	if err != nil {
		panic(err)
	}
	fmt.Printf("sum: %d\n", sum)

	// part two
	sum2, err := ExtrctFileAndCalcGearRatio("day3.txt")
	if err != nil {
		panic(err)
	}
	fmt.Printf("sum2: %d\n", sum2)
}
