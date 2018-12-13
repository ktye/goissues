package main

import (
	"fmt"
	"math/cmplx"
)

func main() {
	x := complex128(1)
	y := -x
	fmt.Println(y, cmplx.Sqrt(y))

	z := -complex(1, 0)
	fmt.Println(z, cmplx.Sqrt(z))
}
