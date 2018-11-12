package main

import (
	"fmt"
	"reflect"
)

func main() {
	first := 1.5 + 2.78i
	second := complex64(1.3 + 17i)
	third := complex(15, 1.7)

	fmt.Printf("first (%s) - %v\n", reflect.TypeOf(first), first)
	fmt.Printf("second (%s) - %v\n", reflect.TypeOf(second), second)
	fmt.Printf("third (%s) - %v\n", reflect.TypeOf(third), third)
}