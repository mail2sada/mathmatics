package main

import "fmt"
import "github.com/mail2sada/mathmatics/airthmatic"
import "github.com/mail2sada/mathmatics/geometry"

func init() {
	fmt.Println("In main init")
	fmt.Println(airthmatic.Add(10, 20, 20))
	fmt.Println(geometry.Area(10))
}

func main() {
	fmt.Println("In main")
}
