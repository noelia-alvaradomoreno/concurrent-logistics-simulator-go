package dispatcher

import (
	"fmt"
	"logistic/packages" 
)


func DispatchPackage(pkg *packages.Package, zone string) {
	// Cambia el estado del paquete a despachado
	pkg.Status = "Despachado"

	
	fmt.Printf("Dispatching Package %d to Zone %s. Status: %s\n", pkg.PackageID, zone, pkg.Status)
	
}
