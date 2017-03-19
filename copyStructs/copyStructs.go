package main

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

func main() {
	var (
		user  = User{Name: "Jinzhu", Age: 18, Role: "Admin"}
		users = []User{{Name: "Jinzhu", Age: 18, Role: "Admin"},
			{Name: "jinzhu 2", Age: 30, Role: "Dev"}}
		employee  = Employee{}
		employees = []Employee{}
	)

	copier.Copy(&employee, &user)

	fmt.Printf("%#v \n", employee)
	// Employee{
	//    Name: "Jinzhu",           // Copy from field
	//    Age: 18,                  // Copy from field
	//    DoubleAge: 36,            // Copy from method
	//    EmployeeId: 0,            // Ignored
	//    SuperRule: "Super Admin", // Copy to method
	// }

	// Copy struct to slice
	copier.Copy(&employees, &user)

	fmt.Printf("%#v \n", employees)
	// []Employee{
	//   {Name: "Jinzhu", Age: 18, DoubleAge: 36, EmployeId: 0, SuperRule: "Super Admin"}
	// }

	// Copy slice to slice
	employees = []Employee{}
	copier.Copy(&employees, &users)

	fmt.Printf("%#v \n", employees)
	// []Employee{
	//   {Name: "Jinzhu", Age: 18, DoubleAge: 36, EmployeId: 0, SuperRule: "Super Admin"},
	//   {Name: "jinzhu 2", Age: 30, DoubleAge: 60, EmployeId: 0, SuperRule: "Super Dev"},
	// }
}
