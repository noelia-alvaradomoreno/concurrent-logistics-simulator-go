package utils

import "math/rand"

func GeneratePackageID() int {
	return rand.Intn(1000)
}
