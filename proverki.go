package main

import "fmt"

type Person struct {
	Name           string
	SecondName     string
	Age            uint
	Weight         uint
	Race           string
	workExperience uint
	usersCell      int
}

func (p *Person) IsSleep() {
	fmt.Printf("%s Спит", p.Name)
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

func (p *Person) aPersonLearns() {
	fmt.Printf("%s %s Учится", p.Name, p.SecondName)
	fmt.Println()
}
func (p *Person) thePersonIsPracticing() {
	fmt.Printf("%s %s Практикует", p.Name, p.SecondName)
	fmt.Println()
}
func (p *Person) aPersonTeaches() {
	fmt.Printf("%s %s Обучает", p.Name, p.SecondName)
	fmt.Println()
}
func (p *Person) usersСell() {
	fmt.Printf("%d   номер на диске", p.usersCell)
}

func main() {
	human1 := Person{
		Name:           "Nikita",
		SecondName:     "Kriukov",
		Age:            21,
		Weight:         75,
		Race:           "European",
		workExperience: 2,
		usersCell:      1,
	}
	human2 := Person{
		Name:           "Uwuwuwewe",
		SecondName:     "Osas",
		Age:            34,
		Weight:         81,
		Race:           "Negroid",
		workExperience: 12,
		usersCell:      2,
	}
	human3 := Person{
		Name:           "Ching",
		SecondName:     "Chong",
		Age:            54,
		Weight:         59,
		Race:           "ASSsian",
		workExperience: 25,
		usersCell:      3,
	}

	personIsEating(human1)
	personIsEating(human2)
	personIsEating(human3)

	human1.aPersonLearns()
	human2.thePersonIsPracticing()
	human3.aPersonTeaches()

	human1.IsGettingOld()
	human2.IsGettingOld()
	human3.IsGettingOld()

	human1.IsSleep()
	human2.IsSleep()
	human3.IsSleep()

}
