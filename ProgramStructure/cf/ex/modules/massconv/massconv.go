// Package massconv performs metric and imperial unit mass computations.
package massconv

import "fmt"

// Kilogram interprets floats as kilograms
type Kilogram float64

// Pound interprets floats as pounds
type Pound float64

// KToP converts a mass in kilograms to pounds
func KToP(k Kilogram) Pound { return Pound(k * 2.205) }

// PToK converts a mass in pounds to kilograms
func PToK(p Pound) Kilogram { return Kilogram(p * 0.45359237) }

func (k Kilogram) String() string { return fmt.Sprintf("%g kg", k) }
func (p Pound) String() string    { return fmt.Sprintf("%g lb", p) }
