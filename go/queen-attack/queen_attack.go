package queenattack

import (
	"fmt"
	"sort"
	"strings"
)

// Local benchmark results
//goos: linux
//goarch: amd64
//pkg: queenattack
//BenchmarkCanQueenAttack
//BenchmarkCanQueenAttack-8   	   19803	     60454 ns/op	   40277 B/op	    1326 allocs/op
//PASS
//ok  	queenattack	1.812s


var letters = []string{"a", "b", "c", "d", "e", "f", "g", "h"}
var numbers = []string{"1", "2", "3", "4", "5", "6", "7", "8"}

func CanQueenAttack(w, b string) (bool, error) {
	board := createBoardWithAlgebraicNotation()
	// check if the passed pieces are actually on the board
	if !contains(board, w) || !contains(board, b) || w == b {
		return false, fmt.Errorf("error, invalid position of one of the pieces: b=%v w=%v", b, w)
	}

	splitSliceW := strings.Split(w, "")
	splitSliceB := strings.Split(b, "")

	// Check horizontals and verticals for attack
	for _, wLocSplit := range splitSliceW {
		for _, bLocSplit := range splitSliceB {
			if wLocSplit == bLocSplit {
				return true, nil
			}
		}
	}
	// Check diagonals for attack
	diagonalLocationsB := getDiagonalsForLocation(b)
	if !contains(diagonalLocationsB, w){
		return false, nil
	}

	return true, nil
}

func createBoardWithAlgebraicNotation() []string {
	var board []string

	for _, letter := range letters {
		for _, number := range numbers {
			board = append(board, letter+number)
		}
	}

	return board
}

func getDiagonalsForLocation(inLoc string) []string {
	var diagonalsOfLocation []string

	var inLocLetterIndex int

	var inLocNumberIndex int

	splitLoc := strings.Split(inLoc, "")
	for _, n := range splitLoc {
		if contains(numbers, n) {
			inLocNumberIndex = sort.SearchStrings(numbers, n)
		} else if contains(letters, n) {
			inLocLetterIndex = sort.SearchStrings(letters, n)
		}
	}

	for i := 0; i <= 8; i++ {
		if !(inLocLetterIndex+i >= len(letters)) && !(inLocNumberIndex+i >= len(numbers)) {
			diagonalsOfLocation = append(diagonalsOfLocation, letters[inLocLetterIndex+i]+numbers[inLocNumberIndex+i])
		}

		if inLocLetterIndex-i >= 0 && !(inLocNumberIndex+i >= len(numbers)) {
			diagonalsOfLocation = append(diagonalsOfLocation, letters[inLocLetterIndex-i]+numbers[inLocNumberIndex+i])
		}

		if !(inLocLetterIndex+i >= len(letters)) && inLocNumberIndex-i >= 0 {
			diagonalsOfLocation = append(diagonalsOfLocation, letters[inLocLetterIndex+i]+numbers[inLocNumberIndex-i])
		}

		if inLocLetterIndex-i >= 0 && inLocNumberIndex-i >= 0 {
			diagonalsOfLocation = append(diagonalsOfLocation, letters[inLocLetterIndex-i]+numbers[inLocNumberIndex-i])
		}
	}

	return diagonalsOfLocation
}

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}

	return false
}
