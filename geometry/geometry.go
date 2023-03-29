package geometry

import "fmt"

func init() {
	fmt.Println("In geometry init")
}

func Area(side int) int {
	return side * side
}
