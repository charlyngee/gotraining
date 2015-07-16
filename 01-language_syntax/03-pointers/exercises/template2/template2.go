// All material is licensed under the GNU Free Documentation License
// https://github.com/ArdanStudios/gotraining/blob/master/LICENSE

// http://play.golang.org/p/b6-FNFOToO

// Declare a struct type and create a value of this type. Declare a function
// that can change the value of some field in this struct type. Display the
// value before and after the call to your function.
package main

// Add imports.
import "fmt"

// Declare a type named user.
type user struct {
	field1 int
	field2 string
}

// Create a function that changes the value of one of the user fields.
func funcName(uptr *user, newfield1 int) {
	// Use the pointer to change the value that the
	// pointer points to.
	uptr.field1 = newfield1
}

// main is the entry point for the application.
func main() {
	// Create a variable of type user and initialize each field.
	u1 := user{
		field1: 1,
		field2: "hello",
	}

	// Display the value of the variable.
	fmt.Printf("%+v", u1)

	// Share the variable with the function you declared above.
	funcName(&u1, 12345)

	// Display the value of the variable.
	fmt.Println("field1 = ", u1.field1)
}
