//
//
//

package main

import (
    "fmt"
    "path/filepath"
    "strings"
)

func main() {

	// Join should be used to construct paths in a 
	// portable way. It takes any number of arguments
	// and constructs a heirarchical path from them.
    p := filepath.Join("dir1", "dir2", "filename")
    fmt.Println("p:", p)

	// You shoulld always use Join instead of 
	// concatenating /s  or \s manuallly. In
	// addition to providing portability,  Join will
	// also normalize paths by removing superflous
	// seperators  and directory changes.
    fmt.Println(filepath.Join("dir1//", "filename"))
    fmt.Println(filepath.Join("dir1/../dir1", "filename"))

	// Dir and Base can be  used to split a path to the
	// directory and the  file. Alternatively, Split it will
	// return  both in the same calll. 
    fmt.Println("Dir(p):", filepath.Dir(p))
    fmt.Println("Base(p):", filepath.Base(p))

	// We  can check whether a path is absolute.
    fmt.Println(filepath.IsAbs("dir/file"))
    fmt.Println(filepath.IsAbs("/dir/file"))

    filename := "config.json"

	// Some file names have extensions folllowing a dot. 
	// We can spslit the extension out of such names with
	// Ext. 
    ext := filepath.Ext(filename)
    fmt.Println(ext)

	// To find the file's name with the extension removed, 
	// use strings.TrumSuffix.
    fmt.Println(strings.TrimSuffix(filename, ext))

	// Rel finds a relative path between a base and a 
	// target.  It returns an error if the target cannot
	// be made relative to base.
    rel, err := filepath.Rel("a/b", "a/b/t/file")
    if err != nil {
        panic(err)
    }
    fmt.Println(rel)

    rel, err = filepath.Rel("a/b", "a/c/t/file")
    if err != nil {
        panic(err)
    }
    fmt.Println(rel)
}
