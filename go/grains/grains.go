package grains

// Local benchmark results
//goos: linux
//goarch: amd64
//pkg: grains
//BenchmarkSquare
//BenchmarkSquare-8   	 1983249	       600 ns/op	     304 B/op	       8 allocs/op
//BenchmarkTotal
//BenchmarkTotal-8    	 2066004	       595 ns/op	       0 B/op	       0 allocs/op
//PASS
//ok  	grains	3.622s

import "fmt"

var (
	numOfSquares         = 64
	startValue           = 1
	multiplyConst uint64 = 2
	SquareMap = populateSquares()
)

// return total number of grains from the whole board, sum of all square grain values
func Total() uint64 {
	var sum uint64
	for i := numOfSquares / numOfSquares; i <= numOfSquares; i++ {
		sum = sum + SquareMap[i]
	}
	return sum
}

// get the grain value of this square
func Square(input int) (uint64, error) {

	// check that the input is in bounds of the board
	if input <= 0 || input > numOfSquares {
		return SquareMap[input], fmt.Errorf("invalid input, input either out of range or non-positive integer %v", input)
	}

	return SquareMap[input], nil
}

// pre-populate squares with grain values, easier for later lookups
func populateSquares() map[int]uint64 {
	var m = make(map[int]uint64, numOfSquares)

	for i := numOfSquares / numOfSquares; i <= numOfSquares; i++ {
		// initialize the value
		if i <= 1 {
			m[i] = uint64(startValue)
			continue
		}
		// get the value from the previous square
		prevGrainVal := m[i-1]
		// multiply previous value by multiplyConst and get the current Square grain value
		currVal := prevGrainVal * multiplyConst
		// assign current value
		m[i] = currVal
	}
	return m
}
