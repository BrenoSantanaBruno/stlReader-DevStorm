package stl

import (
	"bufio"
	"fmt"
	"math"
	"mime/multipart"
	"strings"
)

// ProcessSTLFile processa o arquivo STL e retorna a área total e o número de triângulos.
func ProcessSTLFile(file multipart.File) (float64, int, error) {
	areaTotal := 0.0
	numTriangles := 0

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		if strings.HasPrefix(line, "facet normal") {
			numTriangles++
			for scanner.Scan() && !strings.HasPrefix(strings.TrimSpace(scanner.Text()), "outer loop") {
			}

			var vertices [3][3]float64
			for i := 0; i < 3; i++ {
				scanner.Scan()
				parts := strings.Fields(scanner.Text())[1:]
				for j, part := range parts {
					coord := strings.Replace(part, "e", "e+", 1)
					fmt.Sscanf(coord, "%f", &vertices[i][j])
				}
			}

			area := CalculateTriangleArea(vertices)
			areaTotal += area
		}
	}

	if err := scanner.Err(); err != nil {
		return 0, 0, err
	}

	return areaTotal, numTriangles, nil
}

// CalculateTriangleArea calcula a área de um triângulo com base em seus vértices.
func CalculateTriangleArea(vertices [3][3]float64) float64 {
	v1 := vertices[0]
	v2 := vertices[1]
	v3 := vertices[2]

	edge1 := [3]float64{v2[0] - v1[0], v2[1] - v1[1], v2[2] - v1[2]}
	edge2 := [3]float64{v3[0] - v1[0], v3[1] - v1[1], v3[2] - v1[2]}

	normal := [3]float64{
		edge1[1]*edge2[2] - edge1[2]*edge2[1],
		edge1[2]*edge2[0] - edge1[0]*edge2[2],
		edge1[0]*edge2[1] - edge1[1]*edge2[0],
	}

	normalMagnitude := Length(normal)

	area := 0.5 * normalMagnitude

	return area
}

// Length calcula a magnitude de um vetor tridimensional.
func Length(v [3]float64) float64 {
	return math.Sqrt(v[0]*v[0] + v[1]*v[1] + v[2]*v[2])
}