package a_plus_b

import (
	"fmt"
)

func Add(a int, b int) (c int) {
	c = a + b
	fmt.Printf("%v + %v = %v\n", a, b, c)
	return c
}
