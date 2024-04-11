package main

import "fmt"

// Interface for general human behavior
type Human interface {
	Greet()
	Eat()
	Sleep()
}

// Interface for programmer specific behavior
type ProgrammerInterface interface {
	Code()
	LearnNewLanguage()
}

// Person structure with common fields and methods
type Person struct {
	Name string
}

// Greet method for Person structure
func (p *Person) Greet() {
	fmt.Printf("Hello, my name is %s\n", p.Name)
}

// Eat method for Person structure
func (p *Person) Eat() {
	fmt.Println("I am eating.")
}

// Sleep method for Person structure
func (p *Person) Sleep() {
	fmt.Println("I am sleeping.")
}

// Programmer structure with programmer specific fields and methods
type Programmer struct {
	Person
	Language string
}

// Code method for Programmer structure
func (p *Programmer) Code() {
	fmt.Printf("%s is coding in %s\n", p.Name, p.Language)
}

// LearnNewLanguage method for Programmer structure
func (p *Programmer) LearnNewLanguage() {
	fmt.Println("Learning new programming language...")
}

func main() {
	// Create instance of Programmer which contains an instance of Person
	programmer := &Programmer{
		Person: Person{
			Name: "John Doe",
		},
		Language: "Go",
	}

	// Call methods from both structures
	programmer.Greet()
	programmer.Code()
	programmer.LearnNewLanguage()

	// Cast programmer to Person and call its methods
	person := Programmer.(*Person)
	person.Greet()
	person.Eat()
	person.Sleep()
}
