
package main

import (
	"fmt"
	"math/rand"
	"math"
)

var c, python, java bool

func add(x, y int) int {
	return x + y
}

func swap(x, y string) (string, string) {
	return y, x
}

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}
func main() {
        rand.Seed(4)
	fmt.Println("My favorite number is", rand.Intn(10))
        fmt.Printf("Now you have %g problems.\n", math.Sqrt(7))
        fmt.Println(math.Pi)
        fmt.Println(add(42, 13))
        a, b := swap("hello", "world")
	fmt.Println(a, b)
        //#"naked" return.
        fmt.Println(split(17))
        //#variables declared
        var i int
	fmt.Println(i, c, python, java)
        //#short assignment
        //#:= short assignment
        x := 67
	fmt.Println(x)
}

