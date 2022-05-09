package main

import (
	"fmt"
)

func main() {
	fmt.Println("Primitive Types")

	// boolTypes()
	// integerTypes()
	// floatingTypes()
	// runeType()
}

func runeType() {
	var r rune = 'a'
	fmt.Printf("%v, %T\n", r, r)
}

func floatingTypes() {

	var n float64 = 3.14
	n = 3.7e72
	n = 2.1e14

	fmt.Printf("%v, %T", n, n)

}

func integerTypes() {

	// integers
	a := 10 // 1010
	b := 3  // 0011

	// math operator
	fmt.Println("Math operator")
	fmt.Println(a + b)
	fmt.Println(a - b)
	fmt.Println(a * b)
	fmt.Println(a / b)
	fmt.Println(a % b)

	// bit operator
	fmt.Println("Bit operator")
	fmt.Println(a & b)  // 0010
	fmt.Println(a | b)  // 1011
	fmt.Println(a ^ b)  // 1001
	fmt.Println(a &^ b) // 0100

	// bit shifting
	fmt.Println("Bit shifting")
	a = 8               // 2^3
	fmt.Println(a << 3) // 2^3 * 2^3 = 2^6
	fmt.Println(a >> 3) // 2^3 / 2^3 = 2^0

	var x int = 8
	var y int8 = 16

	// Gottta convert it first go is strict
	fmt.Println("Int Conversion")
	fmt.Println(x + int(y))

}

func boolTypes() {
	fmt.Println("Bool types")

	// Boolean types
	n := 1 == 1
	m := 1 == 2

	// default bool var
	var p bool

	fmt.Printf("%v, %T\n", n, n)
	fmt.Printf("%v, %T\n", m, m)
	fmt.Printf("%v, %T\n", p, p)

	// Numeric types

}
