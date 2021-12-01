package solutions

import (
	"AdventOfCode2021/utils"
	"strconv"
)

/*
--- Day 1: Sonar Sweep ---
https://adventofcode.com/2021/day/1
*/

func D1P1() (int, error) {
	input, err := utils.ReadFile("solutions/day_1_input.txt")
	if err != nil {
		return 0, err
	}

	prev, count := -1, 0
	for _, line := range input {
		value, _ := strconv.Atoi(line)
		if prev != -1 && value > prev {
			count += 1
		}
		prev = value
	}

	return count, nil
}

func D1P2() (int, error) {
	input, err := utils.ReadFile("solutions/day_1_input.txt")
	if err != nil {
		return 0, err
	}

	values := make([]int, len(input))
	for i := 0; i < len(values); i++ {
		values[i], _ = strconv.Atoi(input[i])
	}

	prev, count := -1, 0
	for i := 0; i < len(values)-2; i++ {
		sum := values[i] + values[i+1] + values[i+2]
		if prev != -1 && sum > prev {
			count += 1
		}
		prev = sum
	}

	return count, nil
}
