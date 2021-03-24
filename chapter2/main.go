package main

import (
	"fmt"
	"./sports"
)

func main() {
	fmt.Println(App())
	fmt.Println(sports.Soccer())
	fmt.Println(sports.Baseball())
	fmt.Println(sports.Basketball())
}