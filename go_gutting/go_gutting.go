package main

import (
	"errors"
	"fmt"
	"math"
)

type Person struct {
	name string
	age  int
}

func gutting() {
	// CMD print
	fmt.Println("Hello, playground")

	// Variable Declaration!
	x := 5        // Same as below
	var y int = 7 // Same as above
	var sum int = x + y
	fmt.Println(sum)

	// Conditionals!
	if x > 3 {
		fmt.Println("More than 3")
	} else if x == 5 {
		fmt.Println("X is 5")
	} else {
		fmt.Println("idk what x is")
	}

	// Fixed Length Arrays!
	var a [5]int               // Initialized with zero values [0 0 0 0 0]
	a[2] = 7                   // Index setting
	b := [5]int{5, 4, 3, 2, 1} // Shorthanded with values! [5 4 3 2 1]
	fmt.Println(a)
	fmt.Println(b)

	// Slices (Abstraction over arrays, providing dynamically lenghted arrays)
	c := []int{9, 8, 7, 6, 5}
	c = append(c, 2)
	fmt.Println(c)

	// Key-Value Pairs (Maps or Dictionaries)
	students := make(map[int]string)
	students[289388] = "Joshua"
	students[128137] = "Emily"
	students[830180] = "Terren"

	delete(students, 830180)

	fmt.Println(students)
	fmt.Println(students[128137])

	// Loops
	// For loop
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	// While loop
	l := 15
	for l > 10 {
		fmt.Println(l)
		l--
	}

	// Loop through arrays
	arr := []string{"a", "b", "c"}
	for index, value := range arr {
		fmt.Println("index:", index, "value:", value)
	}

	// Loop through maps
	vertices := make(map[string]int)
	vertices["square"] = 4
	vertices["triangle"] = 3

	for key, value := range vertices {
		fmt.Println("key:", key, "value:", value)
	}

	// Functions!
	result := summation(1, 2)
	fmt.Println(result)

	// Multi-Return Functions!
	sqrtResult, err := sqrt(16)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(sqrtResult)
	}

	// Structs
	p := Person{name: "Joshua", age: 90000000}
	fmt.Println(p)
	fmt.Println(p.age)

	// Pointers
	o := 7
	fmt.Println(o)  // Value
	fmt.Println(&o) // Memory address, a pointer to [o]
	inc(o)          // Does nothing, o is passed by value. Original [o] is not mutated.
	fmt.Println(o)
	incPointer(&o) // Mutates the value at the memory address
	fmt.Println(o)
}

// x is passed by value, no mutation occurs to original
func inc(x int) {
	x++
}

func incPointer(x *int) {
	// This is wrong. Apparently this increments the memory address itself. Which is
	// kinda useless, as we want to mutate the VALUE at the memory address.
	// x++

	// Dereferencing. Increments the actual value.
	*x++
}

func summation(x int, y int) int {
	return x + y
}

func sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, errors.New("Undefined for negative numbers")
	}

	return math.Sqrt(x), nil
}
