// Package weightconv performs kilogram and pound conversions
package weightconv

import "fmt"

type Kilogram float64
type Pound float64

func (k Kilogram) String() string { return fmt.Sprintf("%g kg", k) }
func (p Pound) String() string    { return fmt.Sprintf("%g lb", p) }

const multiplier = 2.205

// KToP converts a kilogram weight to pound
func KToP(k Kilogram) Pound { return Pound(k * multiplier) }

// PToK converts a pound weight to kilogram
func PToK(p Pound) Kilogram { return Kilogram(p / multiplier) }
