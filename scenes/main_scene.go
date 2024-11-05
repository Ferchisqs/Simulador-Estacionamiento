package scenes

import (
	"SimuladorSemaforo/models"
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"image/color"
	"sync"
	"time"
)

type MainScene struct {
	window fyne.Window
}

func NewMainScene(window fyne.Window) *MainScene {
	return &MainScene{
		window: window,
	}
}

var carsContainer = container.NewWithoutLayout()

func (s *MainScene) Show() {
	background := canvas.NewRectangle(color.RGBA{235, 235, 235, 255})
	background.Resize(fyne.NewSize(600, 700))
	carsContainer.Add(background)

	entryArea := canvas.NewRectangle(color.RGBA{180, 180, 255, 255})
	entryArea.Resize(fyne.NewSize(150, 50))
	entryArea.Move(fyne.NewPos(20, 50))
	carsContainer.Add(entryArea)

	parkingArea := canvas.NewRectangle(color.RGBA{200, 200, 200, 255})
	parkingArea.Resize(fyne.NewSize(400, 500))
	parkingArea.Move(fyne.NewPos(20, 120))
	carsContainer.Add(parkingArea)

	exitArea := canvas.NewRectangle(color.RGBA{255, 180, 180, 255})
	exitArea.Resize(fyne.NewSize(150, 50))
	exitArea.Move(fyne.NewPos(450, 50))
	carsContainer.Add(exitArea)

	s.window.SetContent(carsContainer)
}

func (s *MainScene) Run() {
	parking := models.NewParking(make(chan int, 20), make(chan struct{}, 1), &sync.Mutex{})
	var wg sync.WaitGroup
	poisson := models.NewPoissonDist()

	for i := 0; i < 100; i++ {
		wg.Add(1)
		car := models.NewCar(i)
		parking.Attach(car)

		go s.runCar(car, parking, &wg)

		arrivalInterval := time.Duration(poisson.Generate(3) * float64(time.Second))
		time.Sleep(arrivalInterval)
	}

	wg.Wait()
	fmt.Println("Simulación terminada")
}

func (s *MainScene) runCar(car *models.Car, parking *models.Parking, wg *sync.WaitGroup) {
	defer wg.Done()
	carImage := car.GetCarImage()
	carImage.Resize(fyne.NewSize(40, 25))
	carImage.Move(fyne.NewPos(-50, 75))

	carsContainer.Add(carImage)
	carsContainer.Refresh()

	car.Park(parking, carsContainer)
	parking.Detach(car)
}
