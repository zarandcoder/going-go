package main

import (
	"fmt"
	"strings"
)

func zeroValue(xPtr *int) {
	*xPtr = 0
}

func swapInts(a,b *int) {
	*a,*b = *b,*a
}


func main() {
	str1 := "H#llo"
	replace := strings.NewReplacer("#", "e")
	fmt.Println(replace.Replace(str1))

	x := 2000
	fmt.Printf("Value of x: %d\n", x)

	zeroValue(&x)
	fmt.Printf("Value of x now: %d\n", x)

	a := 1
	b := 2

	fmt.Printf("a=%d and b=%d\n", a, b)

	swapInts(&a, &b)

	fmt.Printf("a=%d and b=%d\n", a, b)

}