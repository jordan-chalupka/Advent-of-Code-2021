package solutions

import (
	"AdventOfCode2021/utils"
	"strconv"
	"strings"
)

/*
--- Day 2: Dive! ---
https://adventofcode.com/2021/day/2
*/

const (
	forward string = "forward"
	up             = "up"
	down           = "down"
)

type distance struct {
	x   int
	y   int
	aim int
}

func D2P1() (int, error) {
	input, err := utils.ReadFile("solutions/day_2_input.txt")
	if err != nil {
		return 0, err
	}

	d := distance{}
	for _, line := range input {
		fields := strings.Fields(line)
		dir := fields[0]
		mag, _ := strconv.Atoi(fields[1])

		switch dir {
		case forward:
			d.x += mag
		case up:
			d.y -= mag
		case down:
			d.y += mag
		}
	}

	totalDistance := d.x * d.y
	return totalDistance, nil
}

func D2P2() (int, error) {
	input, err := utils.ReadFile("solutions/day_2_input.txt")
	if err != nil {
		return 0, err
	}

	d := distance{}
	for _, line := range input {
		fields := strings.Fields(line)
		dir := fields[0]
		mag, _ := strconv.Atoi(fields[1])

		switch dir {
		case forward:
			d.x += mag
			d.y += d.aim * mag
		case up:
			d.aim -= mag
		case down:
			d.aim += mag
		}
	}

	totalDistance := d.x * d.y
	return totalDistance, nil
}
