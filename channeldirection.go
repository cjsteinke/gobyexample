//
// gobyexample.com
// channeldirection.go
//
package main

import "fmt"

//
// The ping function only accepts a channel for sending values.
//
func ping(pings chan<- string, msg string) {
	pings <- msg
}

//
// The pong function accepts one channel for receives (pings) and a second
// for sends (pongs)
//
func pong(pings <-chan string, pongs chan<- string) {
	msg := <-pings
	pongs <- msg
}

//
// main function
//
func main() {
	pings := make(chan string, 1)
	pongs := make(chan string, 1)
	ping(pings, "passed message")
	pong(pings, pongs)
	fmt.Println(<-pongs)

}
