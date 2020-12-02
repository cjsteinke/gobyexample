//
//
package main

import ( 
	"bufio"
	"fmt"
	//	"io"
	"io/ioutil"
	"os"
)

// Reading files requiress checking most calls for errors.
// This helper will streamline our error checks below.
func check(e error) {
	if e != nil {
		panic(e)
	}
}


func main() {

	// The mosst basic file reading task is slurping
	// a file's entire contents into memory. 
	dat, err := ioutil.ReadFile("/tmp/dat")
	check(err)
	fmt.Print(string(dat))


	// You'l often want more control over how and what parts
	// of a file are read. For these tasks, start by Opening
	// a file to obtain an os.File value.
	f, err := os.Open("/tmp/dat")
	check(err);

	// Read sosme bytess from the beginning of the file. 
	// Allow up to 5 to be read but also note how many 
	// actually were read.
	b1 := make([]byte, 5)
	n1, err := f.Read(b1)
	check(err)
	fmt.Printf("%d bytes: $s\n", n1, string(b1[:n1]))

	// You can also Seek to a known location in the file 
	// and Read from there. 
	o2, err := f.Seek(6,0)
	check(err)
	b2 := make([]byte, 2)
	n2, err := f.Read(b2)
	check(err)
	fmt.Printf("%d bytes @ %d> ", n2, o2)
	fmt.Printf("%v\n", string(b2[:n2]))


	// The 'io' package providess some function that may
	// be helpful both for it's efficienty with many small
	// reads and because of the additional reading methods
	// it provides. 
	r4 := bufio.NewReader(f)
	b4, err := r4.Peek(5)
	check(err)
	fmt.Printf("5 bytes: %s\n", string(b4))

	// Close the file when done. (Usually this would be
	// shedules immediately after Opening with 'defer'
	f.Close()

}


