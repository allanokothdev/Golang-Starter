package values

import "fmt"

// Go has various value types including Strings, Integers, Floats. Booleans

func values() {
	// String Concatenation
	fmt.Println("go" + "lang")

	//Float & Integers addition
	fmt.Println("1+1 =", 1+1)
	fmt.Println("7.0/3.0 =", 7.0/3.0)

	fmt.Println(true && false)
	fmt.Println(true || false)
	fmt.Println(!true)
}