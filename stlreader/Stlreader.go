package stlreader

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"
)

// CalculateSTLFileArea reads an STL file and returns the total area and number of triangles.
func CalculateSTLFileArea(fileName string) (float64, int, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return 0, 0, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	totalArea := 0.0
	numTriangles := 0

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
			totalArea += area
		}
	}

	if err := scanner.Err(); err != nil {
		return 0, 0, err
	}

	return totalArea, numTriangles, nil
}

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

	normalMagnitude := math.Sqrt(normal[0]*normal[0] + normal[1]*normal[1] + normal[2]*normal[2])

	area := 0.5 * normalMagnitude

	return area
}
