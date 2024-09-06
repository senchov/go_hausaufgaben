package main

import (
	"testing"
)

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
		if v > i+2 || v < i {
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
