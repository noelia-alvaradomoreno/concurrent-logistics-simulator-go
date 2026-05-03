package main

import (
	"fmt"
	"logistic/dispatcher"
	"logistic/packages"
	"logistic/sorter"
)

func main() {
	//Datos de entrada
	packageList := []*packages.Package{
		{PackageID: 1, SenderName: "Alice", Destination: "North", Status: "Recibido"},
		{PackageID: 2, SenderName: "Marian", Destination: "South", Status: "Recibido"},
		{PackageID: 3, SenderName: "Luis", Destination: "East", Status: "Recibido"},
		{PackageID: 4, SenderName: "Ricardo", Destination: "Costa Rica", Status: "Recibido"},
	}

	//Canales para el flujo de datos (Pipeline)
	sortChannel := make(chan *packages.Package)
	dispatchChannel := make(chan *packages.Package)
	done := make(chan bool)

	// Orquestación del Pipeline

	// Goroutine de Clasificación: Delega la lógica al paquete 'sorter'
	go func() {
		sorter.ProcessPackages(packageList, sortChannel)
		close(sortChannel)
	}()

	// Goroutine de Transferencia: Mueve paquetes entre etapas de forma segura
	go func() {
		for pkg := range sortChannel {
			fmt.Printf("Transferring Package %d to dispatch queue\n", pkg.PackageID)
			dispatchChannel <- pkg
		}
		close(dispatchChannel)
	}()

	// Goroutine de Despacho: Delega la acción final al paquete 'dispatcher'
	go func() {
		for pkg := range dispatchChannel {
			dispatcher.DispatchPackage(pkg, pkg.Destination)
		}
		done <- true
	}()

	// Espera final
	<-done
	fmt.Println("Logística completada: Todos los paquetes fueron procesados.")
}
