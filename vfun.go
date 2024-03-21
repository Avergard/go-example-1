package main

import "fmt"

func main() {
	//max min
	s := min("two", "one", "three")
	fmt.Println(s)
	y := max("two", "one", "three")
	fmt.Println(y)

	f := "My name is Max"
	g := 4 - 1
	make := func() int {
		return 14
	}
	fmt.Println(f, g, make())

	x := min(9.99, 3.14, 5.27)
	fmt.Println(x)

	t := min("one", "two", "three")
	fmt.Println(t)

	type ID int

	id1 := ID(7)
	id2 := ID(42)
	id := max(id1, id2)
	fmt.Println(id)
	fmt.Println(min(10))

	fmt.Println(min(10, 9))

	fmt.Println(min(10, 9, 8))

	//clear

	m := map[string]int{"one": 1, "two": 2, "three": 3}
	clear(m)

	fmt.Printf("%#v\n", m)

	l := []string{"one", "two", "three"}
	clear(l)

	fmt.Printf("%#v\n", l)

	k := []int{1, 2, 3}
	clear(k)

	fmt.Printf("%#v\n", k)

	/*func customClear[T []string | map[string]int](container T) {
	clear(container)


		s := []string{"one", "two", "three"}
		customClear(s)
		fmt.Printf("%#v\n", s)
		// []string{"", "", ""}

		m := map[string]int{"one": 1, "two": 2, "three": 3}
		customClear(m)
		fmt.Printf("%#v\n", m)
		// map[string]int{}




	append     добавляет значения в срез
	clear      удаляет или зануляет элементы контейнера

	close      закрывает канал

	complex    создают и разбирают комплексные числа
	real
	imag

	delete     удаляет элемент карты по ключу

	len        возвращает длину контейнера
	cap        возвращает вместимость контейнера

	make       создает новый срез, карту или канал
	new        выделяет память под переменную

	min        выбирает минимальный из переданных аргументов
	max        выбирает максимальный из переданных аргументов

	panic      создают и обрабатывают панику
	recover

	print      печатают аргументы
	println
	*/

}
