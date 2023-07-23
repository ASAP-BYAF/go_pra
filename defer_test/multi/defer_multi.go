package multi

import "fmt"

func DeferMultiTest() {
    fmt.Println("counting")

	// LIFO
	for i:=0; i<10; i++{
		defer fmt.Println(i)
	}

	fmt.Println("done")
}