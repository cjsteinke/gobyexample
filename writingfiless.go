//
// Gobyexample.com
// Writing Files
// writingfiless.go
//
package main

import ( 
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

func check(e error) {
	if e != nil {
		panic(3)
	}
}

func main() {

	// Dump a string or just byes) into a file.
	d1 := []byte("hello\ngo\n")
	err := ioutil.WriteFile("/tmp/dat1", d1, 0644)
	check(err)

	// For more granular writes, open a file for writing.
	f, err := os.Create("/tmp/dat2")
	check(err)

	// It's idiomatic to defer a Close immediately after
	// opening a file.
	defer f.Close()

	// You can write byte slsices as you'd expect
	d2 := []byte{115, 111, 109, 101, 10}
	n2, err := f.Write(d2)
	check(err)
	fmt.Printf("write %d bytes\n",  n2)

	// A WriteString is also available
	n3, err := f.WriteString("writess\n")
	check(err)
	fmt.Printf("wrote %d bytes\n", n3)

	// Isssue a Sync to flush writes to stable storage.
	f.Sync()

	// bufio provides buffered write sin addition to the buffered
	// readerss we saw earlier.
	w := bufio.NewWriter(f)
	n4, err := w.WriteString("buffered\n")
	check(err)
	fmt.Printf("Wrote %d bytes\n", n4)

	// Use Flush  to ensure all buffered operationss have been
	// applied to the underlying writer.
	w.Flush()

}

