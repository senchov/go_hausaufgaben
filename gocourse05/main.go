package main

import (
	"fmt"
	"gocourse05/security"
)

func main() {
	records := []security.AnimalRecord{}
	fmt.Println(records)
	ptz := security.NewPTZCamera(0, 10)
	nv := security.NewNightvisionCamera(30, 6)

	ptz.Record(*security.NewConditions(false, 2, 100), &records, *security.NewAnimalRecord(security.Bear, security.Sleep))
	ptz.Record(*security.NewConditions(true, 4, 100), &records, *security.NewAnimalRecord(security.Bear, security.Sleep))
	nv.Record(*security.NewConditions(true, 5, 100), &records, *security.NewAnimalRecord(security.Tiger, security.Walking))
	ptz.Record(*security.NewConditions(false, 5, 100), &records, *security.NewAnimalRecord(security.Rat, security.Eat))
	nv.Record(*security.NewConditions(true, 50, 100), &records, *security.NewAnimalRecord(security.Tiger, security.Walking))
	nv.Record(*security.NewConditions(true, 3, 100), &records, *security.NewAnimalRecord(security.Tiger, security.Walking))

	for _, v := range records {
		fmt.Println(v)
	}
}
