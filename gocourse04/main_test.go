package main

import (
	"testing"
)

func TestNewZoo(t *testing.T) {
	zoo := NewZoo(CreateAreas())
	if zoo == nil {
		t.Error("Zoo is nil")
	}

	if len(zoo.Areas) != 3 {
		t.Logf("Not enough areas expected=%v actual=%v", 3, len(zoo.Areas))
	}
}

func TestFindAnimal(t *testing.T) {
	zoo := NewZoo(CreateAreas())
	a := zoo.FindAnimal("Timon")
	if a == nil {
		t.Error("Timon is absent")
	}

	fake := zoo.FindAnimal("FakeTimon")
	if fake != nil {
		t.Error("FakeTimon is present")
	}
}

func TestMoveAnimal(t *testing.T) {
	zoo := NewZoo(CreateAreas())
	zoo.MoveAnimal("Pumbaa", Meerkat)
	animalsInMeerkat := *zoo.Areas[Mammals].Sectors[Meerkat].Animals

	if animalsInMeerkat[1].Name != "Pumbaa" {
		t.Error("Pumbaa not in place")
	}
}

func TestAddAnimalToSector(t *testing.T) {
	zoo := NewZoo(CreateAreas())
	sector := zoo.Areas[Mammals].Sectors[Meerkat]
	sector.AddAnimal(Animal{Name: "Test"})
	if len(*sector.Animals) != 2 {
		t.Logf("Not right amount of animals: want=%v, but got=%v", 2, len(*sector.Animals))
	}
}

func TestFeedAnimal(t *testing.T) {
	zoo := NewZoo(CreateAreas())
	a := zoo.FindAnimal("Donald")
	a.FeedAnimal()
	if !a.IsAte {
		t.Error("Donald is hungry")
	}
}
