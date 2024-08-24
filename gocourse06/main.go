package main

import "fmt"

func main() {
	fmt.Println("start")
	data := make(chan string)
	go CollectData(Animal{Name: "Dave", Mood: Sad, Hunger: 56, Health: 22}, data)
	res := <-data
	fmt.Println(res)
}

// Створіть горутину для кожної тварини в зоопарку. Кожна горутина збирає дані про стан тварини
// (наприклад, рівень здоров'я, голод, настрій) і відправляє їх через канал до центральної системи моніторингу.
type Mood string

const (
	Sad  Mood = "Sad"
	Cool Mood = "Cool"

	Undefined string = "Undefined"
)

type Animal struct {
	Name   string
	Health int
	Hunger int
	Mood
}

func CollectData(a Animal, data chan string) {
	if a.Name == "" {
		data <- Undefined
		return
	}

	r := fmt.Sprintf("Animal->%s is mood->%s health->%v hunger->%v", a.Name, a.Mood, a.Health, a.Hunger)
	data <- r
}

// Керування доступом до вольєрів: Імплементуйте горутину, яка контролює доступ до вольєрів,
// використовуючи канали для отримання запитів на відкриття/закриття.
func IsShouldOpenCage(hour int, data chan bool) {
	if hour < 8 || hour > 18 {
		data <- false
	} else {
		data <- true
	}

}

// Розробіть горутини для управління автоматичними кормушками, які відправляють статус кормушки (порожня/повна)
// через канал.
type Trough struct {
	ID     int
	Amount int
}

func CheckTrough(t Trough, full chan int, empty chan int) {
	if t.Amount > 50 {
		full <- t.ID
	} else {
		empty <- t.ID
	}
}
