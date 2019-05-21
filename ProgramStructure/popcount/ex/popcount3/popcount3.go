// Package popcount3 returns the population count (i.e. number of set bits) in a uint64 number.
// Counts bits using the fact that x & (x - 1) clears the rightmost non-zero bit of x.
package popcount3

// PopCount returns the population count (number of set bits) of x.
func PopCount(x uint64) int {
	var result int
	for x != 0 {
		x = x & (x - 1)
		result++
	}
	return result
}
