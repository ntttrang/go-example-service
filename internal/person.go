package internal

import "fmt"

type Person struct {
	Name string
	Age  int
}

// use Stringer interface
func (p Person) String() string {
	return fmt.Sprintf("%v (%v years)", p.Name, p.Age)
}

// Check Person co implement Stringer hay khong?
// TODO:
// func MyPrintln(p Person) {
// 	if _, ok := p.(fmt.Stringer); ok {
// 		fmt.Println(p.String())
// 	} else {
// 		fmt.Println("No")
// 	}
// }
