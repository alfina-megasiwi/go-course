package main

import "fmt"

var yy = 42

// declared that variable with identifier "z" is of type string
// go is a static programming language
// a variable is declared to hold a value of a certain type
var zz = "Shaken, not stirred"
var aa string = `James said, 
"Shaken,

not stirred"`

func ExploringType() {
	fmt.Println(yy)
	fmt.Printf("%T\n", yy)
	fmt.Println(zz)
	fmt.Printf("%T\n", zz)
	fmt.Println(aa)
	fmt.Printf("%T\n", aa)
	// zz = 43 we couldn't assign 43

}
