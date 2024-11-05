package models

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"image/color"
	"math/rand"
	"time"
)

type Car struct {
	id          int
	parkingTime time.Duration
	image       *canvas.Rectangle
}

func NewCar(id int) *Car {
	return &Car{
		id:          id,
		parkingTime: time.Duration(rand.Intn(5)+3) * time.Second,
		image:       canvas.NewRectangle(color.RGBA{uint8(rand.Intn(255)), uint8(rand.Intn(255)), uint8(rand.Intn(255)), 255}),
	}
}

// Update only handles notifications specific to this car
func (c *Car) Update(status string) {
	fmt.Printf("Car %d recibió notificación: %s\n", c.id, status)
}

// GetID returns the unique ID of the car
func (c *Car) GetID() int {
	return c.id
}

func (c *Car) Enter(parking *Parking, carsContainer *fyne.Container) {
	parking.doorAccess <- struct{}{}
	fmt.Printf("Car %d ha ingresado\n", c.id)
	parking.NotifyObservers(c.id, fmt.Sprintf("Car %d ha ingresado", c.id))

	for x := 0; x < 200; x += 20 {
		c.image.Move(fyne.NewPos(float32(20+x), 75))
		carsContainer.Refresh()
		time.Sleep(100 * time.Millisecond)
	}
	parking.spaces <- c.id
	<-parking.doorAccess
}

func (c *Car) Park(parking *Parking, carsContainer *fyne.Container) {
	c.Enter(parking, carsContainer)

	for y := 0; y < 200; y += 20 {
		c.image.Move(fyne.NewPos(120, float32(120+y)))
		carsContainer.Refresh()
		time.Sleep(100 * time.Millisecond)
	}
	time.Sleep(c.parkingTime)
	c.Leave(parking, carsContainer)
}

func (c *Car) Leave(parking *Parking, carsContainer *fyne.Container) {
	<-parking.spaces
	parking.doorAccess <- struct{}{}
	fmt.Printf("Car %d ha salido\n", c.id)
	parking.NotifyObservers(c.id, fmt.Sprintf("Car %d ha salido", c.id))

	for x := 0; x < 300; x += 20 {
		c.image.Move(fyne.NewPos(float32(400-x), 620))
		carsContainer.Refresh()
		time.Sleep(100 * time.Millisecond)
	}
	carsContainer.Remove(c.image)
	carsContainer.Refresh()
	<-parking.doorAccess
}

func (c *Car) GetCarImage() *canvas.Rectangle {
	return c.image
}
