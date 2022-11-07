package greet

import "fmt"

type person struct {
	id    int
	name  string
	hobby string
}

func NewPerson() person {
	return person{}
}

func (p person) SetId(id int) {
	p.id = id
}
func (p person) SetName(name string) {
	p.name = name
}
func (p person) SetHobby(hobby string) {
	p.hobby = hobby
}

func (p person) GetId() int {
	return p.id
}

func (p person) GetName() string {
	return p.name
}

func (p person) GetHobby() string {
	return p.hobby
}

func (p person) ToString() string {
	return fmt.Sprintf("Halo, nama saya %s, hobi saya adalah %s", p.name, p.hobby)
}

func HelloFromToni() {
	fmt.Println("Hai Semua, ini terupdate!")
}
