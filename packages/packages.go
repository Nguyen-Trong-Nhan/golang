package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

// var c, python, java bool
var i, j int = 1, 2

func main() {
	fmt.Print("RANDOM NUMBER\n")
	// For example, `rand.Intn` returns a random `int` n,
	// `0 <= n < 100`.
	fmt.Print(rand.Intn(100), ",")
	fmt.Print(rand.Intn(100))
	fmt.Println()

	// `rand.Float64` returns a `float64` `f`,
	// `0.0 <= f < 1.0`.
	fmt.Println(rand.Float64())

	// This can be used to generate random floats in
	// other ranges, for example `5.0 <= f' < 10.0`.
	fmt.Print((rand.Float64()*5)+5, ",")
	fmt.Print((rand.Float64() * 5) + 5)
	fmt.Println()

	// The default number generator is deterministic, so it'll
	// produce the same sequence of numbers each time by default.
	// To produce varying sequences, give it a seed that changes.
	// Note that this is not safe to use for random numbers you
	// intend to be secret, use `crypto/rand` for those.
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)

	// Call the resulting `rand.Rand` just like the
	// functions on the `rand` package.
	fmt.Print(r1.Intn(100), ",")
	fmt.Print(r1.Intn(100))
	fmt.Println()

	// If you seed a source with the same number, it
	// produces the same sequence of random numbers.
	s2 := rand.NewSource(42)
	r2 := rand.New(s2)
	fmt.Print(r2.Intn(100), ",")
	fmt.Print(r2.Intn(100))
	fmt.Println()

	s3 := rand.NewSource(42)
	r3 := rand.New(s3)
	fmt.Print(r3.Intn(100), ",")
	fmt.Print(r3.Intn(100))
	fmt.Println()

	s4 := rand.NewSource(time.Now().UnixNano())
	r4 := rand.New(s4)
	fmt.Println("My favorite number is: ", r4.Intn(10))
	fmt.Println()

	//Export names
	//In Go, a name is exported if it begins with a capital letter.
	//For example, Pizza is an exported name, as is Pi, which is exported from the math package.
	//pizza and pi do not start with a capital letter, so they are not exported.

	fmt.Print("EXPORT NAMES\n")
	fmt.Printf("Now you have %g problems\n", math.Sqrt(7))
	fmt.Println("Pi number is ", math.Pi)
	fmt.Println()

	//FUNCTION
	fmt.Print("FUNCTIONS\n")
	fmt.Println(add(1, 1))
	fmt.Println(substract(1, 1))
	a := "hello"
	b := "world"
	//a, b := swap("hello", "world")
	fmt.Println(swap(a, b))
	fmt.Println(split(10))

	//VARIBLES
	fmt.Print("VARIBLES\n")
	// var i int
	// fmt.Println(i, c, python, java)
	var c, python, java = true, false, "no!"
	fmt.Println(i, j, c, python, java)

}

func add(x int, y int) int {
	return x + y
}

// /When two or more consecutive named function parameters share a type, you can omit the type from all but the last.
func substract(x, y int) int {
	return x - y
}

//A function can return any number of results.
//The swap function returns two strings.
func swap(x, y string) (string, string) {
	return y, x
}

//A return statement without arguments returns the named return values. This is known as a "naked" return.
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

//The var statement declares a list of variables; as in function argument lists, the type is last.
//A var statement can be at package or function level. We see both in this example.
