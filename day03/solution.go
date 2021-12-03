package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)

func part1(scanner *bufio.Scanner) (int, error) {
	const bitlen = 12
	var bits [bitlen][2]int

	for scanner.Scan() {
		valTxt := scanner.Text()
		for i := 0; i < bitlen; i++ {
			if valTxt[i] == '1' {
				bits[i][1] += 1
			} else {
				bits[i][0] += 1
			}
		}
	}
	gamma := 0
	epsilon := 0

	for i := 0; i < bitlen; i++ {
		if bits[i][1] > bits[i][0] {
			gamma += (1 << (bitlen - 1 - i))
		} else {
			epsilon += (1 << (bitlen - 1 - i))
		}
	}
	fmt.Printf("Bits = %v\n", bits)
	fmt.Printf("Gamma = 0b%b (%d)\n", gamma, gamma)
	fmt.Printf("Epsilon = 0b%b (%d)\n", epsilon, epsilon)
	return gamma * epsilon, nil
}

func part2(scanner *bufio.Scanner) int {
	return 0
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
	} else {
		product = part2(scanner)
	}
	fmt.Printf("Gamma * Epsilon = %d\n", product)

}
