package main

import "fmt"

type user interface {
	getName() string
	getSeconName() string
	getWorkExperience() bool
}

type Person struct {
	Name           string
	SecondName     string
	Age            uint
	Weight         uint
	Race           string
	workExperience string
	usersCell      int
}

func (p *Person) getName() string {
	return p.Name
}
func (p *Person) getSecondName() string {
	return p.SecondName
}
func (p *Person) getAge() uint {
	return p.Age
}

func (p *Person) IsSleep() {
	fmt.Printf("%s drochit", p.Name)
	fmt.Println()

}
func (p *Person) IsGettingOld() {
	fmt.Printf("%s %s %s %d  год,", p.Name, p.SecondName, ": возраст", p.Age)
	fmt.Println()
}
func personIsEating(k Person) {
	fmt.Printf("%s%s кушает", k.Name, k.SecondName)
	fmt.Println()
}

//PROGRAMMER

type Programmer struct {
	Person
	programmingLanguageKnowledge string
	buildSomeServers             bool
	allDocPages                  int
}

func ProgrammerIsEating(l Programmer) {
	fmt.Printf("%s%s кушает", l.Name, l.SecondName)
	fmt.Println()
}

func (Pr *Programmer) getProgrammingLanguageKnowledge() string {
	return Pr.programmingLanguageKnowledge
}

// Меняет AllDocPages
func (pr *Programmer) changeDoc(DocPagesToday int) {
	pr.allDocPages = DocPagesToday

}

func main() {
	human1 := Person{
		Name:           "Nikita",
		SecondName:     "Kriukov",
		Age:            21,
		Weight:         75,
		Race:           "European",
		usersCell:      1,
		workExperience: "0 years",
	}
	human2 := Programmer{
		Person: Person{
			Name:           "Uwuwuwewe",
			SecondName:     "Osas",
			Age:            34,
			Weight:         81,
			Race:           "Negroid",
			usersCell:      2,
			workExperience: "12 years",
		},
		buildSomeServers:             true,
		programmingLanguageKnowledge: "middle",
		allDocPages:                  110,
	}
	fmt.Println(human2.Name, human2.SecondName, "doc for all time", human2.allDocPages)

	human2.changeDoc(2)

	fmt.Println(human2.Name, human2.SecondName, "doc today", human2.allDocPages)

	human3 := Programmer{
		Person: Person{
			Name:           "Ching",
			SecondName:     "Chong",
			Age:            54,
			Weight:         59,
			Race:           "ASSsian",
			usersCell:      3,
			workExperience: "26 years",
		},
		buildSomeServers:             true,
		programmingLanguageKnowledge: "senior",
		allDocPages:                  1229,
	}
	fmt.Println(human3.Name, human3.SecondName, "doc for all time", human3.allDocPages)

	human3.changeDoc(12)

	fmt.Println(human3.Name, human3.SecondName, "doc today", human3.allDocPages)

	personIsEating(human1)
	ProgrammerIsEating(human2)
	ProgrammerIsEating(human3)

	human1.IsGettingOld()
	human2.IsGettingOld()
	human3.IsGettingOld()

	human1.IsSleep()
	human2.IsSleep()
	human3.IsSleep()

	human2.getProgrammingLanguageKnowledge()
	human3.getProgrammingLanguageKnowledge()

	fmt.Println(human2, human1, human3)
}
