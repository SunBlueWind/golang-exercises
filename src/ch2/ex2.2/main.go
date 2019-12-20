package main

import (
	"fmt"
	"os"
	"strconv"

	"ch2/ex2.2/lenconv"
	"ch2/ex2.2/tempconv"
	"ch2/ex2.2/weightconv"
)

func main() {
	for _, arg := range os.Args[1:] {
		n, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%v\n", err)
			os.Exit(1)
		}
		celsius := tempconv.Celsius(n)
		fahrenheit := tempconv.Fahrenheit(n)
		meter := lenconv.Meter(n)
		feet := lenconv.Feet(n)
		kilogram := weightconv.Kilogram(n)
		pound := weightconv.Pound(n)

		fmt.Printf("%s = %s, %s = %s\n",
			fahrenheit, tempconv.FToC(fahrenheit),
			celsius, tempconv.CToF(celsius))
		fmt.Printf("%s = %s, %s = %s\n",
			meter, lenconv.MToF(meter),
			feet, lenconv.FToM(feet))
		fmt.Printf("%s = %s, %s = %s\n",
			kilogram, weightconv.KToP(kilogram),
			pound, weightconv.PToK(pound))
	}
}
