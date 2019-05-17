// Package tempconv performs Celsius and Fahrenheit temperature computations.
package tempconv

import "fmt"

// Celsius interprets floats as Celsius degrees
type Celsius float64

// Fahrenheit interprets floats as Fahrenheit degrees
type Fahrenheit float64

const (
	// AbsoluteZeroC is the temperature in celsius of the absolute zero
	AbsoluteZeroC Celsius = -273.15
	// FreezingC is the temperature at which water freezes under standard conditions
	FreezingC Celsius = 0
	// BoilingC is the temperature at which water boils under standard conditions
	BoilingC Celsius = 100
)

func (c Celsius) String() string    { return fmt.Sprintf("%g°C", c) }
func (f Fahrenheit) String() string { return fmt.Sprintf("%g°F", f) }
