package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func part1(scanner *bufio.Scanner) int {

	horizontal := 0
	depth := 0

	for scanner.Scan() {
		split := strings.Split(scanner.Text(), " ")
		command := split[0]
		value, err := strconv.Atoi(split[1])
		if err != nil {
			fmt.Println(err)
			continue
		}
		switch command {
		case "forward":
			horizontal += value
		case "up":
			depth -= value
		case "down":
			depth += value
		}
	}
	return horizontal * depth
}

func part2(scanner *bufio.Scanner) int {
	horizontal := 0
	depth := 0
	aim := 0

	for scanner.Scan() {
		split := strings.Split(scanner.Text(), " ")
		command := split[0]
		value, err := strconv.Atoi(split[1])
		if err != nil {
			fmt.Println(err)
			continue
		}
		switch command {
		case "up":
			aim -= value
		case "down":
			aim += value
		case "forward":
			horizontal += value
			depth += aim * value
		}
	}
	return horizontal * depth
}

func main() {
	var part int
	flag.IntVar(&part, "part", 1, "Specify part 1 or part 2 to run")
	flag.Parse()

	file, err := os.Open("data.txt")
	if err != nil {
		fmt.Println(err)
	}
	scanner := bufio.NewScanner(file)
	fmt.Printf("Running part %d\n", part)
	var product int
	if part == 1 {
		product = part1(scanner)
	} else {
		product = part2(scanner)
	}
	fmt.Printf("Horizontal * Depth = %d\n", product)

}
