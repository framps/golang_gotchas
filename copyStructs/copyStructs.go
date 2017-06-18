package main

// Copy structs by using github.com/jinzhu/copier
//
// Copyright (C) 2017 framp at linux-tips-and-tricks dot de

import (
	"fmt"

	"github.com/jinzhu/copier"
)

// User -
type User struct {
	Name string
	Role string
	Age  int32
}

// DoubleAge -
func (user *User) DoubleAge() int32 {
	return 2 * user.Age
}

// Employee -
type Employee struct {
	Name      string
	Age       int32
	DoubleAge int32
	EmployeID int64
	SuperRule string
}

// Role -
func (employee *Employee) Role(role string) {
	employee.SuperRule = "Super " + role
}

// Color -
type Color int

const (
	white Color = iota
	black
	red
	green
	blue
	pink
)

// Animal -
type Animal struct {
	Name  string
	Age   int
	Color Color
}

// Elephant -
type Elephant struct {
	Name string
	Age  int
}

// Pig -
type Pig struct {
	Name  string
	Color Color
}

func main() {
	var (
		user  = User{Name: "Jinzhu", Age: 18, Role: "Admin"}
		users = []User{{Name: "Jinzhu", Age: 18, Role: "Admin"},
			{Name: "jinzhu 2", Age: 30, Role: "Dev"}}
		employee  = Employee{}
		employees = []Employee{}
	)

	copier.Copy(&employee, &user)

	fmt.Printf("%+v \n", employee)
	// Employee{
	//    Name: "Jinzhu",           // Copy from field
	//    Age: 18,                  // Copy from field
	//    DoubleAge: 36,            // Copy from method
	//    EmployeeId: 0,            // Ignored
	//    SuperRule: "Super Admin", // Copy to method
	// }

	// Copy struct to slice
	copier.Copy(&employees, &user)

	fmt.Printf("%+v \n", employees)
	// []Employee{
	//   {Name: "Jinzhu", Age: 18, DoubleAge: 36, EmployeId: 0, SuperRule: "Super Admin"}
	// }

	// Copy slice to slice
	employees = []Employee{}
	copier.Copy(&employees, &users)

	fmt.Printf("%+v \n", employees)
	// []Employee{
	//   {Name: "Jinzhu", Age: 18, DoubleAge: 36, EmployeId: 0, SuperRule: "Super Admin"},
	//   {Name: "jinzhu 2", Age: 30, DoubleAge: 60, EmployeId: 0, SuperRule: "Super Dev"},
	// }

	elephant := Elephant{Name: "Dumbo", Age: 42}
	pig := Pig{Name: "Piggy", Color: pink}

	fmt.Printf("Pig: %+v\n", pig)
	// Pig: {Name:Piggy Color:5}

	fmt.Printf("Elephant: %+v\n", elephant)
	// Elephant: {Name:Dumbo Age:42}

	var animal Animal
	copier.Copy(&animal, &elephant)
	fmt.Printf("Animal from elephant: %+v\n", animal)
	// Animal from elephant: {Name:Dumbo Age:42 Color:0}

	copier.Copy(&animal, &pig)
	fmt.Printf("Animal from pig: %+v\n", animal)
	// Animal from pig: {Name:Piggy Age:42 Color:5}

	animal = Animal{Name: "Animal", Age: 4711}
	fmt.Printf("Animal: %+v\n", animal)
	// Animal: {Name:Animal Age:4711 Color:0}

	elephant = Elephant{Name: "Dumbo", Age: 42}
	pig = Pig{Name: "Piggy", Color: pink}
	fmt.Printf("Pig: %+v\n", pig)
	// Pig: {Name:Piggy Color:5}
	fmt.Printf("Elephant: %+v\n", elephant)
	// Elephant: {Name:Dumbo Age:42}

	copier.Copy(&elephant, &animal)
	fmt.Printf("Elephant from animal: %+v\n", animal)
	// Elephant from animal: {Name:Animal Age:4711 Color:0}

	copier.Copy(&pig, &animal)
	fmt.Printf("Pig form animal: %+v\n", pig)
	//	Pig form animal: {Name:Animal Color:0}

}
