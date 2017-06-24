package main

import "fmt"

// User -
type User struct {
	Name string
}

func hasChanged(u User) {
	if u.Name != "Leto" {
		fmt.Printf("!!! Name changed to %s\n", u.Name)
	} else {
		fmt.Println("??? Name NOT changed")
	}
}

// ModifyPtr - modify with pointer interface
func ModifyPtr(u *User) {
	fmt.Printf("Receiving u as pointer\nType: %T: \tAddress: %p \tValue: %v\n", u, &u, u)
	u = &User{Name: "Paul"} // linter warning : ineffectual assignment to u
}

// ModifyValue - modify with value interface
func ModifyValue(u User) {
	fmt.Printf("Receiving u as value\nType: %T: \tAddress: %p \tValue: %v\n", u, &u, u)
	u.Name = "Duncan"
}

// ModifyPtrPtr - modify with pointer pointer interface
func ModifyPtrPtr(u **User) {
	fmt.Printf("Receiving u as pointer to pointer:\nType: %T: \tAddress: %p \tValue: %v\n", u, &u, u)
	*u = &User{Name: "Bob"}
}

// ModifyPtr2 - modify with pointer interface
func ModifyPtr2(u *User) {
	fmt.Printf("Receiving u as pointer\nType: %T: \tAddress: %p \tValue: %v\n", u, &u, u)
	u.Name = "Paul"
}

func main() {

	u := &User{Name: "Leto"}
	fmt.Printf("Passing u as value:\nType: %T: \tAddress: %p \tValue: %v\n", u, &u, u)
	ModifyPtr(u)
	hasChanged(*u)
	// Name not changed
	// u -> User, u is passed by value/copy as u' -> User and u' will be modified to point to NewUser,
	//   but u will still point to User

	v := User{Name: "Leto"}
	fmt.Printf("Passing v as value:\nType: %T: \tAddress: %p \tValue: %v\n", v, &v, v)
	ModifyValue(v)
	hasChanged(v)
	// Name not changed, user passed by value/copy

	w := &User{Name: "Leto"}
	fmt.Printf("Passing w as pointer:\nType: %T: \tAddress: %p \tValue: %v\n", w, &w, w)
	ModifyPtrPtr(&w)
	hasChanged(*w)
	// Name changed, address of var w which points to user is passed and will be changed
	// w -> User, &w is passed by value/copy as w' -> w -> User and contents of w which points to User
	//    can be changed to NewUser

	x := &User{Name: "Leto"}
	fmt.Printf("Passing u as value:\nType: %T: \tAddress: %p \tValue: %v\n", x, &x, x)
	ModifyPtr2(x)
	hasChanged(*x)
	// Name changed, address to user contents (name) will be changed, no new user will be returned
	// x -> User, x is passed by value/copy x'-> User and User Name can be modified

	y := User{Name: "Leto"}
	fmt.Printf("Passing u as value:\nType: %T: \tAddress: %p \tValue: %v\n", y, &y, y)
	ModifyPtr2(&y)
	hasChanged(y)
	// Name changed, address to user passed and contents (name) will be changed, no new user will be returned
	// y -> User, &y is passed by valued and y' -> User and User Name can be modified

}

/* Run result:

Passing u as value:
Type: *main.User: 	Address: 0xc42002c020 	Value: &{Leto}
Receiving u as pointer
Type: *main.User: 	Address: 0xc42002c030 	Value: &{Leto}
??? Name NOT changed
Passing v as value:
Type: main.User: 	Address: 0xc42000a3a0 	Value: {Leto}
Receiving u as value
Type: main.User: 	Address: 0xc42000a3d0 	Value: {Leto}
??? Name NOT changed
Passing w as pointer:
Type: *main.User: 	Address: 0xc42002c038 	Value: &{Leto}
Receiving u as pointer to pointer:
Type: **main.User: 	Address: 0xc42002c040 	Value: 0xc42002c038
!!! Name changed to Bob
Passing u as value:
Type: *main.User: 	Address: 0xc42002c048 	Value: &{Leto}
Receiving u as pointer
Type: *main.User: 	Address: 0xc42002c050 	Value: &{Leto}
!!! Name changed to Paul
Passing u as value:
Type: main.User: 	Address: 0xc42000a4c0 	Value: {Leto}
Receiving u as pointer
Type: *main.User: 	Address: 0xc42002c058 	Value: &{Leto}
!!! Name changed to Paul

*/
