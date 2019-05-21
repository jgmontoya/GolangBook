// Package popcount2 returns the population count (i.e. number of set bits) in a uint64 number.
// Counts bits by shifting it argument through 64 bit positions.
package popcount2

// PopCount returns the population count (number of set bits) of x.
func PopCount(x uint64) int {
	var result uint64
	for i := uint64(0); i < 64; i++ {
		result += (x & (1 << i))
	}
	return int(result)
}
