/*bool

string

int  int8  int16  int32  int64
uint uint8 uint16 uint32 uint64 uintptr

byte // alias for uint8

rune // alias for int32
     // represents a Unicode code point

float32 float64

complex64 complex128*/

package main

import (
	"fmt"
	"math"
	"math/cmplx"
)

var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

func main() {
	fmt.Printf("Value %v has Type %T\n", ToBe, ToBe)
	fmt.Printf("Value %v has Type %T\n", MaxInt, MaxInt)
	fmt.Printf("Value %v has Type %T\n", z, z)

	/*Variables declared without an explicit initial value are given their zero value.
		  The zero value is:
		      0 for numeric types,
		      false for the boolean type, and
	          "" (the empty string) for strings. */

	var i int
	var f float64
	var b bool
	var s string
	fmt.Printf("%v %v %v %q\n", i, f, b, s)

	//The expression T(v) converts the value v to the type T.
	var x, y int = 3, 4
	var fl float64 = math.Sqrt(float64(x*x + y*y))
	var z uint = uint(fl)
	fmt.Println(x, y, z)
	//Or, put more simply:
	j := 42
	flt := float64(j)
	u := uint(flt)
	fmt.Println(j, flt, u)

	/*When declaring a variable without specifying an explicit type (either by using the := syntax or var = expression syntax),
	the variable's type is inferred from the value on the right hand side.*/
	in := 26
	floa := 3.14
	g := 0.867 + 0.5i
	fmt.Printf("\nin is of type %T ", in)
	fmt.Printf("\nfloa is of type %T ", floa)
	fmt.Printf("\ng is of type %T ", g)

	/*Constants
	Constants are declared like variables, but with the const keyword.
	Constants can be character, string, boolean, or numeric values.
	Constants cannot be declared using the := syntax.*/

	const World = "She"
	fmt.Println("\nHello ", World)
	fmt.Println("\nHappy ", Pi, " Day")

	const Truth = true
	fmt.Println("\nGo rule? ", Truth)
	fmt.Println()

	fmt.Println(needInt(Small))
	fmt.Println(needFloat(Small))
	fmt.Println(needFloat(Big))
}

const Pi = 3.14

const (
	// Create a huge number by shifting a 1 bit left 100 places.
	// In other words, the binary number that is 1 followed by 100 zeroes.
	Big = 1 << 100
	// Shift it right again 99 places, so we end up with 1<<1, or 2.
	Small = Big >> 99
)

func needInt(x int) int {
	return x*10 + 1
}

func needFloat(x float64) float64 {
	return x * 0.1
}
