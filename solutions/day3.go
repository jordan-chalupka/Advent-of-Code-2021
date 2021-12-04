package solutions

import (
	"AdventOfCode2021/utils"
)

/*
--- Day 3: Binary Diagnostic ---
https://adventofcode.com/2021/day/3
*/

func boolArrayToInt(boolArray []bool) int {
	out := 0
	for _, b := range boolArray {
		out <<= 1
		if b {
			out += 1
		}
	}

	return out
}

func D3P1() (int, error) {
	input, err := utils.ReadFile("solutions/day_3_input.txt")
	if err != nil {
		return 0, err
	}

	bitCount := make([]int, len(input[0]))
	for _, line := range input {
		for i, bit := range line {
			if bit == '1' {
				bitCount[i] += 1
			} else {
				bitCount[i] -= 1
			}
		}
	}

	gammaBoolValues := make([]bool, len(bitCount))
	epsilonBoolValues := make([]bool, len(bitCount))
	for i, bit := range bitCount {
		if bit > 0 {
			gammaBoolValues[i] = true
			epsilonBoolValues[i] = false
		} else {
			gammaBoolValues[i] = false
			epsilonBoolValues[i] = true
		}
	}

	gamma := boolArrayToInt(gammaBoolValues)
	epsilon := boolArrayToInt(epsilonBoolValues)

	return gamma * epsilon, nil
}

func D3P2() (int, error) {
	// input, err := utils.ReadFile("solutions/day_3_input.txt")
	// if err != nil {
	// 	return 0, err
	// }
	return 0, nil
}
