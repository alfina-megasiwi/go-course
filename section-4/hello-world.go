package main

import "fmt"

func HelloWorld() {

	n, _ := fmt.Println("Hello everyone!", 42, true)

	// n, err := fmt.Println("Hello everyone!", 42, true)
	// fmt.Println(err)

	fmt.Println(n)

	foo()
	fmt.Println("something more")

	// iterative
	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}

	bar()
}

func foo() {
	fmt.Println("I'm in foo")
}

func bar() {
	fmt.Println("and then we exited")
}
