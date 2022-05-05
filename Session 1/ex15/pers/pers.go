package pers

type Person struct {
    firstName string
    lastName string
}

func (p *Person) FirstName() string {   // getter
    return p.firstName
}

func (p *Person) SetFirstName(newName string) { // setter
    p.firstName = newName
}
