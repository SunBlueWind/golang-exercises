package popcount

// Popcount returns the population count (number of set bits of x
func Popcount(x uint64) (ret int) {
	if x == 0 {
		return 0
	}
	return 1 + Popcount(x&(x-1))
}
