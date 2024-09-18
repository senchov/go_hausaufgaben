package main

import (
	"fmt"
	"gocourse05/animal"
	"gocourse05/security"
	"math/rand"
)

func main() {
	records := []animal.AnimalRecord{}
	fmt.Println(records)
	ptz := security.NewPTZCamera(250, 10)
	nv := security.NewNightvisionCamera(30, 6)
	fc := security.NewFixedCamera(23)

	records = ptz.Record(*security.NewCondition(false, 2, 100), records, *animal.NewAnimalRecord(animal.Bear, animal.Sleep))
	records = ptz.Record(*security.NewCondition(true, 4, 100), records, *animal.NewAnimalRecord(animal.Bear, animal.Sleep))
	records = nv.Record(*security.NewCondition(true, 5, 100), records, *animal.NewAnimalRecord(animal.Tiger, animal.Walking))
	records = ptz.Record(*security.NewCondition(false, 5, 100), records, *animal.NewAnimalRecord(animal.Rat, animal.Eat))
	records = nv.Record(*security.NewCondition(true, 50, 100), records, *animal.NewAnimalRecord(animal.Tiger, animal.Walking))
	records = nv.Record(*security.NewCondition(true, 3, 100), records, *animal.NewAnimalRecord(animal.Tiger, animal.Walking))
	records = ptz.Record(*security.NewCondition(false, 150, 100), records, *animal.NewAnimalRecord(animal.Rat, animal.Walking))
	records = ptz.Record(*security.NewCondition(false, 4, 100), records, *animal.NewAnimalRecord(animal.Bear, animal.Sleep))
	records = nv.Record(*security.NewCondition(true, 55, 0), records, *animal.NewAnimalRecord(animal.Tiger, animal.Walking))
	records = ptz.Record(*security.NewCondition(false, 27, 100), records, *animal.NewAnimalRecord(animal.Rat, animal.Eat))
	records = nv.Record(*security.NewCondition(true, 33, 100), records, *animal.NewAnimalRecord(animal.Rat, animal.Walking))
	records = nv.Record(*security.NewCondition(true, 3, 100), records, *animal.NewAnimalRecord(animal.Bear, animal.Eat))
	records = fc.Record(*security.NewCondition(true, 33, 100), records, *animal.NewAnimalRecord(animal.Rat, animal.Walking))
	records = fc.Record(*security.NewCondition(false, 27, 100), records, *animal.NewAnimalRecord(animal.Rat, animal.Eat))
	records = fc.Record(*security.NewCondition(true, 33, 100), records, *animal.NewAnimalRecord(animal.Rat, animal.Walking))
	records = fc.Record(*security.NewCondition(true, 3, 100), records, *animal.NewAnimalRecord(animal.Bear, animal.Eat))

	for _, v := range records {
		v.Print()
	}

	animals := []animal.AnimalType{animal.Bear, animal.Rat, animal.Tiger}
	actions := []animal.AnimalAction{animal.Eat, animal.Sleep, animal.Walking}
	recorders := []Recorder{ptz, nv}
	rotators := []Rotater{ptz, nv}

	for range 100 {
		isNight := rand.Intn(2) == 1
		distToAnimal := security.Range(rand.Float32() * 20)
		battery := rand.Intn(101)

		animalType := animals[rand.Intn(3)]
		action := actions[rand.Intn(3)]
		for j := range 2 {
			rotators[j].Rotate(security.Angle(rand.Float32() * 180))
			records = recorders[j].Record(*security.NewCondition(isNight, distToAnimal, battery), records, *animal.NewAnimalRecord(animalType, action))
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
	Record(c security.Condition, records []animal.AnimalRecord, rec animal.AnimalRecord) []animal.AnimalRecord
}
