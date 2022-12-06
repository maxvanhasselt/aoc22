package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

var inputfile = "./input"

type Shape int

const (
	Rock Shape = iota + 1
	Paper
	Scissors
)

var shapeMap = map[byte]Shape{
	'A': Rock,
	'B': Paper,
	'C': Scissors,
	'X': Rock,
	'Y': Paper,
	'Z': Scissors,
}

type Result int

const (
	Loss Result = 0
	Draw        = 3
	Win         = 6
)

var matchupMap = map[Shape]map[Shape]Result{
	Rock:     {Rock: Draw, Paper: Loss, Scissors: Win},
	Paper:    {Rock: Win, Paper: Draw, Scissors: Loss},
	Scissors: {Rock: Loss, Paper: Win, Scissors: Draw},
}

func Score(game string) int {
	opponent := shapeMap[game[0]]
	me := shapeMap[game[2]]
	return int(me) + int(matchupMap[me][opponent])
}

func main() {
	file, err := os.Open(inputfile)
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)

	score := 0
	for scanner.Scan() {
		score += Score(scanner.Text())
	}

	fmt.Println(score)
}
