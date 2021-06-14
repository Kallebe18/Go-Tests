// +build windows
// +build 386

package main

import "fmt"

type Animal interface {
	getOld() int
	doSomething(action string)
}

type Person struct {
	name string
	age int
}

type Dog struct {
	name string
	breed string
	age int
}

func (person Person) getOld() int {
	person.age += 1
	return 0
}

// with * the value of created object is changed outside the function
// without * the value is changed only inside the function
//   (person *Person) theres a difference right here 
func (person Person) doSomething(action string) {
	switch action {
		case "falar":
			fmt.Println("Olá")
		default:
			fmt.Println("Esta ação não está definida.")
	}
}

func (dog Dog) getOld() int {
	dog.age += 1
	return dog.age
}

func (dog Dog) doSomething(action string) {
	switch action {
		case "latir":
			fmt.Println("au au!")
		default:
			fmt.Println("Esta ação não está definida.")
	}
}

func callAnimalActions(animal Animal) {
	fmt.Println(animal.getOld())
	animal.doSomething("falar")
}

func main() {
	var dog1 Dog
	dog1.age = 0
	dog1.name = "Bob"
	dog1.breed = "labrador"

	var person1 Person
	person1.age = 0
	person1.name = "João"
	person1.getOld()
	person1.getOld()
	person1.getOld()
	fmt.Println(person1.age)


	// callAnimalActions(person1)
	// callAnimalActions(dog1)
}
