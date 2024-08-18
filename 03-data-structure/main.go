/*

Built-In data structures

- Arrays: Fixed-size
- Slices: Dynamic arrays.
- Maps  : Key-Value stores.

*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	// fmt.Println("-----arrays-----")
	// arrays()
	// fmt.Println("\n-----slices-----")
	// slices()
	// fmt.Println("\n------maps------")
	// maps()

	myAgenda()
}

func arrays() {
	// Arrays
	// var array_name = [length]datatype{values} -> length defined
	// var array_name = [...]datatype{values}		 -> length inferred

	array01 := [3]int{1, 5, 7}
	fmt.Printf("array01: %v\n", array01)

	array02 := [...]string{"Minalo", "Meca"}
	fmt.Printf("array02: %v\n", array02)

	// change an element
	array01[0] = 12
	fmt.Printf("array01: %v\n", array01)

	// array initialization
	arr01 := [5]int{}              // not initializated
	arr02 := [5]int{1, 2}          // parcially initializated
	arr03 := [5]int{1, 2, 3, 4, 5} // fully initializated

	fmt.Printf("arr01: %v\n", arr01)
	fmt.Printf("arr02: %v\n", arr02)
	fmt.Printf("arr03: %v\n", arr03)

	// specific initialization
	arr01 = [5]int{1: 10, 4: 15}
	fmt.Printf("arr01: %v\n", arr01)

	// length of an array
	fmt.Printf("len(arr01): %v\n", len(arr01))

	// remove element
	arr03[4] = 0

	fmt.Printf("arr03: %v\n", arr03)

	// sorting
	arr03 = [5]int{5, 2364, 71, 564, 147}
	fmt.Println("before sort:", arr03)
	sort.Ints(arr03[:])
	fmt.Println("after sort:", arr03)

	// matrix
	matrix := [2][3]int{{1, 2, 3}, {6, 5, 4}}
	fmt.Printf("matrix: %v\n", matrix)
}

func slices() {
	// slice_name := []datatype{values}

	slice01 := []int{3, 4, 5}
	fmt.Println("slice01", slice01)
	fmt.Println("len", len(slice01)) // length of the slice
	fmt.Println("cap", cap(slice01)) // capacity of the slice

	// slice from array
	arr01 := [5]int{4, 5, 6, 8, 2}
	slice02 := arr01[1:3]
	fmt.Println("slice02", slice02)

	// create with make
	slice03 := make([]int, 3, 10)
	fmt.Println("slice03", slice03)
	fmt.Println("len", len(slice03))
	fmt.Println("cap", cap(slice03))

	// data manipulation
	// append
	slice04 := []int{1, 2, 3, 4, 5}
	slice04 = append(slice04, 6)
	fmt.Println("slice04", slice04)
	fmt.Println("len", len(slice04))
	fmt.Println("cap", cap(slice04))

	// modification
	slice04[1] = 12
	fmt.Println("slice04 mod", slice04)
	fmt.Println("len", len(slice04))
	fmt.Println("cap", cap(slice04))

	// delete
	slice04 = append(slice04[:2], slice04[2+1:]...)
	fmt.Println("slice04 del", slice04)
	fmt.Println("len", len(slice04))
	fmt.Println("cap", cap(slice04))
}

func maps() {
	// the map key can be of any data type for which the equality operator (==) is defined. These include:
	// 	Booleans
	// Numbers
	// Strings
	// Arrays
	// Pointers
	// Structs
	// Interfaces (as long as the dynamic type supports equality)

	// b := map[KeyType]ValueType{key1:value1, key2:value2,...}
	map01 := map[string]int{"kazte": 64, "pepe": 2}
	fmt.Println("map01", map01)

	// using make
	// b := make(map[KeyType]ValueType)
	map02 := make(map[string]string)
	map02["elpepe"] = "minalo"
	map02["kazte"] = "meca"
	fmt.Println("map02", map02)

	// access
	fmt.Println(map02["kazte"])

	// update and add
	map02["elpepe"] = "pepe"
	map02["bucho"] = "frienddd!"

	fmt.Println(map02)

	// remove
	delete(map02, "kazte")
	fmt.Println(map02)

	// specific key
	// val, ok :=map_name[key]
	val, ok := map02["kum"]
	fmt.Println(val, ok)
	val, ok = map02["bucho"]
	fmt.Println(val, ok)

	// maps are references
	a := map[string]string{"brand": "Ford", "year": "1998"}
	b := a

	fmt.Println("a", a)
	fmt.Println("b", b)

	b["year"] = "1997"

	fmt.Println("after change b", a)
	fmt.Println("after change b", b)

	// iterate maps
	// maps are unordered data structures
	fmt.Println("unordered")
	map03 := map[string]int{"one": 1, "two": 2, "three": 3, "four": 4}
	for key, value := range map03 {
		fmt.Println(key, value)
	}

	// iterate in a specific order
	fmt.Println("ordered")
	map04 := map[string]int{"one": 1, "two": 2, "three": 3, "four": 4}

	var c []string
	c = append(c, "one", "two", "three", "four")

	for _, element := range c {
		fmt.Println(element, map04[element])
	}
}

func myAgenda() {
	showPrompt()
	scanner := bufio.NewScanner(os.Stdin)
	contacts := map[string]int{}

	for scanner.Scan() {
		text := scanner.Text()

		if text == "5" {
			break
		}

		if text == "" {
			fmt.Println("Select an option")
		}

		optionSelected(text, contacts)
		showPrompt()

	}

	if err := scanner.Err(); err != nil {
		fmt.Println("error", err)
	}
}

func showPrompt() {
	fmt.Println("1. Search")
	fmt.Println("2. Add")
	fmt.Println("3. Update")
	fmt.Println("4. Delete")
	fmt.Println("5. Quit")
}

func optionSelected(text string, contacts map[string]int) {
	if text == "1" {
		fmt.Println("Who are you looking for?")
		fmt.Println("Type -1 to return to menu")
		scanner := bufio.NewScanner(os.Stdin)

		for scanner.Scan() {
			text := scanner.Text()

			if text == "-1" {
				break
			}

			val, ok := contacts[text]

			if ok {
				fmt.Printf("%v: %v", contacts[text], val)
			} else {
				fmt.Println("Contact", text, "not found:")
				fmt.Println("Type -1 to return to menu")
				fmt.Println("Try again:")
			}
		}

	} else if text == "2" {
		scanner := bufio.NewScanner(os.Stdin)
		name := ""
		number := -1

		fmt.Println("Please write the name of your contact and then the number (must be lesse than 11 digits and only numbers)")

		for scanner.Scan() {
			text := scanner.Text()

			if text == "" {
				break
			}

			name = text
		}

		if name != "" {
			for scanner.Scan() {
				text := scanner.Text()

				if text == "" {
					break
				}

				i, err := strconv.Atoi(text)

				if err != nil || len(text) > 11 {
					fmt.Println("Wrong number...")
				} else {
					number = i
					break
				}
			}
		}

		if number != -1 {
			contacts[name] = number
			fmt.Println("Added", name, "with number", number)
		}

	} else if text == "3" {

	} else if text == "4" {

	}
}
