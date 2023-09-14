package main

import (
	"fmt"
	"stlReader-DevStorm/stlreader"
)

func main() {
	fileName := "/Users/brenobruno/Desktop/Moon.stl" // TODO: Alterar para o caminho do arquivo STL desejado.
	totalArea, numTriangles, err := stlreader.CalculateSTLFileArea(fileName)
	if err != nil {
		fmt.Println("Error calculating the area of the STL file:", err)
		return
	}

	fmt.Printf("Number of Triangles: %d\n", numTriangles)
	fmt.Printf("Total area of triangles in the STL file: %.4f\n", totalArea)
}
