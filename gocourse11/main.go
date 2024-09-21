package main

import (
	"fmt"
	"gocourse11/aquarium"
	"gocourse11/filter"
)

func main() {

	first := aquarium.New(aquarium.WithAnimal(aquarium.Cat),
		aquarium.WithChemicals([]aquarium.Chemicals{
			aquarium.N,
			aquarium.Mg,
			aquarium.Mg,
			aquarium.N}),
		aquarium.WithSaltLevel(5),
		aquarium.WithSize(10))
	RunFilter(first, "first")

	second := aquarium.New(aquarium.WithAnimal(aquarium.Cat),
		aquarium.WithChemicals([]aquarium.Chemicals{
			aquarium.N,
			aquarium.Mg,
			aquarium.Mg,
			aquarium.N,
			aquarium.Mg,
			aquarium.Mg,
			aquarium.N}),
		aquarium.WithSaltLevel(5),
		aquarium.WithSize(10))
	RunFilter(second, "second")

	third := aquarium.New(aquarium.WithAnimal(aquarium.Cat),
		aquarium.WithChemicals([]aquarium.Chemicals{
			aquarium.N,
			aquarium.Mg,
			aquarium.Mg,
			aquarium.N,
			aquarium.Mg,
			aquarium.Mg,
			aquarium.N,
			aquarium.N,
			aquarium.Mg,
			aquarium.Mg,
			aquarium.N,
			aquarium.Mg}),
		aquarium.WithSaltLevel(3),
		aquarium.WithSize(8))
	RunFilter(third, "third")
}

func RunFilter(a *aquarium.Aquarium, name string) {
	filterSystem := filter.System{}
	pollution := filterSystem.CalculatePollution(*a)
	if pollution <= 5 {
		fmt.Printf("Aquarium %s is clean\n", name)
		return
	}

	filterSystem.Filter(pollution, a)
	fmt.Printf("Aquarium %s is polluted with value of %d\n", name, pollution)
	if pollution > 10 {
		fmt.Printf("Aquarium %s filtered with chemicals K and Na\n", name)
		return
	}

	fmt.Printf("Aquarium %s filtered with chemical K\n", name)
}
