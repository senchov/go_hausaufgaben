package filter

import (
	"gocourse11/aquarium"
)

//import "gocourse11/aqurium"
// Потрібно вирахувати рівень забруднення і запустити системи фільтрації в зміненому режимі,
//також додати до води необхідні хімічні речовини, які покращать її стан.

type System struct{}

func (System) CalculatePollution(a aquarium.Aquarium) int {
	var n, mg int
	for _, v := range a.Chemichals {
		switch v {
		case aquarium.N:
			n += 1
		case aquarium.Mg:
			mg += 1
		}
	}

	return n + mg
}

func (System) Filter(polution int, a *aquarium.Aquarium) {
	if polution > 5 {
		a.Chemichals = append(a.Chemichals, aquarium.K)
	}

	if polution > 10 {
		a.Chemichals = append(a.Chemichals, aquarium.Na)
	}
}
