package main

import "fmt"

func main() {
	// Arithmetic operators
	fmt.Println("Arithmetic operators:")
	fmt.Println("Add: 10 + 3 =", 10+3)
	fmt.Println("Sub: 10 - 3 =", 10-3)
	fmt.Println("Mul: 10 * 3 =", 10*3)
	fmt.Println("Div: 10 / 3 =", 10/3)
	fmt.Println("Mod: 10 % 3 =", 10%3)
	fmt.Println()

	// Comparison operators
	// ==    equal
	// !=    not equal
	// <     less
	// <=    less or equal
	// >     greater
	// >=    greater or equal
	fmt.Println("Comparison operators:")
	fmt.Println("Equal: 10 == 3", 10 == 3)
	fmt.Println("Not equal: 10 != 3", 10 != 3)
	fmt.Println("Less: 10 < 3", 10 < 3)
	fmt.Println("Less or equal: 10 <= 3", 10 <= 3)
	fmt.Println("Greater: 10 > 3", 10 > 3)
	fmt.Println("Greater or equal: 10 >= 3", 10 >= 3)
	fmt.Println()

	// Logical operators
	// &&    conditional AND
	// ||    conditional OR
	// !     NOT
	fmt.Println("Logical operators:")
	fmt.Println("AND: 10 + 3 == 13 && 3 - 1 == 4 ->", 10+3 == 13 && 3-1 == 4)
	fmt.Println("OR: 10 + 3 == 10 || 3 - 1 == 2 ->", 10+3 == 13 || 3-1 == 4)
	fmt.Println("OR: 10 + 3 != 14 ->", 10+3 != 14)
	fmt.Println()

	// Asignation operators
	fmt.Println("Asigation Operators:")
	var my_number int = 42
	fmt.Println("Declaration and asignation: var my_number int = 42 -> ", my_number)

	my_number2 := 42
	fmt.Println("Declaration: my_number2 := 42 -> ", my_number2)
	fmt.Println()

	// Binary operators
	fmt.Println("Binary operators:")
	a := 10 // 1010
	b := 3  // 0011

	fmt.Println("AND: 10 & 3 =", a&b)           // 0010 = 2
	fmt.Println("OR: 10 | 3 =", a|b)            // 1011 = 11
	fmt.Println("XOR: 10 ^ 3 =", a^b)           // 1001 = 9
	fmt.Println("AND NOT: 10 &^ 3 =", a&^b)     // 0100 = 8
	fmt.Println("NOT: ^10 =", ^a)               // 0101 = -11
	fmt.Println("Right Shift: 10 >> 2 =", a>>2) // 0010 = 2
	fmt.Println("Left Shift: 10 << 2 =", a<<2)  // 101000 = 40
	fmt.Println()

	// Conditional
	var my_string = "kazte"

	if my_string == "kazte" {
		fmt.Println("my_string is kazte")
	} else if my_string == "pepe" {
		fmt.Println("my_string is pepe")
	} else {
		fmt.Println("my_string is not kazte neither pepe")
	}
	fmt.Println()

	// simple statement before
	// executes the evaluation
	// if x := f(); x < y{
	// 	return x
	// }else{
	// 	return y
	// }

	var my_int = 5

	switch my_int {
	default:
		fmt.Println("Default, my_int not in [1,2,3,4,6,7,8]")
	case 1, 2, 3, 4:
		fmt.Println("my_int is in [1,2,3,4]")
	case 6, 7:
		fmt.Println("my_int is in [6,7]")
	}

	switch {
	case my_int > 5:
		fmt.Println("my_int is greater than 5")
	case my_int < 5:
		fmt.Println("my_int is lesser than 5")
	case my_int == 5:
		fmt.Println("my_int is 5")
	}
	fmt.Println()

	// switch x := f(); { // missing switch expression means "true"
	// case x < 0:
	// 	return -x
	// default:
	// 	return x
	// }

	// Loop statements

	// Single condition for

	var c = 0
	fmt.Println("Single condition")
	for c < 10 {
		fmt.Println(c)
		c++
	}

	// normal for
	fmt.Println("Normal For")
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	var numbers []int = []int{1, 2, 3, 4, 5}

	for i, n := range numbers {
		fmt.Println("n:", n, "\ti:", i)
	}

	// Manejo de excepciones

	// err, num := (10 / 0)

	// if err {
	// 	fmt.Println("Error")
	// } else {
	// 	fmt.Println(num)
	// }

	

}
