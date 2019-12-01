//
//
// gobyexample.com
// pointers.go
//
package main

import "fmt"

func zeroval(ival int) {
	ival = 0
}

func zeroptr(ipter *int) {
	*ipter = 0
}

func main() {

	i := 1
	fmt.Println("Initial value: ", i)

	zeroval(i)
	fmt.Println("Zero value: ", i)

	zeroptr(&i)
	fmt.Println("Zero pointer: ", i)

	fmt.Println("Pointer: ", &i)
}
