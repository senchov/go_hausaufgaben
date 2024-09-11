package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var datas []Data

	tigerCollar := Collar{Pulse: 25, Temparature: 35}
	subject := tigerCollar.Initialize()
	tigerData := make(chan Data)
	go tigerCollar.CheckSounds(subject, tigerData)
	datas = append(datas, <-tigerData)

	bearCollar := Collar{Pulse: 75, Temparature: 80}
	subject = bearCollar.Initialize()
	bearData := make(chan Data)
	go bearCollar.CheckSounds(subject, bearData)
	datas = append(datas, <-bearData)

	gorillaCollar := Collar{Pulse: 125, Temparature: 110}
	subject = gorillaCollar.Initialize()
	gorillaData := make(chan Data)
	go gorillaCollar.CheckSounds(subject, gorillaData)
	datas = append(datas, <-gorillaData)

	for _, v := range datas {
		switch v.Subgect.(type) {
		case Tiger[float32]:
			tiger := v.Subgect.(Tiger[float32])
			fmt.Printf("The tiger with dexeterity=%f intensity=%d make sound=%s\n", tiger.Dexeterity, v.Intensity, v.Sound)
		case Bear[float32]:
			bear := v.Subgect.(Bear[float32])
			fmt.Printf("The bear with speed=%f intensity=%d make sound=%s\n", bear.Speed, v.Intensity, v.Sound)
		case Gorilla[int8]:
			gorilla := v.Subgect.(Gorilla[int8])
			fmt.Printf("The gorilla with strength=%d intensity=%d make sound=%s\n", gorilla.Strength, v.Intensity, v.Sound)
		}
	}
}

type Animal struct {
}

type Tiger[T any] struct {
	Dexeterity T
}

type Bear[T any] struct {
	Speed T
}

type Gorilla[T any] struct {
	Strength T
}

type Collar struct {
	Pulse       int
	Temparature int
}

type Data struct {
	Subgect   any
	Intensity int
	Sound     string
}

// При ініціалізації нашийник відслідковує пульс і температуру тварини, і з цього робить висновок, що це за тварина.
func (c Collar) Initialize() any {
	switch {
	case c.Pulse < 50 && c.Temparature < 50:
		return Tiger[float32]{Dexeterity: 98}
	case c.Pulse < 100 && c.Temparature < 100:
		return Bear[float32]{Speed: 100}
	case c.Pulse < 150 && c.Temparature < 150:
		return Gorilla[int8]{Strength: 15}
	default:
		return Animal{}
	}
}

// Далі записує інтенсивність дихання і звуків тварини.
func (c Collar) CheckSounds(animal any, data chan Data) {
	intensity := rand.Intn(100)
	sounds := [...]string{"hgg", "op", "aughh"}

	data <- Data{
		Subgect:   animal,
		Intensity: intensity,
		Sound:     sounds[rand.Intn(len(sounds))],
	}
}
