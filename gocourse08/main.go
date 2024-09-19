package main

import (
	"fmt"
	"math/rand/v2"
)

func main() {
	var datas []Data

	tigerCollar := Collar{Pulse: 25, Temperature: 35}
	subject := tigerCollar.Initialize()
	tigerData := make(chan Data)
	go tigerCollar.CheckSounds(subject, tigerData)
	datas = append(datas, <-tigerData)

	bearCollar := Collar{Pulse: 75, Temperature: 80}
	subject = bearCollar.Initialize()
	bearData := make(chan Data)
	go bearCollar.CheckSounds(subject, bearData)
	datas = append(datas, <-bearData)

	gorillaCollar := Collar{Pulse: 125, Temperature: 110}
	subject = gorillaCollar.Initialize()
	gorillaData := make(chan Data)
	go gorillaCollar.CheckSounds(subject, gorillaData)
	datas = append(datas, <-gorillaData)

	for _, v := range datas {
		switch v.Subject.(type) {
		case Tiger[float32]:
			tiger := v.Subject.(Tiger[float32])
			fmt.Printf("The tiger with dexeterity=%f intensity=%d make sound=%s\n", tiger.Dexterity, v.Intensity, v.Sound)
		case Bear[float32]:
			bear := v.Subject.(Bear[float32])
			fmt.Printf("The bear with speed=%f intensity=%d make sound=%s\n", bear.Speed, v.Intensity, v.Sound)
		case Gorilla[int8]:
			gorilla := v.Subject.(Gorilla[int8])
			fmt.Printf("The gorilla with strength=%d intensity=%d make sound=%s\n", gorilla.Strength, v.Intensity, v.Sound)
		}
	}
}

type Animal struct {
}

type Tiger[T any] struct {
	Dexterity T
}

type Bear[T float32 | float64 | int8 | int32] struct {
	Speed T
}

type Gorilla[T int8 | int32 | int | int64] struct {
	Strength T
}

type Collar struct {
	Pulse       int
	Temperature int
}

type Data struct {
	Subject   any
	Intensity int
	Sound     string
}

// При ініціалізації нашийник відслідковує пульс і температуру тварини, і з цього робить висновок, що це за тварина.
func (c Collar) Initialize() any {
	switch {
	case c.Pulse < 50 && c.Temperature < 50:
		return Tiger[float32]{Dexterity: 98}
	case c.Pulse < 100 && c.Temperature < 100:
		return Bear[float32]{Speed: 100}
	case c.Pulse < 150 && c.Temperature < 150:
		return Gorilla[int8]{Strength: 15}
	default:
		return Animal{}
	}
}

// Далі записує інтенсивність дихання і звуків тварини.
func (c Collar) CheckSounds(animal any, data chan Data) {
	intensity := rand.Int32N(100)
	sounds := [...]string{"hgg", "op", "aughh"}
	index := int(rand.Int32N(int32(len(sounds))))

	data <- Data{
		Subject:   animal,
		Intensity: int(intensity),
		Sound:     sounds[index],
	}
}
