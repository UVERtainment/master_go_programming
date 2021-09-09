package main

import (
	"fmt"
)

//var version = "3.1"
type duration int
type mile float64
type kilometer float64

func main() {

	//050 named types and aliases
	//ex1
	// var hour duration

	// fmt.Printf("hour value = %v, hour type = %T\n", hour, hour)

	// hour = 3600
	// fmt.Println(hour)

	//ex2
	// var hour duration = 3600
	// minute := 60
	// fmt.Println(int(hour) != minute)

	//ex3
	const m2km = 1.609
	var mileBerlinToParis mile = 655.3
	var kmBerlinToParis kilometer
	kmBerlinToParis = kilometer(mileBerlinToParis) * m2km
	fmt.Println(kmBerlinToParis)

	//049 operators and conversions
	//ex1
	// var i int = 3
	// var f float64 = 3.2

	// var x = float64(i)
	// var y = int(f)

	// fmt.Printf("x type = %T, x value = %v\n", x, x)
	// fmt.Printf("y type = %T, y value = %v", y, y)

	//ex2
	// var i = 3
	// var f = 3.2
	// var s1, s2 = "3.14", "5"

	// var iStr = strconv.Itoa(i)
	// fmt.Printf("iStr type = %T, iStr value = %v\n", iStr, iStr)

	// var s2Int, _ = strconv.Atoi(s2)
	// fmt.Printf("s2Int type = %T, s2Int value = %v\n", s2Int, s2Int)

	// var fStr = fmt.Sprintf("%f", f)
	// fmt.Printf("fStr type = %T, fStr value = %v\n", fStr, fStr)

	// var s1Float, _ = strconv.ParseFloat(s1, 64)
	// fmt.Printf("s1Float type = %T, s1Float value = %v\n", s1Float, s1Float)

	//ex3
	// x, y := 4, 5.1

	// z := x * int(y)
	// fmt.Println(z)

	// a := 5
	// b := 6.2 * float64(a)
	// fmt.Println(b)

	//ex4
	// distanceInKm := 149.6
	// speedInM := 299792458
	// speedInKm := speedInM / 1000
	// time := distanceInKm / float64(speedInKm)

	// fmt.Println(time)

	//ex5
	// x, y := 0.1, 5
	// var z float64

	// // Write the correct logical operator (&&, || , !)
	// // inside the expression so that result1 will be false and result2 will be true.

	// result1 := x < z && !(int(x) != int(z))
	// fmt.Println(result1)

	// result2 := y == 1*5 || int(z) == 0
	// fmt.Println(result2)

	//048 package fmt
	//ex1
	// x, y, z := 10, 15.5, "Gophers"
	// score := []int{10, 20, 30}

	// fmt.Printf("x = %d\ny = %f\nz = %s\n", x, y, z)
	// fmt.Printf("score is %#v\n", score)
	// fmt.Printf("z quoted = %q\n", z)
	// fmt.Printf("x = %v\ny = %v\nz = %v\nscore = %v\n", x, y, z, score)
	// fmt.Printf("type of y = %T, type of score = %T", y, score)

	//ex2
	// const x float64 = 1.422349587101
	// fmt.Printf("x 4 decimal = %.4f", x)

	//ex3
	// shape := "circle"
	// radius := 3.2
	// const pi float64 = 3.14159
	// circumference := 2 * pi * radius

	// fmt.Printf("Shape: %q\n", shape)
	// fmt.Printf("Circle's circumference with radius %f is %f\n", radius, circumference)
	// _ = shape

	//047 Constant
	//ex1
	// const (
	// 	daysWeek   = 7
	// 	lightSpeed = 299792458
	// 	pi         = 3.14159
	// 	secPerDay  = 86400
	// 	daysYear   = 365
	// )

	// fmt.Printf("There are %d seconds in a year.", secPerDay*daysYear)

	//ex4
	// const x int = 10

	// // declaring a constant of type slice int ([]int)
	// const m = []int{1: 3, 4: 5, 6: 8}	// cannot declare a const of slice int!!
	// _ = m

	//ex5
	// const a = 7
	// const b = 5.6
	// const c = a * b

	// // x := 8
	// // const xc int = x	//cannot use a variable to declare a const

	// // const noIPv6 = math.Pow(2, 128)	//cannot initialize a const at run time

	//ex6
	// const (
	// 	Jun = iota + 6
	// 	Jul
	// 	Aug
	// )

	// fmt.Println(Jun, Jul, Aug)

	//046 Declare Variables
	//ex1
	// var a int
	// var b float64
	// var c bool
	// var d string

	// x := 20
	// y := 15.5
	// z := "Gopher!"

	// fmt.Println(a, b, c, d, x, y, z)

	//ex2
	// var (
	// 	a int
	// 	b float64
	// 	c bool
	// 	d string
	// )

	// x, y, z := 20, 15.5, "Gopher!"

	// _, _, _, _, _, _, _ = a, b, c, d, x, y, z

	//ex3
	// var a float64 = 7.1

	// x, y := true, 3.7

	// //a, x := 5.5, false
	// a, x = 5.5, false

	// _, _, _ = a, x, y

	//ex4
	// name := "Golang"
	// fmt.Println(name)
}
