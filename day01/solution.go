package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
)

func part1(scanner *bufio.Scanner) int {

	increases := 0
	line := 0
	var previous int
	var current int
	var err error

	for scanner.Scan() {
		current, err = strconv.Atoi(scanner.Text())
		if err != nil {
			fmt.Println(err)
		}

		if line > 0 && current > previous {
			increases++
		}
		line++
		previous = current
	}
	return increases
}

func has_increase(buffer [4]int) bool {
	return (buffer[3] + buffer[2] + buffer[1]) > (buffer[2] + buffer[1] + buffer[0])
}

func shuffle(num int, buffer [4]int) [4]int {
	for i := 0; i < 3; i++ {
		buffer[i] = buffer[i+1]
	}
	buffer[3] = num
	return buffer
}

func part2(scanner *bufio.Scanner) int {
	var buffer [4]int
	increases := 0
	line := 1

	for scanner.Scan() {
		num, err := strconv.Atoi(scanner.Text())
		if err != nil {
			fmt.Println(err)
		}
		buffer = shuffle(num, buffer)
		if line > 3 && has_increase(buffer) {
			increases++
		}
		line++
	}
	return increases
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
	fmt.Printf("Running part %v\n", part)
	var increases int
	if part == 1 {
		increases = part1(scanner)
	} else {
		increases = part2(scanner)
	}

	fmt.Printf("%v increases found\n", increases)

}
