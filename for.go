package for

import "fmt"

func for() {

	// We declare and initialize i with 1
	i := 1

	// Basic for with a single condition
	for i <= 3 {
		fmt.Println(i)
		i=i+1
	}

	// For without a condition will loop repeatedly until you break out of the loop or return from the enclosing function
	for {
		fmt.Println("Loop")
		break
	}

	// classic initial/condition/after for loop
	for j:=7;j<=9;j++ {
		fmt.Println(j)
	}

	for n := 0; n <= 5; n++ {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}
}