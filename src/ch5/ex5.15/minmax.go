package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
)

func min(x float64, val ...float64) float64 {
	res := x
	for _, v := range val {
		res = math.Min(res, v)
	}
	return res
}

func max(x float64, val ...float64) float64 {
	res := x
	for _, v := range val {
		res = math.Max(res, v)
	}
	return res
}

func main() {
	if len(os.Args) > 1 {
		args := make([]float64, len(os.Args)-1)
		for i := range args {
			if val, err := strconv.ParseFloat(os.Args[i+1], 64); err == nil {
				args[i] = val
			}
		}
		minVal := min(args[0], args[1:]...)
		maxVal := max(args[0], args[1:]...)
		fmt.Printf("Min value is: %f\n", minVal)
		fmt.Printf("Max value is: %f\n", maxVal)
	}
}
