package filter

import (
	"gocourse11/aquarium"
)

type System struct{}

func (System) CalculatePollution(a aquarium.Aquarium) int {
	var n, mg int
	for _, v := range a.Chemicals {
		switch v {
		case aquarium.N:
			n++
		case aquarium.Mg:
			mg++
		}
	}

	return n + mg
}

func (System) Filter(pollution int, a *aquarium.Aquarium) {
	if pollution > 5 {
		a.Chemicals = append(a.Chemicals, aquarium.K)
	}

	if pollution > 10 {
		a.Chemicals = append(a.Chemicals, aquarium.Na)
	}
}
