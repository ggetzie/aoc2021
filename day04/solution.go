package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type BingoNumber struct {
	value  int
	marked bool
}

type BingoBoard [5][5]BingoNumber

func (p *BingoBoard) MarkValue(value int) {
	board := *p
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if board[i][j].value == value {
				board[i][j].marked = true
			}
		}
	}
	*p = board
}

func (p *BingoBoard) Check() (bool, int) {
	board := *p
	res := false
	total := 0
	for i := 0; i < 5; i++ {
		rowRes := true
		colRes := true
		for j := 0; j < 5; j++ {
			rowRes = rowRes && board[i][j].marked
			colRes = colRes && board[j][i].marked
			if !board[i][j].marked {
				total += board[i][j].value
			}
		}
		if rowRes || colRes {
			res = true
		}
	}
	return res, total
}

func (p *BingoBoard) Print() {
	board := *p
	for _, row := range board {
		for _, col := range row {
			fmt.Printf("%v ", col)
		}
		fmt.Print("\n")
	}
}

func (p *BingoBoard) Reset() {
	board := *p
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			board[i][j].marked = false
		}
	}
	*p = board
}

func MakeBoard(lines [5]string) BingoBoard {
	var res BingoBoard
	s := regexp.MustCompile(`\D+`)
	for i, line := range lines {
		for j, numStr := range s.Split(strings.TrimSpace(line), 5) {
			num, err := strconv.Atoi(numStr)
			if err != nil {
				log.Fatal(err)
			}
			res[i][j] = BingoNumber{
				value:  num,
				marked: false,
			}
		}
	}
	return res
}

func ReadData(scanner *bufio.Scanner) ([]int, []BingoBoard) {
	// get numbers from first line
	scanner.Scan()
	line1 := scanner.Text()
	numList := strings.Split(line1, ",")
	var numbers []int
	for _, numStr := range numList {
		num, err := strconv.Atoi(numStr)
		if err != nil {
			log.Fatal(err)
		}
		numbers = append(numbers, num)
	}
	var boards []BingoBoard
	var buffer [5]string
	i := 0
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}
		buffer[i] = line
		i++
		if i == 5 {
			board := MakeBoard(buffer)
			boards = append(boards, board)
			i = 0
		}
	}
	return numbers, boards
}

func part1(scanner *bufio.Scanner) (int, error) {
	// part 1 solution
	numbers, boards := ReadData(scanner)
	fmt.Printf("Numbers: %v\n", numbers)
	for _, number := range numbers {
		for i := 0; i < len(boards); i++ {
			boards[i].MarkValue(number)
			winner, total := boards[i].Check()
			if winner {
				fmt.Println("Winning Board:")
				boards[i].Print()
				fmt.Printf("Total: %d\n", total)
				fmt.Printf("Last number called: %d\n", number)
				return number * total, nil
			}
		}
	}
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
		fmt.Printf("Winner Score = %d\n", product)
	} else {
		product, err = part2(scanner)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Printf("Oxygen * c02 = %d\n", product)
	}

}
