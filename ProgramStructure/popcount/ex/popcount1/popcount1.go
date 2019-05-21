// Package popcount1 returns the population count (i.e. number of set bits) in a uint64 number.
// Uses a loop instead of a simple expression.
package popcount1

// pc[i] is the population count of i.
var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

// PopCount returns the population count (number of set bits) of x.
func PopCount(x uint64) int {
	var result byte
	for i := uint64(0); i < 8; i++ {
		result += pc[byte(x>>(i*8))]
	}
	return int(result)
}
