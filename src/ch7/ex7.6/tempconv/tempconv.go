// Package tempconv performs Celsius, Fahrenheit, and Kelvin conversions
package tempconv

import (
	"flag"
	"fmt"
)

// Celsius represents temperature in Celsius
type Celsius float64

// Fahrenheit represents temperature in Fahrenheit
type Fahrenheit float64

// Kelvin represents temperature in Kelvin
type Kelvin float64

type celciusFlag struct{ Celsius }

func (f *celciusFlag) Set(s string) error {
	var unit string
	var value float64
	fmt.Sscanf(s, "%f%s", &value, &unit)
	switch unit {
	case "C", "°C":
		f.Celsius = Celsius(value)
		return nil
	case "F", "°F":
		f.Celsius = FToC(Fahrenheit(value))
		return nil
	case "K", "°K":
		f.Celsius = KToC(Kelvin(value))
		return nil
	}
	return fmt.Errorf("invalid temperature %q", s)
}

// CelsiusFlag defines a Celsius flag with the specified name,
// default value, and usage, and returns the address of the flag variable.
// The flag argument must have a quantity and a unit, e.g. "100°C".
func CelsiusFlag(name string, value Celsius, usage string) *Celsius {
	f := celciusFlag{value}
	flag.CommandLine.Var(&f, name, usage)
	return &f.Celsius
}

const (
	// AbsoluteZeroC is the absolute zero temperature in Celsius
	AbsoluteZeroC Celsius = -273.15
	// FreezingC is the freezing temperatur of water in Celsius
	FreezingC Celsius = 0
	// BoilingC is the boiling temperatur of water in Celsius
	BoilingC Celsius = 100
)

func (c Celsius) String() string    { return fmt.Sprintf("%g°C", c) }
func (f Fahrenheit) String() string { return fmt.Sprintf("%g°F", f) }
func (k Kelvin) String() string     { return fmt.Sprintf("%g°K", k) }
