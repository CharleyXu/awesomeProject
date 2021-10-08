package main

import (
	"fmt"
	"strings"
)

var a = "123"

var b int

var c bool

func main() {
	const d string = "d"
	const e int8 = 1
	const (
		Unknown = 0
		Female = 1
		Male = 2
	)
	f := "str"
	println(f)
	println(b)
	println(d)
	println(e)
	fmt.Printf("%v %v \n", b, c)
	println(hello("b"))
	fmt.Println(strings.Contains(a, "1"))
}

func hello(name string) string {
	return "a" + name
}
