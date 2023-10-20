package main

import (
	"fmt"
)

func main() {
	variadicFunctions()
	//convertToArrayPointer()
	//passToFunction()
	//sliceWithNew()
	//getSlice()
	//copySlice()
	//deleteElement()
}

func variadicFunctions() {
	showAllElements(1,2,3)
	showAllElements(1,2,3,4,5,6,7)

	firstSlice := []int{5,6,7,8}
	secondSlice := []int{9, 3, 2, 1}

	showAllElements(firstSlice...) // (5, 6, 7, 8)

	newSlice := append(firstSlice, secondSlice...) //(firstSlice, 9, 3, 2, 1)
	fmt.Printf("Type: %T Value: %#v\n", newSlice, newSlice)
}

func showAllElements(values ...int) {
	for _, val := range values {
		fmt.Println("Values:", val)
	}
	fmt.Println()
}

func convertToArrayPointer() {
	SliceInit := []int{1, 2}
	fmt.Printf("Type: %T Value: %#v\n", SliceInit, SliceInit)
	fmt.Printf("Length: %d Capacity: %d\n\n", len(SliceInit), cap(SliceInit))

	intArray := (*[2]int)(SliceInit)
	fmt.Printf("Type: %T Value: %#v\n", intArray, intArray)
	fmt.Printf("Length: %d Capacity: %d\n\n", len(intArray), cap(intArray))
}
//6 18