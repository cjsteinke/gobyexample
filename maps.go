package main

import "fmt"

func main() {

	myMap := make(map[string]int)

	myMap["k1"] = 7
	myMap["k2"] = 13

	fmt.Println("map:", myMap)

	v1 := myMap["k1"]
	fmt.Println("v1: ", v1)

	fmt.Println("len:", len(myMap))

	delete(myMap, "k2")
	fmt.Println("map:", myMap)

	_, prs := myMap["k2"]
	fmt.Println("prs:", prs)

	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map:", n)

}
