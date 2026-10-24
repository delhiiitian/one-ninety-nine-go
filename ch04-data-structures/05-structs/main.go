// Chapter 4.5: Structs (group related data)
// Run: go run ./ch04-data-structures/05-structs
package main

import "fmt"

// A Student record: like one row in a school register.
type Student struct {
	Name   string
	Roll   int
	Class  string
	Marks  []int
	Active bool
}

type Address struct {
	City, Pin string
}

type Customer struct {
	Name    string
	Address Address // a struct inside a struct
}

func main() {
	// Create with field names (clearest)
	s := Student{Name: "Ananya", Roll: 7, Class: "8B", Marks: []int{88, 92, 79}, Active: true}
	fmt.Println(s.Name, "is in class", s.Class)

	// Change a field
	s.Marks = append(s.Marks, 95)
	fmt.Printf("%+v\n", s)

	// Missing fields get zero values
	newKid := Student{Name: "Yash"}
	fmt.Printf("new: %+v\n", newKid)

	// A slice of structs is a simple table
	register := []Student{
		s,
		{Name: "Rahul", Roll: 8, Class: "8B", Marks: []int{70, 65}},
	}
	for _, st := range register {
		fmt.Printf("%2d %-8s %v\n", st.Roll, st.Name, st.Marks)
	}

	c := Customer{Name: "Leela", Address: Address{City: "Kochi", Pin: "682001"}}
	fmt.Println(c.Name, "lives in", c.Address.City)

	// Anonymous struct: handy for one-off data
	point := struct{ X, Y int }{3, 4}
	fmt.Println("point:", point.X, point.Y)

	// Structs are compared field by field
	fmt.Println(Address{"Kochi", "682001"} == c.Address)
}
