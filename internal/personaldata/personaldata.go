package personaldata

//package main

import "fmt"

// Ниже создайте структуру Personal
type Personal struct {
	Name   string
	Weight int
	Height int
}

// Ниже создайте метод Print()
func (p Personal) Print() {
	fmt.Println("Имя:", p.Name)
	fmt.Println("Вес", p.Weight)
	fmt.Println("Рост", p.Height)
}

// проверка работоспособности кода
//func main() {
//personal := Personal{Name: "Anna", Weight: 80, Height: 180}
//fmt.Println(personal)
//personal.Print()
//}
