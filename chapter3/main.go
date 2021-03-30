package main

import (
	"./foo"
	"fmt"
	"math"
)

var (
	a, b int8
	c    string
	d    uint16
	e    float64
	f    rune
	g    [5]int
)

const (
	A = "A"
	B = 1
	C
)

func minus(x, y int) int {
	return x - y
}

func plus(x, y, z int) (int, int) {
	v := x + z
	w := y + z
	return v, w
}

func doSomething() (x, y int) {
	return
}

func returnFunc() func() {
	return func() {
		fmt.Println("return")
	}
}

func callFunc(f func()) {
	f()
}

func later() func(string) string {
	var store string
	return func(next string) string {
		s := store
		store = next
		return s
	}
}

func generator() func() int {
	i := 0
	return func() int {
		i += 1
		return i
	}
}

func runDefer() {
	defer func() {
		fmt.Println("Defer1")
		fmt.Println("Defer2")
		fmt.Println("Defer3")
	}()
	fmt.Println("Done")
}

func init() {
	fmt.Println("Init")
}

func main() {
	a = 127
	b = -128
	c = "C"
	d = 65535
	e = 1.0000000000000001
	f = 'あ'

	g[0] = 1
	h := [...]string{"h"}
	i := h // コピー
	j := minus(5, 4)
	k, l := plus(5, 4, 3)
	m, n := doSomething()
	o := func(x, y int) int { return x + y }
	returnFunc()()
	callFunc(func() {
		fmt.Println("call")
	})
	p := later()
	q := generator()

	fmt.Printf(
		"a=%d b=%d c=%s d=%d e=%v f=%v MaxInt64=%d MinInt64=%d\n",
		a, b, c, d, e, f, math.MaxInt64, math.MinInt64,
	) // => a=127, b=-128, c=C, d=65535, e=1, f=12354, MaxInt64=9223372036854775807, MinInt64=-9223372036854775808
	fmt.Printf(
		"g=%v h=%v i=%v j=%v k=%v l=%v m=%v n=%v o=%v, p1=%v p2=%v q1=%v q2=%v\n",
		g, h, i, j, k, l, m, n, o(1, 2), p("A"), p("B"), q(), q(),
	) // => g=[1 0 0 0 0] h=[h] i=[h] j=1 k=8 l=7 m=0 n=0 o=3, p1= p2=A q1=1 q2=2
	fmt.Printf(
		"A=%v B=%v C=%v\n",
		A, B, C,
	) // => A=A B=1 C=1
	fmt.Println(foo.DoSomething()) // => A=A B=1 C=1

	for i := 1; i < 100; i++ {
		if i%15 == 0 {
			fmt.Println("Fizz Buzz")
		} else if i%3 == 0 {
			fmt.Println("Fizz")
		} else if i%5 == 0 {
			fmt.Println("Buzz")
		} else {
			fmt.Println(i)
		}

		switch {
		case i%15 == 0:
			fmt.Println("Fizz Buzz")
		case i%3 == 0:
			fmt.Println("Fizz")
		case i%5 == 0:
			fmt.Println("Buzz")
		default:
			fmt.Println(i)
		}
	}

LOOP:
	for i := 1; i < 10; i++ {
		for j := 1; j < 10; j++ {
			if i%j == 0 {
				fmt.Printf("i=%d j=%d", i, j)
				break LOOP
			}
		}
	}

	N := 1
	switch N {
	case 1:
		fmt.Println("1")
	case 2, 3:
		fmt.Println("2, 3")
	case 4:
		fmt.Println("4")
	default:
		fmt.Println(i)
	}

	S := "あ"
	switch S {
	case "あ":
		S += "い"
		fallthrough
	case "い":
		S += "う"
		fallthrough
	case "う":
		S += "え"
		fallthrough
	default:
		S += "お"
	}
	runDefer()

}
