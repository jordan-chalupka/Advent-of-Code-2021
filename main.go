package main

import (
	"AdventOfCode2021/solutions"
	"AdventOfCode2021/utils"
	"log"
)

func main() {
	p1, err := solutions.D2P1()
	if err != nil {
		log.Fatal(err)
	}

	p2, err := solutions.D2P2()
	if err != nil {
		log.Fatal(err)
	}

	utils.PrettyPrint([]interface{}{p1, p2})
}
