package sorter

import (
	"fmt"
	"logistic/config"
	"logistic/packages"
	"logistic/utils"
)

// ProcessPackages es la función que el main llamará.
// Recibe la lista completa y decide qué hacer con cada uno.
func ProcessPackages(pkgs []*packages.Package, sortChannel chan<- *packages.Package) {
	for _, pkg := range pkgs {
		if isValidZone(pkg.Destination) {
			// Si es válido, lo procesamos.
			sortAndSend(pkg, sortChannel)
		} else {
			// Si no, informamos el error.
			fmt.Printf("Skipping package from %s: Invalid destination '%s'\n", pkg.SenderName, pkg.Destination)
		}
	}
}

func sortAndSend(pkg *packages.Package, sortChannel chan<- *packages.Package) {
	pkg.PackageID = utils.GeneratePackageID()
	pkg.Status = "Ordenado"

	fmt.Printf("Package %d (from %s) sorted to %s. Status: %s\n",
		pkg.PackageID, pkg.SenderName, pkg.Destination, pkg.Status)

	// Enviamos al canal
	sortChannel <- pkg
}
func isValidZone(dest string) bool {
	for _, zone := range config.Zones {
		if zone == dest {
			return true
		}
	}
	return false
}
