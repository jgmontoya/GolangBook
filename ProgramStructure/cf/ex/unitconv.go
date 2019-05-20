// Unitconv converts its numeric argument to metric and imperial units of the selected dimension

package main

import (
	"GolangBook/ProgramStructure/cf/ex/modules/lengthconv"
	"GolangBook/ProgramStructure/cf/ex/modules/massconv"
	"GolangBook/ProgramStructure/cf/ex/modules/tempconv"
	"flag"
	"fmt"
)

var dim = flag.String("dim", "mass", "dimension (mass, length, temp)")
var n = flag.Float64("n", 0.0, "number to convert")

func main() {
	flag.Parse()
	switch *dim {
	case "mass":
		k := massconv.Kilogram(*n)
		p := massconv.Pound(*n)
		fmt.Printf("%s = %s, %s = %s\n", k, massconv.KToP(k), p, massconv.PToK(p))

	case "temp":
		f := tempconv.Fahrenheit(*n)
		c := tempconv.Celsius(*n)
		fmt.Printf("%s = %s, %s = %s\n", f, tempconv.FToC(f), c, tempconv.CToF(c))

	case "length":
		m := lengthconv.Meter(*n)
		f := lengthconv.Feet(*n)
		y := lengthconv.Yard(*n)
		i := lengthconv.Inch(*n)
		fmt.Printf("%s = %s = %s = %s\n%s = %s\n%s = %s\n%s = %s\n",
			m, lengthconv.MToF(m), lengthconv.MToY(m), lengthconv.MToI(m),
			f, lengthconv.FToM(f), y, lengthconv.YToM(y), i, lengthconv.IToM(i))
	}
}
