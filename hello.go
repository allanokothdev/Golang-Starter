package hello
// Declare main package, packages group functions

// import fmt library containing functions for formatting text, including printing to the console
import "fmt"

// Implement main function. A main fxn executes by default when you run the main package
func hello(){
	fmt.Println("Hello World")
}

// CLI command "go run ." runs the application