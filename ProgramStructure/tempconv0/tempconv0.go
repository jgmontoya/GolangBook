// Package tempconv performs Celsius and Fahrenheit temperature computations.
package tempconv

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

// CToF converts celsius degrees to fahrenheit degrees.
func CToF(c Celsius) Fahrenheit { return Fahrenheit(c*9/5 + 32) }

// FToC converts fahrenheit degrees to celsius degrees.
func FToC(f Fahrenheit) Celsius { return Celsius((f - 32) * 5 / 9) }
