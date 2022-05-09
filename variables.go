package main

import "fmt"

// variable block
/*
var (
	actorName    string = "Elisabeth  Sladen"
	companion    string = "Sarah Jane Smith"
	doctorNumber int    = 3
	season       int    = 11
)
*/

var x string = "Shadow variable"

func main() {

	// Shadow Variable

	fmt.Println(x)

	// first way of creating a variable
	var x string
	x = "hello, playground"
	fmt.Printf("%v, %T\n", x, x)

	// second way of creating a variable
	var y string = "Greetings Traveler"
	fmt.Printf("%v, %T\n", y, y)

	// third way of creating a variable
	z := 43
	fmt.Printf("%v, %T\n", z, z)

	// converting z int to f float
	var f float64
	f = float64(z)
	fmt.Printf("%v, %T\n", f, f)

}
