package feed

import "testing"

func TestSausagesTrough(t *testing.T) {
	sausages := SausagesTrough{Trough: MeetTrought{}}
	if sausages.Food() != "Sausages from Ham, pork and chicken" {
		t.Errorf("Food isn't correct should be %s \n but was %s", "Sausages from Ham, pork and chicken", sausages.Food())
	}
}

func TestMeetTrought(t *testing.T) {
	meet := MeetTrought{}
	if meet.Food() != "Ham, pork and chicken" {
		t.Errorf("Food isn't correct should be %s \n", "Ham, pork and chicken")
	}
}

func TestVegetablesTrought(t *testing.T) {
	vg := VegetablesTrought{}
	if vg.Food() != "Carrots, potatoes and cucumbers" {
		t.Errorf("Food isn't correct should be %s \n", "Carrots, potatoes and cucumbers")
	}
}

func TestFishTrough(t *testing.T) {
	fish := FishTrough{}
	if fish.Food() != "Code and salmon" {
		t.Errorf("Food isn't correct should be %s \n", "Code and salmon")
	}
}
