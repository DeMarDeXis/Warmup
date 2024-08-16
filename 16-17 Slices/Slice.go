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
	showAllElements(1, 2, 3)
	showAllElements(1, 2, 3, 4, 5, 6, 7)

	firstSlice := []int{5, 6, 7, 8}
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

func passToFunction() {
	initialSlice := []int{1, 2}
	fmt.Printf("Type: %T Value: %v\n", initialSlice, initialSlice)
	fmt.Printf("Length: %d Capacity: %d\n\n", len(initialSlice), cap(initialSlice))

	changeValue(initialSlice)

	fmt.Printf("Type: %T Value: %v\n", initialSlice, initialSlice)
	fmt.Printf("Length: %d Capacity: %d\n\n", len(initialSlice), cap(initialSlice))

	newSlice := append(initialSlice, 3)
	fmt.Printf("Type: %T Value: %v\n", newSlice, newSlice)
	fmt.Printf("Length: %d Capacity: %d\n\n", len(newSlice), cap(newSlice))

	newSlice2 := appendValue(newSlice)
	fmt.Printf("Type: %T Value: %v\n", newSlice, newSlice)
	fmt.Printf("Length: %d Capacity: %d\n\n", len(newSlice), cap(newSlice))

	fmt.Printf("Type: %T Value: %v\n", newSlice2, newSlice2)
	fmt.Printf("Length: %d Capacity: %d\n\n", len(newSlice2), cap(newSlice2))
}

func changeValue(slice []int) {
	slice[1] = 15
}

func appendValue(slice []int) []int {
	slice = append(slice, 4)
	fmt.Printf("Type: %T Value: %v\n", slice, slice)
	fmt.Printf("Length: %d Capacity: %d\n\n", len(slice), cap(slice))

	return slice
}

func sliceWithNew() {
	slicePointer := new([]int)

	fmt.Printf("Type: %T Value: %v\n", slicePointer, slicePointer)
	fmt.Printf("Length: %d Capacity: %d\n\n", len(*slicePointer), cap(*slicePointer))

	newSlice2 := append(*slicePointer, 1)
	fmt.Printf("Type: %T Value: %#v\n", newSlice2, newSlice2)
	fmt.Printf("Length %T Capacity: %d\n\n", len(newSlice2), cap(newSlice2))
}

func getSlice() {
	intArr := [...]int{1, 2, 3, 4, 5}
	fmt.Printf("Type: %T Value: %#v\n\n", intArr, intArr)

	intSlice := intArr[1:3]
	fmt.Printf("Type: %T Value: %#v\n", intSlice, intSlice)
	fmt.Printf("Length %T Capacity: %d\n\n", len(intSlice), cap(intSlice))

	fullSlice := intArr[:] // intArr[0:5]
	fmt.Printf("Type: %T Value: %#v\n", fullSlice, fullSlice)
	fmt.Printf("Length %T Capacity: %d\n\n", len(fullSlice), cap(fullSlice))

	sliceFromSlice := fullSlice[:3]
	fmt.Printf("Type: %T Value: %#v\n", sliceFromSlice, sliceFromSlice)
	fmt.Printf("Length %T Capacity: %d\n\n", len(sliceFromSlice), cap(sliceFromSlice))

	intArr[2] = 500
	fmt.Printf("Type: %T Value: %#v\n", intArr, intArr)
	fmt.Printf("Type: %T Value: %#v\n", intSlice, intSlice)
	fmt.Printf("Type: %T Value: %#v\n", fullSlice, fullSlice)
	fmt.Printf("Type: %T Value: %#v\n", sliceFromSlice, sliceFromSlice) //все слайсы ссылаются на этот массив
	//6:06
}
