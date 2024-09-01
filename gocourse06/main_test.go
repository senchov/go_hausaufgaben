package main

import (
	"fmt"
	"testing"
)

func TestCollectData(t *testing.T) {
	daveData := make(chan string)
	mikeData := make(chan string)
	jackData := make(chan string)
	undefinedData := make(chan string)
	go CollectData(Animal{Name: "Dave", Mood: Sad, Hunger: 56, Health: 22}, daveData)
	go CollectData(Animal{Name: "Mike", Mood: Cool, Hunger: 90, Health: 44}, mikeData)
	go CollectData(Animal{Name: "Jack", Mood: Sad, Hunger: 0, Health: 56}, jackData)
	go CollectData(Animal{Mood: Sad, Hunger: 0, Health: 56}, undefinedData)
	daveStr := <-daveData
	fmt.Println(daveStr)
	if len(daveStr) == 0 {
		t.Errorf("Dave data is 0 expected more than zero")
	}

	mikeStr := <-mikeData
	fmt.Println(mikeStr)
	if len(mikeStr) == 0 {
		t.Errorf("Mike data is 0 expected more than zero")
	}

	jackStr := <-jackData
	fmt.Println(jackStr)
	if len(jackStr) == 0 {
		t.Errorf("Jack data is 0 expected more than zero")
	}

	uStr := <-undefinedData
	fmt.Println(uStr)
	if uStr != Undefined {
		t.Error("Data should be undefined if name is empty")
	}
}

func TestIsShouldOpenCage(t *testing.T) {
	firstCageLock := make(chan bool)
	secondCageLock := make(chan bool)
	go IsShouldOpenCage(5, firstCageLock)
	go IsShouldOpenCage(10, secondCageLock)
	f := <-firstCageLock
	if f {
		t.Error("First cage should be locked but it opened")
	}

	s := <-secondCageLock
	if !s {
		t.Error("Second cage should be closed but it opened")
	}
}

func TestCheckTrough(t *testing.T) {
	full := make(chan int, 2)
	empty := make(chan int, 1)
	go CheckTrough(Trough{ID: 1, Amount: 100}, full, empty)
	go CheckTrough(Trough{ID: 2, Amount: 60}, full, empty)
	go CheckTrough(Trough{ID: 3, Amount: 40}, full, empty)
	go CheckTrough(Trough{ID: 4, Amount: 0}, full, empty)

	fullIDs := []int{
		<-full,
		<-full,
	}
	for i, v := range fullIDs {
		if v > i+1 || v < i {
			t.Errorf("Wrong id->%v, expected more then %v and less then %v", v, i+1, i)
		}
	}

	emptyIDs := []int{
		<-empty,
		<-empty,
	}
	for i, v := range emptyIDs {
		if v > i+4 || v < i+2 {
			t.Errorf("Wrong id->%v, expected more then %v and less then %v", v, i+4, i+2)
		}
	}
}
