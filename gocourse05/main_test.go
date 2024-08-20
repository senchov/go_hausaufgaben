package main

import "testing"

func TestRotate(t *testing.T) {
	ptz := PTZCamera{Angle: 0, Range: 0}
	ptz.Rotate(30)
	if ptz.Angle < 30 {
		t.Error("PTZ camera dosen't rotate")
	}

	fixed := FixedCamera{Angle: 0, Range: 0}
	fixed.Rotate(90)
	if fixed.Angle != 0 {
		t.Error("Fixed camera shouldnt rotate")
	}
}

func TestRecord(t *testing.T) {
	ptz := PTZCamera{Angle: 0, Range: 10}
	records := []AnimalRecord{}
	ptz.Record(Conditions{IsNight: false, DistToAnimal: 5, Battery: 100}, &records, AnimalRecord{Bear, Sleep})
	if len(records) == 0 {
		t.Error("Doesen't record anything")
	}

	ptz.Record(Conditions{IsNight: true, DistToAnimal: 5, Battery: 100}, &records, AnimalRecord{Bear, Sleep})
	if len(records) == 2 {
		t.Error("Shouldn't record at night")
	}

	nv := NightvisionCamera{Angle: 30, Range: 6}
	nv.Record(Conditions{IsNight: true, DistToAnimal: 5, Battery: 100}, &records, AnimalRecord{Tiger, Walking})
	if len(records) != 2 {
		t.Error("Nightvision camera doesn't work")
	}
}
