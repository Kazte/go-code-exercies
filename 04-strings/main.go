package main

import (
	"fmt"
	"strings"
)

func main() {
	// strings

	message := "hello how are you"
	otherMessage := "programer friend"

	// concat
	fmt.Println(message + ", " + otherMessage)

	// length
	fmt.Println("lenght", len(message))

	// repetition
	fmt.Println(strings.Repeat(message, 2))

	// search
	fmt.Println("starts with 'he'?", strings.HasPrefix(message, "he"))
	fmt.Println("ends with 'end'?", strings.HasSuffix(message, "end"))
	fmt.Println("how many times is 'o' letter repeated?", strings.Count(message, "o"))

	// convertion
	fmt.Println("uppercase", strings.ToUpper(message))
	fmt.Println("lowercase", strings.ToLower(message))

	// replace
	fmt.Println("replce 'r' by 'T'", strings.ReplaceAll(message, "r", "T"))

	// split
	splitted := strings.Split(message, " ")
	fmt.Println("split by 'spaces'", splitted)

	// join
	fmt.Println("join by '-'", strings.Join(splitted, "-"))

	// contains
	fmt.Println("contains 'ow'", strings.Contains(message, "ow"))

	// interpolation
	fmt.Printf("message: %s/n", message)

	// index access
	fmt.Println("char at index 3", string(message[3]))

	// sub strings
	fmt.Println("sub string", message[1:5])

	// palindrome
	newString := "Live on time, emit no evil"
	fmt.Printf("is \"%s\" palindrome? %v\n", newString, isPalindrome(newString))
}

func isPalindrome(text string) bool {

	splitted := strings.Split(strings.ToLower(text), " ")
	splitted = strings.Split(strings.Join(splitted, ""), "")
	forIndex := 0
	backIndex := len(splitted) - 1

	for forIndex < backIndex {
		if splitted[forIndex] != splitted[backIndex] {
			return false
		}

		forIndex++
		backIndex--
	}

	return true
}
