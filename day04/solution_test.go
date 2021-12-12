package main

import (
	"bufio"
	"os"
	"testing"
)

func TestMakeBoard(t *testing.T) {
	lines := [5]string{"67 97 50 51  1", "47 15 77 31 66", "24 14 55 70 52", "76 46 19 32 73", "34 22 54 75 17"}
	want := BingoBoard{
		[5]BingoNumber{
			{value: 67, marked: false},
			{value: 97, marked: false},
			{value: 50, marked: false},
			{value: 51, marked: false},
			{value: 1, marked: false},
		},
		[5]BingoNumber{
			{value: 47, marked: false},
			{value: 15, marked: false},
			{value: 77, marked: false},
			{value: 31, marked: false},
			{value: 66, marked: false},
		},
		[5]BingoNumber{
			{value: 24, marked: false},
			{value: 14, marked: false},
			{value: 55, marked: false},
			{value: 70, marked: false},
			{value: 52, marked: false},
		},
		[5]BingoNumber{
			{value: 76, marked: false},
			{value: 46, marked: false},
			{value: 19, marked: false},
			{value: 32, marked: false},
			{value: 73, marked: false},
		},
		[5]BingoNumber{
			{value: 34, marked: false},
			{value: 22, marked: false},
			{value: 54, marked: false},
			{value: 75, marked: false},
			{value: 17, marked: false},
		},
	}

	res := MakeBoard(lines)
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if res[i][j] != want[i][j] {
				t.Fatalf(`MakeBoard Failed (%d, %d) wanted %v, got %v`, i, j, want[i][j], res[i][j])
			}
		}
	}

}

func TestReadData(t *testing.T) {
	file, err := os.Open("data.txt")
	if err != nil {
		t.Fatalf("Error opening file")
	}
	scanner := bufio.NewScanner(file)
	numbers, boards := ReadData(scanner)
	t.Logf(`numbers:\n%v`, numbers)
	t.Logf(`boards:\n%v`, boards)
}

func TestMarkBoard(t *testing.T) {
	lines := [5]string{"67 97 50 51  1", "47 15 77 31 66", "24 14 55 70 52", "76 46 19 32 73", "34 22 54 75 17"}
	board := MakeBoard(lines)
	board.MarkValue(97)
	t.Logf(`Board: %v`, board)
	if !board[0][1].marked {
		t.Fatalf(`Value not marked! %v`, board)
	}
}

func TestCheckFalse(t *testing.T) {
	lines := [5]string{"67 97 50 51  1", "47 15 77 31 66", "24 14 55 70 52", "76 46 19 32 73", "34 22 54 75 17"}
	board := MakeBoard(lines)
	board.MarkValue(67)
	board.MarkValue(22)
	board.MarkValue(75)
	winner, total := board.Check()
	if winner {
		t.Fatalf(`Fake winner! Board: %v`, board)
	}
	if total != 1165-67-22-75 {
		t.Fatalf(`Wrong total! %v`, total)
	}
}

func TestCheckRow(t *testing.T) {
	lines := [5]string{"67 97 50 51  1", "47 15 77 31 66", "24 14 55 70 52", "76 46 19 32 73", "34 22 54 75 17"}
	board := MakeBoard(lines)
	wantTotal := 1165
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			board.MarkValue(board[i][j].value)
			wantTotal -= board[i][j].value
		}
		winner, total := board.Check()
		if !winner {
			t.Fatalf(`Shoulda Won! Board: %v`, board)
		}
		if total != wantTotal {
			t.Fatalf(`Wrong total! %v`, total)
		}
		board.Reset()
		wantTotal = 1165
	}
}

func TestCheckCol(t *testing.T) {
	lines := [5]string{"67 97 50 51  1", "47 15 77 31 66", "24 14 55 70 52", "76 46 19 32 73", "34 22 54 75 17"}
	board := MakeBoard(lines)
	wantTotal := 1165
	for j := 0; j < 5; j++ {
		for i := 0; i < 5; i++ {
			board.MarkValue(board[i][j].value)
			wantTotal -= board[i][j].value
		}
		winner, total := board.Check()
		if !winner {
			t.Fatalf(`Shoulda Won! Board: %v`, board)
		}
		if total != wantTotal {
			t.Fatalf(`Wrong total! %v`, total)
		}
		board.Reset()
		wantTotal = 1165
	}
}
