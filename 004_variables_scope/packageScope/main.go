package main


import (
	"fmt"
	"github.com/KapaMiki/GolangLearning/004_variables_scope/packageScope/nestedPackage"
)

func main() {
	fmt.Println(nestedPackage.airplaneName)
	fmt.Println(nestedPackage.Airplane)
}