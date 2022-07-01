package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	// manisha := person{"Manisha", "C"}
	// fmt.Println(manisha.firstName, manisha.lastName)
	// var person1 person
	// fmt.Printf("%+v", person1) // to print with field names

	// update values of strut
	// person1.firstName = "Mani"
	// person1.lastName = "C"
	// fmt.Printf("\n%+v", person1)

	// // Embedding struct
	// person2 := person{"Jhon", "Carter", contactInfo{"jhon@gmail.com", 212}}
	// person2.print()
	// // person2Pointer := &pers
	// person2.updateName("Deck")
	// person2.print()

	user := User{Name: "Bob", Age: 22, Active: true, lastLoginAt: "today"}
	fmt.Println(user)
	u, err := json.Marshal(user)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(u)) // {"full_name":"Bob"}
	fmt.Println(user)
	err1 := json.Unmarshal(u, &user)
	if err1 != nil {
		panic(err1)
	}
	fmt.Println(user)
}

type person struct {
	firstName string
	lastName  string
	constant  contactInfo
}

type contactInfo struct {
	email   string
	zipCode int
}

// struct with receiver
func (p person) print() {
	fmt.Printf("\n%+v", p)
}

func (p *person) updateName(newLasttName string) {
	(*p).lastName = newLasttName
}

// func (pointerToPersson *person) updateName(newFrirstName string) {
// 	(*pointerToPersson).firstName = newFrirstName
// }

// JSON Marshal Struct with JSON Tags
type User struct {
	Name        string `json:"full_name"`
	Age         int    `json:"age,omitempty"` // the field will be disarcarded in the JSON output if the value has a zero value.
	Active      bool   `json:"-"`             // the field will be removed altogether from the JSON output. This is used when you want to keep the field in your Go struct but not in your JSON output
	lastLoginAt string
}
