package security

import "testing"

func TestRotate(t *testing.T) {
	ptz := NewPTZCamera(0, 0)
	ptz.Rotate(30)
	if ptz.Angle < 30 {
		t.Error("PTZ camera dosen't rotate")
	}

	fixed := NewFixedCamera(0, 0)
	fixed.Rotate(90)
	if fixed.Angle != 0 {
		t.Error("Fixed camera shouldnt rotate")
	}
}

func TestRecord(t *testing.T) {
	ptz := NewPTZCamera(10, 0)
	records := []AnimalRecord{}
	ptz.Record(*NewConditions(false, 5, 100), &records, *NewAnimalRecord(Bear, Sleep))
	if len(records) == 0 {
		t.Error("Doesen't record anything")
	}

	ptz.Record(*NewConditions(true, 5, 100), &records, *NewAnimalRecord(Bear, Sleep))
	if len(records) == 2 {
		t.Error("Shouldn't record at night")
	}

	nv := NewNightvisionCamera(30, 6)
	nv.Record(*NewConditions(true, 5, 100), &records, *NewAnimalRecord(Tiger, Walking))
	if len(records) != 2 {
		t.Error("Nightvision camera doesn't work")
	}
}
