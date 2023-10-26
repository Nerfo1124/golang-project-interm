package main

import "fmt"

type Employee struct {
	id       int
	name     string
	vacation bool
}

func (e *Employee) SetId(id int) {
	e.id = id
}

func (e *Employee) SetName(name string) {
	e.name = name
}

func (e *Employee) GetId() int {
	return e.id
}

func (e *Employee) GetName() string {
	return e.name
}

func NewEmployee(id int, name string, vacation bool) *Employee {
	return &Employee{
		id:       id,
		name:     name,
		vacation: vacation,
	}
}

func main() {
	e := Employee{}
	fmt.Printf("%v\n", e)

	e.id = 1
	e.name = "Nerfo"
	fmt.Printf("%v\n", e)

	e.SetId(2)
	e.SetName("NerfoDos")
	fmt.Printf("%v\n", e)
	fmt.Println(e.GetId())
	fmt.Println(e.GetName())

	e2 := Employee{id: 10, name: "Mario", vacation: true}
	fmt.Printf("%v\n", e2)

	e3 := new(Employee)
	fmt.Printf("%v\n", *e3)

	e4 := NewEmployee(1, "Luigi", true)
	fmt.Printf("%v\n", *e4)
}
