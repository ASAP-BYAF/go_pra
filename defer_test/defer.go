package defer_test

import "fmt"

func DeferTest() {
	defer fmt.Printf(" world\n")

	fmt.Printf("hello")
}