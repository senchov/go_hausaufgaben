package security

import (
	"gocourse05/animal"
	"testing"
)

func TestPTZCamera_Rotate(t *testing.T) {
	ptz := NewPTZCamera(0, 0)
	ptz.Rotate(30)
	if ptz.Angle < 30 {
		t.Error("PTZ camera doesn't rotate to angle 30 should be 30 degrees")
	}
}

func TestFixedCamera_Record(t *testing.T) {
	records := []animal.AnimalRecord{}

	nv := NewNightvisionCamera(30, 6)
	records = nv.Record(*NewCondition(true, 5, 100), records, *animal.NewAnimalRecord(animal.Tiger, animal.Walking))
	if len(records) != 1 {
		t.Error("Nightvision camera doesn't work")
	}
}

func TestPTZCamera_Recordt(t *testing.T) {
	ptz := NewPTZCamera(10, 0)
	records := []animal.AnimalRecord{}
	records = ptz.Record(*NewCondition(false, 5, 100), records, *animal.NewAnimalRecord(animal.Bear, animal.Sleep))
	if len(records) == 0 {
		t.Error("Amount of records equals zero but should one")
	}

	records = ptz.Record(*NewCondition(true, 5, 100), records, *animal.NewAnimalRecord(animal.Bear, animal.Sleep))
	if len(records) == 2 {
		t.Error("Amount of records was two should be one, because the camera shouldn't record at night")
	}
}
