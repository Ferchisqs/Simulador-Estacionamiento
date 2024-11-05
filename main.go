package main

import (
	"fmt"
	"SimuladorSemaforo/views"
)

func main() {
	fmt.Println("Iniciando la aplicación de estacionamiento...")

	mainView := views.NewMainView()
	fmt.Println("Vista principal creada.")

	mainView.Run()
	fmt.Println("Aplicación en ejecución.")
}
