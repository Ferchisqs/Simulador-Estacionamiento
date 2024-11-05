package views

import (
	"SimuladorSemaforo/scenes"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/theme"
)

type MainView struct{}

func NewMainView() *MainView {
	return &MainView{}
}

func (v *MainView) Run() {
	myApp := app.New()
	myApp.Settings().SetTheme(theme.LightTheme())
	window := myApp.NewWindow("Simulador de Estacionamiento")
	window.Resize(fyne.NewSize(600, 700))
	window.CenterOnScreen()

	mainScene := scenes.NewMainScene(window)
	mainScene.Show()
	go mainScene.Run()
	window.ShowAndRun()
}
