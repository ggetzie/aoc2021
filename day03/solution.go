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

func filter(values []string, criteria func(string) bool) []string {
	var res []string

	for _, val := range values {
		if criteria(val) {
			res = append(res, val)
		}
	}
	return res
}

func part2(scanner *bufio.Scanner, bitlen int) int {
	var values []string
	var bits [][2]int

	for i := 0; i < bitlen; i++ {
		bits = append(bits, [2]int{0, 0})
	}

	for scanner.Scan() {
		valTxt := scanner.Text()
		values = append(values, valTxt)
		for i := 0; i < bitlen; i++ {
			if valTxt[i] == '1' {
				bits[i][1] += 1
			} else {
				bits[i][0] += 1
			}
		}
	}

	position := 0
	oxygen_candidates := values

	fmt.Printf("Bits: %v\n", bits)

	for len(oxygen_candidates) > 1 {
		oxygen_candidates = filter(oxygen_candidates, func(v string) bool {
			if bits[position][1] >= bits[position][0] {
				return v[position] == '1'
			} else {
				return v[position] == '0'
			}
		})
		if len(oxygen_candidates) < 5 {
			fmt.Printf("Position: %d: oxygen %v\n", position, oxygen_candidates)
		}
		position++
	}
	return 0
}

func main() {
	var part int
	var filename string
	var bitlen int
	flag.IntVar(&part, "part", 1, "Specify part 1 or part 2 to run")
	flag.StringVar(&filename, "file", "data.txt", "Enter the data filename to use")
	flag.IntVar(&bitlen, "bitlen", 12, "Specify the bit length")
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
		product = part2(scanner, bitlen)
	}
	fmt.Printf("Gamma * Epsilon = %d\n", product)

}
