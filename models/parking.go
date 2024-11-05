package models

import "sync"

type Parking struct {
	spaces     chan int
	doorAccess chan struct{}
	lock       *sync.Mutex
	observers  []Observer
}

func NewParking(spaces chan int, doorAccess chan struct{}, lock *sync.Mutex) *Parking {
	return &Parking{
		spaces:    spaces,
		doorAccess: doorAccess,
		lock:      lock,
		observers: []Observer{},
	}
}

func (p *Parking) GetSpaces() chan int {
	return p.spaces
}

func (p *Parking) Attach(observer Observer) {
	p.lock.Lock()
	defer p.lock.Unlock()
	p.observers = append(p.observers, observer)
}

func (p *Parking) Detach(observer Observer) {
	p.lock.Lock()
	defer p.lock.Unlock()
	for i, o := range p.observers {
		if o == observer {
			p.observers = append(p.observers[:i], p.observers[i+1:]...)
			break
		}
	}
}

// NotifyObservers notifies the specific observer about a status update.
func (p *Parking) NotifyObservers(carID int, status string) {
	p.lock.Lock()
	defer p.lock.Unlock()
	for _, observer := range p.observers {
		if observer.GetID() == carID {
			observer.Update(status)
		}
	}
}
