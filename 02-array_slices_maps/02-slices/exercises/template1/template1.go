// All material is licensed under the GNU Free Documentation License
// https://github.com/ArdanStudios/gotraining/blob/master/LICENSE

// http://play.golang.org/p/mPKmyGNR2L

// Declare a nil slice of integers. Declare a nil slice of integers. Create a
// loop that appends 10 values to the slice. Iterate over the slice and display
// each value. Iterate over the slice and display each value.
//
// Declare a slice of five strings and initialize the slice with string literal
// values. Display all the elements. Take a slice of index one and two
// and display the index position and value of each element in the new slice.
package main

import "fmt"

// Add imports.

// main is the entry point for the application.
func main() {
	// Declare a nil slice of integers.
	var myslice []int

	// Appends numbers to the slice.
	for i := 0; i < 10; i++ {
		myslice = append(myslice, i*5)
	}

	// Display each value in the slice.
	for i, v := range myslice {
		fmt.Println(i, "\t", v)
	}

	// Declare a slice of strings and populate the slice with names.
	/*	strslice := make([]string, 5)
		for i := range strslice {
			strslice[i] = "hello" + strconv.Itoa(i)
		}
	*/
	//	strslice := []string{"one", "two", "three", "four", "five"}

	// Display each index position and slice value.
	for i, v := range strslice {
		fmt.Println(i, "\t", v)
	}

	// // Take a slice of index 1 and 2 of the slice of strings.
	smallslice := strslice[1:3]

	// Display each index position and slice values for the new slice.
	for i, v := range smallslice {
		fmt.Println(i, "\t", v)
	}

}
