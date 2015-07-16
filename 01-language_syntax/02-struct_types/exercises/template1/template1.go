// All material is licensed under the GNU Free Documentation License
// https://github.com/ArdanStudios/gotraining/blob/master/LICENSE

// http://play.golang.org/p/ItPe2EEy9X

// Declare a struct type to maintain information about a user (name, email and age).
// Create a value of this type, initalize with values and display each field.
//
// Declare and initialize an anonymous struct type with the same three fields. Display the value.
package main

// Add imports.
import "fmt"

// Add user type and provide comment.
type user struct {
	name	string
	email	string
	age		int
}

// main is the entry point for the application.
func main() {
	// Declare variable of type user and init using a struct literal.
	u1 := user{
		name: "Gee",
		email: "gee@test.com",
		age: 12345,
	}
	// Display the field values.
	fmt.Printf("%+v\n", u1)

	// Declare a variable using an anonymous struct.
	u2 := struct {
		name	string
		email	string
		age		int
	}{
		name: "Charlyn",
		email: "test@test.com",
		age: 32,
	}

	// Display the field values.
	fmt.Printf("%+v\n", u2)
}
