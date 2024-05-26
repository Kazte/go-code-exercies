package main

import "fmt"

func main() {
	// https://golang.org/

	// One line comment

	/*
	 * Multiline comment
	 */

	// Variable
	var name string = "Golang"
	const surname string = "Rocks"

	// Data types
	// string, number, booleans
	var i int8 = 1
	var f float32 = 3.66
	var s string = "Minalomeca"
	var b bool = false

	fmt.Println("Hello from " + name)
	fmt.Println("Variables:")
	fmt.Println("integer:", i)
	fmt.Println("float:", f)
	fmt.Println("string:", s)
	fmt.Println("boolean:", b)

}
