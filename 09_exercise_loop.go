
package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {

	delta := 0.05
	zprima := 0.0
	z := float64(1)
	i := 0
	//for  i < 30 {
	for (math.Abs(zprima-z)>delta) {
	   zprima = z
	   tmp := (z*z) - x
	   z = z - (tmp/2*z)
	   i++
	   fmt.Println("iteracion", i)
	   fmt.Println("imprimiendo z", z)
	   fmt.Println("imprimiendo zprima",zprima)
	   fmt.Println("DIFF",math.Abs(zprima-z))
	}
	return z
}

func main() {
	fmt.Println(Sqrt(2))
}
