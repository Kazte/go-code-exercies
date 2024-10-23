package main

import (
	"fmt"
	"strconv"
)

func main() {
	noParamsFunction()

	paramsFunction("Kazte")

	twoParamsFunction(10, 3)

	pow := returnFunction(2, 5)
	fmt.Println("2^5 =", pow)

	numbers := extraDifficulty("minalo", "meca")
	fmt.Println("numbers:", numbers)

}

func noParamsFunction() {
	a := 3
	b := 4
	sum := a + b

	fmt.Println("Sum 3 + 4 is", sum)
}

func paramsFunction(name string) {
	fmt.Println("Hello", name)
}

func twoParamsFunction(a int, b int) {
	sub := a - b
	fmt.Println("Sub", a, "-", b, "is", sub)
}

func returnFunction(a int, b int) int {
	pow := 1
	for i := 0; i < b; i++ {
		pow *= a
	}

	return pow
}

func extraDifficulty(a string, b string) int {
	count := 0

	for i := 0; i < 100; i++ {

		print := ""
		number := true

		if i%3 == 0 {
			print += a
			number = false
		}

		if i%5 == 0 {
			print += b
			number = false
		}

		if number {
			print = strconv.Itoa(i)
			count++
		}

		fmt.Println(print)
	}

	return count
}
