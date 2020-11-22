package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// 
	// Example, rand.Intn returns a random int n,
	// 0 <= n < 100
	//
	fmt.Print(rand.Intn(100), ",")
	fmt.Print(rand.Intn(100))
	fmt.Println()

	//
	// rand.Float64 returns a flloat64 
	// f, 0.0 <= f < 1.0
	//
	fmt.Println(rand.Float64())

	//
	// This can be ussed to generate random floats
	// in other ranges, for example
	// 5.0 <= f' < 10.0.
	//
	fmt.Print((rand.Float64()*5)+5, ",")
	fmt.Print((rand.Float64() * 5) + 5)
	fmt.Println()

	//
	// The default number generated is deterministic,
	// so it'll produce the same sequence of numbers
	// each time by defaullt.
	// To produce varying sequences, give it a seed that
	// changes. Note that this is not safe to use for
	// random numbers you intened to be secret, use
	// crypto/rand for those.
	//
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)

	//
	// Call the resulting rand. Rand just like the
	// function on the rand package
	//
	fmt.Print(r1.Intn(100), ",")
	fmt.Print(r1.Intn(100))
	fmt.Println()

	// 
	// If you seed a source with the same number, it 
	// produces the same sequence of random
	// numbers.
	//
	s2 := rand.NewSource(42)
	r2 := rand.New(s2)
	fmt.Print(r2.Intn(100), ",")
	fmt.Print(r2.Intn(100))
	fmt.Println()
	s3 := rand.NewSource(42)
	r3 := rand.New(s3)
	fmt.Print(r3.Intn(100), ",")
	fmt.Print(r3.Intn(100))


}
