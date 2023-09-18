package stl_test

import (
	"stlReader-DevStorm/stl"
	"strings"
	"testing"
)

func TestCalculateTriangleArea(t *testing.T) {
	testCases := []struct {
		vertices [3][3]float64
		expected float64
	}{
		{[3][3]float64{{0, 0, 0}, {1, 0, 0}, {0, 1, 0}}, 0.5},
		{[3][3]float64{{0, 0, 0}, {1, 0, 0}, {0, 0, 1}}, 0.5},
		// Adicione mais casos de teste conforme necessário
	}

	for _, testCase := range testCases {
		result := stl.CalculateTriangleArea(testCase.vertices)
		if result != testCase.expected {
			t.Errorf("Esperava área %.2f para vértices %v, mas obteve %.2f", testCase.expected, testCase.vertices, result)
		}
	}
}

func TestProcessSTLFile(t *testing.T) {
	// Crie um arquivo STL de teste fictício
	mockSTLContent := `
solid mock
facet normal 0 0 1
outer loop
vertex 0 0 0
vertex 1 0 0
vertex 0 1 0
endloop
endfacet
facet normal 0 0 1
outer loop
vertex 0 0 0
vertex 1 0 0
vertex 0 0 1
endloop
endfacet
endsolid mock
`

	// Crie um leitor de strings para simular um arquivo multipart
	mockFile := strings.NewReader(mockSTLContent)

	// Chame ProcessSTLFile com o arquivo fictício
	area, numTriangles, err := stl.ProcessSTLFile(mockFile)

	// Verifique erros
	if err != nil {
		t.Errorf("ProcessSTLFile retornou um erro: %v", err)
	}

	// Verifique os resultados
	expectedArea := 1.0 // Área total de dois triângulos no conteúdo fictício
	expectedNumTriangles := 2
	if area != expectedArea || numTriangles != expectedNumTriangles {
		t.Errorf("Esperava área=%.2f e numTriangles=%d, mas obteve área=%.2f e numTriangles=%d",
			expectedArea, expectedNumTriangles, area, numTriangles)
	}
}
