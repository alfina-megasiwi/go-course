package main

import "fmt"

var temp = 42

func FmtPackage() {
	fmt.Println(temp)
	fmt.Printf("%T\n", y)
	fmt.Printf("%b\n", y)
	fmt.Printf("%x\n", y)
	fmt.Printf("%#x\n", y)
	y = 911
	fmt.Printf("%#x\n", y)
	fmt.Printf("%#x\t%b\t%x\n", y, y, y)

	s := fmt.Sprintf("%#x\t%b\t%x", y, y, y)
	fmt.Println(s)
	fmt.Printf("%T\n", s)
	fmt.Printf("%v", y) // the default value

}
