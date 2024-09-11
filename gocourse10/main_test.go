package main

import "testing"

func TestFood(t *testing.T) {
	meet := MeetTrought{}
	if meet.Food() != "Ham, pork and chicken" {
		t.Errorf("Food isn't corect should be %s \n", "Ham, pork and chicken")
	}

	vg := VegetablesTrought{}
	if vg.Food() != "Carrots, potatoes and cucumbers" {
		t.Errorf("Food isn't corect should be %s \n", "Carrots, potatoes and cucumbers")
	}

	fish := FishTrough{}
	if fish.Food() != "Code and salmon" {
		t.Errorf("Food isn't corect should be %s \n", "Code and salmon")
	}

	sauseges := SousegesTrough{Trough: MeetTrought{}}
	if sauseges.Food() != "Sauseges from Ham, pork and chicken" {
		t.Errorf("Food isn't corect should be %s \n but was %s", "Sauseges from Ham, pork and chicken", sauseges.Food())
	}
}
