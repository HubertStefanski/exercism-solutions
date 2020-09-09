package grains

// Local benchmark results
//goos: linux
//goarch: amd64
//pkg: grains
//BenchmarkSquare
//BenchmarkSquare-8   	 2463411	       488 ns/op	     304 B/op	       8 allocs/op
//BenchmarkTotal
//BenchmarkTotal-8    	13781620	        80.6 ns/op	       0 B/op	       0 allocs/op
//PASS
//ok  	grains	2.898s

import "fmt"

var (
	numOfSquares = 64
)

// return total number of grains from the whole board, sum of all square grain values
func Total() uint64 {
	var sum uint64
	for i := numOfSquares / numOfSquares; i <= numOfSquares; i++ {
		sSum, _ := Square(i)
		sum = sum + sSum
	}

	return sum
}

// get the grain value of this square
func Square(input int) (uint64, error) {
	// check that the input is in bounds of the board
	if input <= 0 || input > numOfSquares {
		return 0, fmt.Errorf("invalid input, input either out of range or non-positive integer %v", input)
	}
	return 1 << (input - 1), nil
}
