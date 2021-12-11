package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)

func part1(scanner *bufio.Scanner) (int, error) {
	// part 1 solution
	return 0, nil
}

func part2(scanner *bufio.Scanner) (int, error) {
	// part 2 solution
	return 0, nil

}

func main() {
	var part int
	var filename string
	flag.IntVar(&part, "part", 1, "Specify part 1 or part 2 to run")
	flag.StringVar(&filename, "file", "data.txt", "Enter the data filename to use")
	flag.Parse()

	file, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
	}
	scanner := bufio.NewScanner(file)
	fmt.Printf("Running part %d\n", part)
	var product int
	if part == 1 {
		product, err = part1(scanner)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Printf("Gamma * Epsilon = %d\n", product)
	} else {
		product, err = part2(scanner)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Printf("Oxygen * c02 = %d\n", product)
	}

}
