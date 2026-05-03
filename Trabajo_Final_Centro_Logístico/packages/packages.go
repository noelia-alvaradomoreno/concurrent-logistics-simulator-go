package packages

type Package struct {
	PackageID   int
	SenderName  string
	Destination string
	Status      string // Recibido, ordenado, despachado
}
