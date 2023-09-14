package stlreader_test

import (
	"stlReader-DevStorm/stlreader"
	"testing"
)

func TestCalculateSTLFileArea(t *testing.T) {
	fileName := "/Users/brenobruno/Desktop/Moon.stl" // TODO: Alterar para o caminho do arquivo STL desejado.
	_, _, err := stlreader.CalculateSTLFileArea(fileName)
	if err != nil {
		t.Errorf("Error calculating the area of the STL file: %v", err)
		return
	}

	//if numTriangles != 2304 {
	//	t.Errorf("Expected 2304 triangles, got %d", numTriangles)
	//}

	//if totalArea != 0.000000 {
	//	t.Errorf("Expected 0.000000 area, got %.6f", totalArea)
	//}
}
