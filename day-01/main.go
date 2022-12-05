package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

var inputfile = "./input"

type CalorieCounter struct {
	current int

	list []int
}

func (c *CalorieCounter) Add(calories string) {
	converted, _ := strconv.Atoi(calories)
	c.current += converted
}

func (c *CalorieCounter) Next() {
	c.list = append(c.list, c.current)
	sort.Ints(c.list)
	if len(c.list) == 4 {
		c.list = c.list[1:]
	}
	c.current = 0
}

func (c *CalorieCounter) Top() int {
	total := 0
	for _, i := range c.list {
		total += i
	}
	return total
}

func (c *CalorieCounter) Max() int {
	return c.list[2]
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

	fmt.Printf("Top 3 sum: %d. Max: %d\n", c.Top(), c.Max())
}
