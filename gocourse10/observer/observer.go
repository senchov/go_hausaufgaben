package observer

type Observer interface {
	TimeToEat()
	GetID() int
}

type Item struct {
	ObserverList []Observer
}

func (i *Item) Register(o Observer) {
	i.ObserverList = append(i.ObserverList, o)
}

func (i *Item) Deregister(o Observer) {
	i.ObserverList = removeFromslice(i.ObserverList, o)
}

func (i *Item) NotifyAll() {
	for _, observer := range i.ObserverList {
		observer.TimeToEat()
	}
}

func removeFromslice(observerList []Observer, observerToRemove Observer) []Observer {
	observerListLength := len(observerList)
	for i, observer := range observerList {
		if observerToRemove.GetID() == observer.GetID() {
			observerList[observerListLength-1], observerList[i] = observerList[i], observerList[observerListLength-1]
			return observerList[:observerListLength-1]
		}
	}
	return observerList
}
