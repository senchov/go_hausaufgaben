package main

import (
	"testing"
)

func TestNewZoo(t *testing.T) {
	zoo := NewZoo()
	if zoo == nil {
		t.Error("Zoo is nil")
	}

	if len(zoo.Areas) != 3 {
		t.Error("Not enough areas")
	}
}

func TestFindAnimal(t *testing.T) {
	zoo := NewZoo()
	a := FindAnimal(zoo, "Timon")
	if a == nil {
		t.Error("Timon is absent")
	}

	fake := FindAnimal(zoo, "FakeTimon")
	if fake != nil {
		t.Error("FakeTimon is present")
	}
}

func TestMoveAnimal(t *testing.T) {
	zoo := NewZoo()
	MoveAnimal(zoo, "Pumbaa", Meerkat)
	animalsInMeerkat := *zoo.Areas[Mammals].Sectors[Meerkat].Animals

	if animalsInMeerkat[1].Name != "Pumbaa" {
		t.Error("Pumbaa not in place")
	}
}

func TestAddAnimalToSector(t *testing.T) {
	zoo := NewZoo()
	sector := zoo.Areas[Mammals].Sectors[Meerkat]
	AddAnimalToSector(&sector, &Animal{Name: "Test"})
	if len(*sector.Animals) != 2 {
		t.Error("Op")
	}
}
