package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

var inputfile = "./input"

type CalorieCounter struct {
	Max     int
	current int
}

func (c *CalorieCounter) Add(calories string) {
	converted, _ := strconv.Atoi(calories)
	c.current += converted
}

func (c *CalorieCounter) Next() {
	if c.current >= c.Max {
		c.Max = c.current
	}
	c.current = 0
}

func main() {
	file, err := os.Open(inputfile)
	if err != nil {
		log.Fatal(err)
	}

	c := &CalorieCounter{}
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			// next inventory
			c.Next()
			continue
		}

		c.Add(line)
	}
	c.Next()

	fmt.Println(c.Max)
}
