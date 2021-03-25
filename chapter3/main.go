package main

import (
	"fmt"
	"math"
)

var (
	a, b int8
	c string
	d uint16
	e float64
	f rune
)

func main() {
	a = 127
	b = -128
	c = "C"
	d = 65535
	e = 1.0000000000000001
	f = 'あ'

	fmt.Printf(
		"a=%d, b=%d, c=%s, d=%d, e=%v, f=%v, MaxInt64=%d, MinInt64=%d\n",
		a, b, c, d, e, f, math.MaxInt64, math.MinInt64,
	) // => a=127, b=-128, c=C, d=65535, e=1, f=12354, MaxInt64=9223372036854775807, MinInt64=-9223372036854775808
}