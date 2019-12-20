package popcount

// pc[i] is the population count of i
var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

// Popcount returns the population count (number of set bits of x
func Popcount(x uint64) (ret int) {
	var i uint64
	for ; i < 8; i++ {
		ret += int(pc[byte(x>>(i*8))])
	}
	return
}
