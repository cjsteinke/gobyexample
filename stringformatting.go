//
// stringformatting.go
// Go by Example
//
package main

import (
	"fmt"
	"os"
)

type point struct {
	x, y int
}

func main() {

	// print an instance of our point struct
	p := point{1, 2}
	fmt.Printf("%v\n", p)

	// If the value is a struct, the %+v variant
	// will include the struct’s field names.
	fmt.Printf("%+v\n", p)

	// The %#v variant prints a Go syntax representation of the value,
	// i.e. the source code snippet that would produce that value.
	fmt.Printf("%#v\n", p)

	// Print the type of a value, use %T.
	fmt.Printf("%T\n", p)

	fmt.Printf("%t\n", true)

	fmt.Printf("%d\n", 123)

	fmt.Printf("%b\n", 14)

	fmt.Printf("%c\n", 33)

	fmt.Printf("%x\n", 456)

	fmt.Printf("%f\n", 78.9)

	fmt.Printf("%e\n", 123400000.0)
	fmt.Printf("%E\n", 123400000.0)

	fmt.Printf("%s\n", "\"string\"")

	fmt.Printf("%q\n", "\"string\"")

	fmt.Printf("%x\n", "hex this")

	fmt.Printf("%p\n", &p)

	fmt.Printf("|%6d|%6d|\n", 12, 345)

	fmt.Printf("|%6.2f|%6.2f|\n", 1.2, 3.45)

	fmt.Printf("|%-6.2f|%-6.2f|\n", 1.2, 3.45)

	fmt.Printf("|%6s|%6s|\n", "foo", "b")

	fmt.Printf("|%-6s|%-6s|\n", "foo", "b")

	s := fmt.Sprintf("a %s", "string")
	fmt.Println(s)

	fmt.Fprintf(os.Stderr, "an %s\n", "error")
}
