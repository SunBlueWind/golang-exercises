package popcount

// Popcount returns the population count (number of set bits of x
func Popcount(x uint64) (ret int) {
	var i uint64
	for ; i < 64; i++ {
		ret += int((x >> i) & 1)
	}
	return
}
