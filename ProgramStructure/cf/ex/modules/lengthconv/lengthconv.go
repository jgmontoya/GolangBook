// Package lengthconv performs metric and imperial unit length computations.
package lengthconv

import "fmt"

// Meter interprets floats as meters
type Meter float64

// Feet interprets floats as feet
type Feet float64

// Yard interprets floats as yards
type Yard float64

// Inch interprets floats as inches
type Inch float64

// MToF converts a length in meters to feet
func MToF(m Meter) Feet { return Feet(m * 3.2808) }

// MToY converts a length in meters to yards
func MToY(m Meter) Yard { return Yard(m * 1.0936) }

// MToI converts a length in meters to inches
func MToI(m Meter) Inch { return Inch(m * 39.370) }

// FToM converts a length in feet to meters
func FToM(f Feet) Meter { return Meter(f * 0.3048) }

// YToM converts a length in yards to meters
func YToM(y Yard) Meter { return Meter(y * 0.9144) }

// IToM converts a length in inches to meters
func IToM(i Inch) Meter { return Meter(i * 25.4 / 1000) }

func (m Meter) String() string { return fmt.Sprintf("%g m", m) }
func (f Feet) String() string  { return fmt.Sprintf("%g ft", f) }
func (i Inch) String() string  { return fmt.Sprintf("%g in", i) }
func (y Yard) String() string  { return fmt.Sprintf("%g yd", y) }
