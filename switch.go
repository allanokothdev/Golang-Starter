package switch

import (
	"fmt"
	"time"
)

func switch(){

	i := 2
	fmt.Println("Write ",i," as ")

	//	Basic switch statement
	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	}

	switch time.Now().Weekday() {

	// Use Commas to separate multiple expressions
	case time.Saturday, time.Sunday:
		fmt.Println("It's the weekend")
	default:
		fmt.Println("It's a weekday")
	}

	t:= time.Now()
	//Switch without login is another way of expressing if-else | also the case expressions can be non-constants
	switch {
	case t.Hour()< 12:
		fmt.Println("It's before noon")
	default:
		fmt.Println("it's after noon")
	}

	//	A Type switch compares types instead of values
	whatAmI := func(i interface{}){
		switch t:= i.(type) {
		case bool:
			fmt.Println("I'm a bool")
		case int:
			fmt.Println("I'm an int")
		case string:
			fmt.Println("I'm an String")
		default:
			fmt.Println("Don't know type %T\n",t)
		}
	}
	whatAmI(true)
	whatAmI(1)
	whatAmI("hey")
}