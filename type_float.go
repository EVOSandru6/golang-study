package main

import (
	"fmt"
	"math"
	"reflect"
)

func main() {
	var defaultFloat float32;
	floatVar := 5.5

	fmt.Println("defaultFloat = ", defaultFloat)
	fmt.Printf("floatVar (%s) = %v\n", reflect.TypeOf(floatVar), floatVar) // 5.5

	fmt.Println("\n------------------------------------------------------------------\n")

	fmt.Println("Max float32 = ", math.MaxFloat32)
	fmt.Println("Min float32 = ", math.SmallestNonzeroFloat32, "\n")

	fmt.Println("Max float64 = ", math.MaxFloat64)
	fmt.Println("Min float64 = ", math.SmallestNonzeroFloat64, "\n")

}
