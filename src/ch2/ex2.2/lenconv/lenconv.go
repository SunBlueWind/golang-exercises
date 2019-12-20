// Package lenconv performs feet and meters conversions
package lenconv

import "fmt"

type Meter float64
type Feet float64

func (m Meter) String() string { return fmt.Sprintf("%g m", m) }
func (f Feet) String() string  { return fmt.Sprintf("%g ft", f) }

const multiplier = 3.281

// MToF converts a meter length to feet
func MToF(m Meter) Feet { return Feet(m * multiplier) }

// FToM converts a feet length to meter
func FToM(f Feet) Meter { return Meter(f / multiplier) }
