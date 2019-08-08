
package main

import "fmt"

func main() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)

	/* forcontinued */
	sum2 := 1
	for ; sum2 < 1000; {
		sum2 += sum2
	}
	fmt.Println(sum2)
        /* for go while */
	sum := 1
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)
	/* forever = while (true) */
	for {
	}
}

