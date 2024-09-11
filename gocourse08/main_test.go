package main

import (
	"reflect"
	"testing"
)

func TestInitialize(t *testing.T) {
	tigerCollar := Collar{Pulse: 25, Temparature: 35}
	subject := tigerCollar.Initialize()
	v, ok := subject.(Tiger[float32])
	if !ok {
		t.Errorf("Expected Tiger but received %s", reflect.TypeOf(v))
	}

	bearCollar := Collar{Pulse: 65, Temparature: 87}
	subject = bearCollar.Initialize()
	v1, ok := subject.(Bear[float32])
	if !ok {
		t.Errorf("Expected Bear but received %s", reflect.TypeOf(v1))
	}

	animalCollar := Collar{Pulse: 1000, Temparature: 33}
	subject = animalCollar.Initialize()
	v2, ok := subject.(Animal)
	if !ok {
		t.Errorf("Expected Animal but received %s", reflect.TypeOf(v2))
	}
}

func TestCheckSounds(t *testing.T) {
	tigerCollar := Collar{Pulse: 25, Temparature: 35}
	subject := tigerCollar.Initialize()
	tigerData := make(chan Data)
	go tigerCollar.CheckSounds(subject, tigerData)
	d := <-tigerData
	t.Logf("Data is %v", d)
	if d.Subgect == nil {
		t.Error("Data subject is nil")
	}
}
