package main

import (
	"fmt"

	"golang.org/x/example/stringutil"
	"rsc.io/quote"
)

func main() {
	fmt.Println("Hello, World!")
	fmt.Println("quote= ", quote.Go())
	fmt.Println("stringutil= ", stringutil.Reverse("Hello"))
}
