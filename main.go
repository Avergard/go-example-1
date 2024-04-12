package main

import "fmt"

type person interface {
	getName() string
	getSeconName() string
	getWorkExperience() bool
}
type programmer interface {
	getBuildSomeServers() bool
	getAllDocPages() int
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

func (p *Person) getName() {

	fmt.Printf("name %s", p.Name)

}

func (p *Person) getSecondName() string {
	return p.SecondName
}
func (p *Person) getAge() uint {
	return p.Age
}

func (p *Person) IsSleep() {
	fmt.Printf("%s: dolboeb", p.Name)
	fmt.Println()

}
func (p *Person) IsGettingOld() {
	fmt.Printf("%s %s: возраст %d  год,", p.Name, p.SecondName, p.Age)
	fmt.Println()
}
func personIsEating(k Person) {
	fmt.Printf("%s%s: кушает", k.Name, k.SecondName)
	fmt.Println()
}
func (p *Person) changeAge(AgeNow uint) {
	p.Age = AgeNow
}

//PROGRAMMER

type Programmer struct {
	Person
	programmingLanguageKnowledge string
	buildSomeServers             bool
	allDocPages                  int
}

func ProgrammerIsEating(l Programmer) {
	fmt.Printf("%s%s: кушает", l.Name, l.SecondName)
	fmt.Println()
}

func (P *Programmer) getProgrammingLanguageKnowledge() string {
	return P.programmingLanguageKnowledge
}

// Меняет AllDocPages
func (p *Programmer) changeDoc(DocPagesToday int) {
	p.allDocPages = DocPagesToday

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
	fmt.Println(human1.Name, human1.SecondName, "age", human1.Age)

	human1.changeAge(23)

	fmt.Println(human1.Name, human1.SecondName, "age now", human1.Age)

	human2 := Programmer{
		Person: Person{
			Name:           "stas",
			SecondName:     "nikitov",
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
			Name:           "Qxin",
			SecondName:     "Qxan",
			Age:            54,
			Weight:         59,
			Race:           "Asian",
			usersCell:      3,
			workExperience: "26 years",
		},
		buildSomeServers:             true,
		programmingLanguageKnowledge: "senior",
		allDocPages:                  1229,
	}
	fmt.Println()
	fmt.Println(human3.Name, human3.SecondName, "doc for all time", human3.allDocPages)

	human3.changeDoc(12)

	fmt.Println(human3.Name, human3.SecondName, "doc today", human3.allDocPages)

	personIsEating(human1)
	ProgrammerIsEating(human2)
	ProgrammerIsEating(human3)

	(&human1).getName()

	human1.IsGettingOld()
	human2.IsGettingOld()
	human3.IsGettingOld()

	human1.IsSleep()
	human2.IsSleep()
	human3.IsSleep()

	human2.getProgrammingLanguageKnowledge()
	human3.getProgrammingLanguageKnowledge()

	fmt.Println(human1)
	fmt.Println(human2)
	fmt.Println(human3)
}
