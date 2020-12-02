//
// gobyexample.com
// Temporary  Files and Directories
//  temp.go
//
package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)


func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {

	// The easiesst way to create a temporary file
	// is by  calling ioutil.TempFile. It createss
	// a file and opens it for reading and writing. 
	// We provide "" as the first argument, sso 
	// ioutil.Tempfile will create the file in the 
	// default location for out OS.
	f, err := ioutil. TempFile("", "sample")
	check(err)

	// Display the name of the temporary file.
	// On Unix-based OSes the directory will likely be /tmp.
	// The file name starts with the prefix given as the
	// second argument to ioutil.TempFile and the rest is
	// chosen automatically to ensure that concurrent calls
	// will always create different file names.
	fmt.Println("Temp file  name:", f.Name())

	//  Clean up the file after we’re done. 
	// The OS is likely to clean up temporary files by
	// itself after some time, but it’s good practice to
	// do this explicitly.
	defer  os.Remove(f.Name())

	// We can write sosme data to  the file.  
	_, err =  f.Write([]byte{1,2,3,4})
	check(err)
	
	// If we intend to write many temporary files,
	// we may prefer to create a temporary directory.
	// ioutil.TempDir’s arguments are the same as TempFile’s,
	// but it returns a directory name rather than an open file.
	dname, err := ioutil.TempDir("", "sampledir")
    check(err)
    fmt.Println("Temp dir name:", dname)

	defer os.RemoveAll(dname)

	// Now we can synthesize temporary  file namess by 
	// prefixing them with our temporary directory.
	fname := filepath.Join(dname, "file1")
    err = ioutil.WriteFile(fname, []byte{1, 2}, 0666)
    check(err)
}


