package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(4)
	go CollectData(Animal{Name: "Jo", Mood: Sad, Hunger: 42, Health: 11}, &wg)
	go CollectData(Animal{Name: "Tim", Mood: Cool, Hunger: 21, Health: 44}, &wg)
	go CollectData(Animal{Name: "Jack", Mood: Sad, Hunger: 16, Health: 8}, &wg)
	go CollectData(Animal{Mood: Sad, Hunger: 0, Health: 56}, &wg)

	wg.Wait()
	fmt.Println("Data collected")
	fmt.Println("-----------")

	full := make(chan int, 2)
	empty := make(chan int, 1)
	go CheckTrough(Trough{ID: 1, Amount: 100}, full, empty)
	go CheckTrough(Trough{ID: 2, Amount: 60}, full, empty)
	go CheckTrough(Trough{ID: 3, Amount: 40}, full, empty)
	go CheckTrough(Trough{ID: 4, Amount: 0}, full, empty)

	fullIDs := []int{<-full, <-full}
	for i, v := range fullIDs {
		fmt.Printf("full id=>%d index=%d \n", v, i)
	}

	emptyIDs := []int{<-empty, <-empty}
	for i, v := range emptyIDs {
		fmt.Printf("empty id=>%d index=%d\n", v, i)
	}
	fmt.Println("-----------")

	firstCageLock := make(chan bool)
	secondCageLock := make(chan bool)
	thirdCageLock := make(chan bool)
	fourthCageLock := make(chan bool)
	go IsShouldOpenCage(5, firstCageLock)
	go IsShouldOpenCage(10, secondCageLock)
	go IsShouldOpenCage(12, thirdCageLock)
	go IsShouldOpenCage(3, fourthCageLock)
	first := <-firstCageLock
	PrintLockState(first, "First")
	second := <-secondCageLock
	PrintLockState(second, "Second")
	third := <-thirdCageLock
	PrintLockState(third, "Third")
	fourth := <-fourthCageLock
	PrintLockState(fourth, "Fourth")
}

func PrintLockState(isOpen bool, lockName string) {
	if isOpen {
		fmt.Printf("%s lock is open \n", lockName)
	} else {
		fmt.Printf("%s lock is closed \n", lockName)
	}
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

func CollectData(a Animal, wg *sync.WaitGroup) {

	defer wg.Done()
	fmt.Printf("Animal->%s is mood->%s health->%v hunger->%v \n", a.Name, a.Mood, a.Health, a.Hunger)
}

// Керування доступом до вольєрів: Імплементуйте горутину, яка контролює доступ до вольєрів,
// використовуючи канали для отримання запитів на відкриття/закриття.
func IsShouldOpenCage(hour int, data chan<- bool) {
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

func CheckTrough(t Trough, full chan<- int, empty chan<- int) {
	if t.Amount > 50 {
		full <- t.ID
	} else {
		empty <- t.ID
	}
}
