package main

import (
	"fmt"

	"github.com/samir/fuzz/fuzz"
)

func main() {
	fmt.Println("My string compare is cooler and better than yours!")
	result := fuzz.Reverse("abcd")
	fmt.Println(result)
}
