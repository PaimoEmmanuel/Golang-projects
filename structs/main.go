package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}
type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}

func main() {
	user := person{firstName: "Emmanuel", lastName: "Paimo", contact: contactInfo{email: "paimoemmanuel2016@gmail.com", zipCode: 101110}}
	user.updatename("Ayomide")
	user.print()
	name := []string{"Bill"}
 
    log(name)
	str := "paimo"
	strPoint := &str
	fmt.Println(&strPoint)
	fmt.Println(&str)
}

func (pointerToPerson *person) updatename(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}
func (p person) print() {
	fmt.Println(p)

}
func log(n []string) {
	n[0] = "omooo"
}