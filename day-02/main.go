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

var resultMap = map[byte]Result{
	'X': Loss,
	'Y': Draw,
	'Z': Win,
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

var matchFixMap = map[Shape]map[Result]Shape{
	Rock:     {Draw: Rock, Win: Paper, Loss: Scissors},
	Paper:    {Loss: Rock, Draw: Paper, Win: Scissors},
	Scissors: {Win: Rock, Loss: Paper, Draw: Scissors},
}

func Score(game string) int {
	opponent := shapeMap[game[0]]
	me := shapeMap[game[2]]

	return int(me) + int(matchupMap[me][opponent])
}

func FixMatchAndScore(game string) int {
	opponent := shapeMap[game[0]]
	result := resultMap[game[2]]

	return int(result) + int(matchFixMap[opponent][result])
}

func main() {
	file, err := os.Open(inputfile)
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)

	score := 0
	scorePart2 := 0
	for scanner.Scan() {
		game := scanner.Text()
		score += Score(game)
		scorePart2 += FixMatchAndScore(game)
	}

	fmt.Println("part 1:", score)
	fmt.Println("part 2:", scorePart2)
}
