package main

import (
	"fmt"
	"math"
)

func main() {
	var B int
	pntr := &B
	B = 54
	*pntr = 66
	fmt.Println(*pntr)

	var value float64
	fmt.Scan(&value)
	R := &value
	var O float64
	O = math.Pow(*R, 2) * 3.14
	fmt.Println(O)
}
