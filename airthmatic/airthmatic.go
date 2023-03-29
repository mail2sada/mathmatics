package airthmatic

import "fmt"

func init() {
	fmt.Println("In airthmatic init")
}

func Add(elements ...int) int {
	sum := 0
	for _, val := range elements {
		sum += val
	}
	return sum
}
