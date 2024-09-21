package filter

import (
	"gocourse11/aquarium"
	"testing"
)

func TestCalculatePollution(t *testing.T) {
	a := aquarium.Aquarium{Chemicals: []aquarium.Chemicals{aquarium.N,
		aquarium.Mg, aquarium.Mg}}
	system := System{}
	pollution := system.CalculatePollution(a)
	if pollution != 3 {
		t.Errorf("Pollution isn't right should be 3 but was %d\n", pollution)
	}

	a = aquarium.Aquarium{Chemicals: []aquarium.Chemicals{aquarium.N,
		aquarium.Mg, aquarium.Mg, aquarium.Na, aquarium.K}}
	pollution = system.CalculatePollution(a)
	if pollution != 3 {
		t.Errorf("Pollution isn't right should be 3 but was %d\n", pollution)
	}
}

func TestFilter(t *testing.T) {
	a := aquarium.Aquarium{Chemicals: []aquarium.Chemicals{
		aquarium.N,
		aquarium.Mg,
		aquarium.Mg,
		aquarium.Mg,
		aquarium.Mg,
		aquarium.N,
	}}
	system := System{}
	pollution := system.CalculatePollution(a)
	system.Filter(pollution, &a)
	if len(a.Chemicals) != 7 {
		t.Errorf("Should be seven chemical elements but was %d", len(a.Chemicals))
	}

	a = aquarium.Aquarium{Chemicals: []aquarium.Chemicals{
		aquarium.N,
		aquarium.Mg,
		aquarium.Mg,
		aquarium.Mg,
		aquarium.Mg,
		aquarium.N,
		aquarium.N,
		aquarium.Mg,
		aquarium.Mg,
		aquarium.Mg,
		aquarium.Mg,
		aquarium.N,
	}}
	pollution = system.CalculatePollution(a)
	system.Filter(pollution, &a)
	if len(a.Chemicals) != 14 {
		t.Errorf("Should be 14 chemical elements but was %d", len(a.Chemicals))
	}
}
