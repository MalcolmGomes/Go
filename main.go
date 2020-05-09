package main

import (
	"errors"
	"fmt"
)

func main() {

	// helloWorld()

	// variables() // Variables

	// conditionals() // Conditionals

	// arrays() // Arrays

	// maps() // Maps

	// loop() // For

	// Functions
	// result := sum(2, 1)
	// fmt.Println(result)

	// Structs
	// p := person{name: "Malcolm", age: 20}
	// fmt.Println(p)

	// Pointers:
	// i := 7
	// inc(&i)
	// fmt.Println(i)
}

func inc(x *int) {
	*x++
}

type person struct {
	name string
	age  int
}

func sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, errors.New("undefined for negative numbers")
	}
	return 0, nil
}

func sum(x int, y int) int {
	return x + y
}

func loop() {
	// For Loop
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	// "While" Loop
	// i := 0
	// for i < 10 {
	// 	fmt.Println(i)
	// 	i++
	// }

	// Looping over array via range
	// arr := []string{"a", "b", "c"}

	// for index, value := range arr {
	// 	fmt.Println("index", index, "value", value)
	// }

	// Looping over map
	// m := make(map[string]string)
	// m["a"] = "alpha"
	// m["b"] = "beta"

	// for key, value := range m {
	// 	fmt.Println("key:", key, "value:", value)
	// }
}

func maps() {
	vertices := make(map[string]int)
	vertices["triangle"] = 2
	vertices["square"] = 3
	vertices["dodecagon"] = 12

	delete(vertices, "square")
	fmt.Println(vertices)
}

func arrays() {
	// var a [5]int // Array
	a := []int{5, 4, 3, 2, 1} // Slice
	a[2] = 7
	a = append(a, 13)
	fmt.Println(a)
}

func conditionals() {
	x := 4
	if x > 6 {
		fmt.Println("More than 6.")
	} else if x < 6 {
		fmt.Println("Less than 6.")
	} else {
		fmt.Println("Exactly 6.")
	}
}

func helloWorld() {
	fmt.Println("Hello, World!")
}

func variables() {
	var x int = 5
	y := 7
	sum := x + y
	fmt.Println(sum)
}
