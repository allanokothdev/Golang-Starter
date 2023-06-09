package constantFxn

import (
	"fmt"
	"math"
)

// CONST declares a constant value
const s string = "constant"

func constantFxn() {
	fmt.Println(s)

	const n = 50000000

	const d = 3e20 / n
	fmt.Println(d)

	fmt.Println(int64(d))

	fmt.Println(math.Sin(n))
}