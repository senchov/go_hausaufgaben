package main

import (
	"fmt"
	"gocourse05/security"
	"math/rand"
)

func main() {
	records := []security.AnimalRecord{}
	fmt.Println(records)
	ptz := security.NewPTZCamera(250, 10)
	nv := security.NewNightvisionCamera(30, 6)

	records = ptz.Record(*security.NewConditions(false, 2, 100), records, *security.NewAnimalRecord(security.Bear, security.Sleep))
	records = ptz.Record(*security.NewConditions(true, 4, 100), records, *security.NewAnimalRecord(security.Bear, security.Sleep))
	records = nv.Record(*security.NewConditions(true, 5, 100), records, *security.NewAnimalRecord(security.Tiger, security.Walking))
	records = ptz.Record(*security.NewConditions(false, 5, 100), records, *security.NewAnimalRecord(security.Rat, security.Eat))
	records = nv.Record(*security.NewConditions(true, 50, 100), records, *security.NewAnimalRecord(security.Tiger, security.Walking))
	records = nv.Record(*security.NewConditions(true, 3, 100), records, *security.NewAnimalRecord(security.Tiger, security.Walking))
	records = ptz.Record(*security.NewConditions(false, 150, 100), records, *security.NewAnimalRecord(security.Rat, security.Walking))
	records = ptz.Record(*security.NewConditions(false, 4, 100), records, *security.NewAnimalRecord(security.Bear, security.Sleep))
	records = nv.Record(*security.NewConditions(true, 55, 0), records, *security.NewAnimalRecord(security.Tiger, security.Walking))
	records = ptz.Record(*security.NewConditions(false, 27, 100), records, *security.NewAnimalRecord(security.Rat, security.Eat))
	records = nv.Record(*security.NewConditions(true, 33, 100), records, *security.NewAnimalRecord(security.Rat, security.Walking))
	records = nv.Record(*security.NewConditions(true, 3, 100), records, *security.NewAnimalRecord(security.Bear, security.Eat))

	for _, v := range records {
		v.Print()
	}

	animals := []security.AnimalType{security.Bear, security.Rat, security.Tiger}
	actions := []security.AnimalAction{security.Eat, security.Sleep, security.Walking}
	recorders := []Recorder{ptz, nv}
	rotators := []Rotater{ptz, nv}

	for i := 0; i < 100; i++ {
		isNight := rand.Intn(2) == 1
		distToAnimal := security.Range(rand.Float32() * 20)
		battery := rand.Intn(101)

		animal := animals[rand.Intn(3)]
		action := actions[rand.Intn(3)]
		for j := 0; j < 2; j++ {
			rotators[j].Rotate(security.Angle(rand.Float32() * 180))
			records = recorders[j].Record(*security.NewConditions(isNight, distToAnimal, battery), records, *security.NewAnimalRecord(animal, action))
		}
	}

	for _, v := range records {
		v.Print()
	}
}

type Rotater interface {
	Rotate(angle security.Angle)
}

type Recorder interface {
	Record(c security.Condition, records []security.AnimalRecord, rec security.AnimalRecord) []security.AnimalRecord
}
