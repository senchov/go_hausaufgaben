package main

import (
	"fmt"
	"log"

	"github.com/senchov/go_hausaufgaben/gocourse10/feed"
	"github.com/senchov/go_hausaufgaben/gocourse10/observer"
)

func main() {
	log.Println("Start")
	// strategy
	collar := feed.Collar{Trough: feed.MeetTrought{}, ID: 1}
	feed.Feed(collar)
	collar.SetTrough(&feed.FishTrough{})
	feed.Feed(collar)
	collar.SetTrough(&feed.VegetablesTrought{})
	feed.Feed(collar)

	// decorator
	collar.SetTrough(&feed.SausagesTrough{Trough: feed.MeetTrought{}})
	feed.Feed(collar)

	item := observer.Item{}
	item.Register(collar)

	for i := range 24 {
		if i == 9 {
			fmt.Println("Go to breakfast")
			item.NotifyAll()
			collar.SetTrough(&feed.VegetablesTrought{})
		}

		if i == 13 {
			fmt.Println("Go to the lunch")
			item.NotifyAll()
			collar.SetTrough(&feed.FishTrough{})
		}

		if i == 18 {
			fmt.Println("Go to the dinner")
			item.NotifyAll()
			collar.SetTrough(&feed.MeetTrought{})
		}
	}
}
