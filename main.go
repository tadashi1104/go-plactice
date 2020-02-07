package main

import (
	"fmt"
	"math"
	"math/cmplx"
	"math/rand"
	"time"
)

var c, python, java bool
var i, j int = 1, 2

// まとめて定義
var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

func main() {
	fmt.Printf("Hello world\n")
	fmt.Println("The time is", time.Now())
	fmt.Println("My favorite number is", rand.Intn(10))
	fmt.Printf("Now you have %g problems.", math.Sqrt(7))
	fmt.Println(math.Pi)
	fmt.Println(add(42, 13))
	a, b := swap("hello", "world")
	fmt.Println(a, b)
	fmt.Println(split(17))
	var i int
	fmt.Println(i, c, python, java)
	var c, python, java = true, false, "no!"
	fmt.Println(i, j, c, python, java)
	k := 3
	d, qython, kava := true, false, "Oh!"
	fmt.Println(i, j, k, d, qython, kava)
	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value: %v\n", z, z)
	var inta int
	var f float64
	var boola bool
	var s string
	fmt.Printf("%v %v %v %q\n", inta, f, boola, s)
	typeConversions()
	constants()
}

// Pi
const Pi = 3.14

func constants() {
	const World = "世界"
	fmt.Println("Hello", World)
	fmt.Println("Happy", Pi, "Day")

	const Truth = true
	fmt.Println("Go rules?", Truth)
}

func typeConversions() {
	var x, y int = 3, 4
	var f float64 = math.Sqrt(float64(x*x + y*y))
	var z uint = uint(f)
	fmt.Println(x, y, z)
}

func add(x, y int) int {
	return x + y
}

func swap(x, y string) (string, string) {
	return y, x
}

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}
