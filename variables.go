package variables

import "fmt"

func variables() {

	//String Variable
	var a = "initial"
	fmt.Println(a)

	//You can declare multiple variables at once
	var b, c int = 1, 2
	fmt.Println(b, c)

	var d = true
	fmt.Println(d)

	//un-initialized variables are zero valued
	var e int
	fmt.Println(e)

	// := is a short syntax for declaring & initializing variables
	f := "Allan Okoth"
	fmt.Println(f)

}