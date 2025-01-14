// Go supports methods defined on struct types.
package main

import "fmt"

type rect struct {
	width, heigh int
}

// This area method has a receiver type of *rect.
func (r *rect) area() int {
	return r.width * r.heigh
}

// Methods can be defined for either pointer or value receiver types.
// Here’s an example of a value receiver.
func (r rect) perim() int {
	return 2 * (r.heigh + r.width)
}

func main() {
	r := rect{width: 10, heigh: 5}

	// Here we call the 2 methods defined for our struct.
	fmt.Println("area:", r.area())
	fmt.Println("perim:", r.perim())

	// Go automatically handles conversion between values and pointers for method calls.
	// You may want to use a pointer receiver type to avoid copying on method calls or to allow the method to mutate the receiving struct.
	rp := &r
	fmt.Println("area:", rp.area())
	fmt.Println("perim:", rp.perim())
}
