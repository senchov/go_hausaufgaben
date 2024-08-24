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
		t.Logf("Dave data is 0 expexted more than zero")
	}

	mikeStr := <-mikeData
	fmt.Println(mikeStr)
	if len(mikeStr) == 0 {
		t.Logf("Mike data is 0 expexted more than zero")
	}

	jackStr := <-jackData
	fmt.Println(jackStr)
	if len(jackStr) == 0 {
		t.Logf("Jack data is 0 expexted more than zero")
	}

	uStr := <-undefinedData
	fmt.Println(uStr)
	if uStr != Undefined {
		t.Error("Data should be undefined if name is empthy")
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
		t.Error("Second cage should be closed but it oepned")
	}
}
