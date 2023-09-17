package stl

import (
	"bufio"
	"fmt"
	"math"
	"mime/multipart"
	"strings"
)

// ProcessSTLFile processes the STL file and returns the total area and the number of triangles.
func ProcessSTLFile(file multipart.File) (float64, int, error) {
	areaTotal := 0.0
	numTriangles := 0

	// Create a scanner to read lines from the input file
	scanner := bufio.NewScanner(file)

	// Iterate through each line of the STL file
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		// Check if the line starts with "facet normal"
		if strings.HasPrefix(line, "facet normal") {
			numTriangles++

			// Skip lines until "outer loop" is encountered
			for scanner.Scan() && !strings.HasPrefix(strings.TrimSpace(scanner.Text()), "outer loop") {
			}

			// Initialize an array to store vertex coordinates
			var vertices [3][3]float64

			// Read and parse each line of vertex coordinates
			for i := 0; i < 3; i++ {
				scanner.Scan()
				parts := strings.Fields(scanner.Text())[1:] // Skip the first field
				for j, part := range parts {
					// Parse and store the vertex coordinates
					coord := strings.Replace(part, "e", "e+", 1)
					fmt.Sscanf(coord, "%f", &vertices[i][j])
				}
			}

			// Calculate the area of the triangle and add it to the total area
			area := CalculateTriangleArea(vertices)
			areaTotal += area
		}
	}

	// Check for scanner errors and return the results
	if err := scanner.Err(); err != nil {
		return 0, 0, err
	}

	return areaTotal, numTriangles, nil
}

// CalculateTriangleArea calculates the area of a triangle based on its vertices.
func CalculateTriangleArea(vertices [3][3]float64) float64 {
	v1 := vertices[0]
	v2 := vertices[1]
	v3 := vertices[2]

	// Calculate the vectors representing the edges of the triangle
	edge1 := [3]float64{v2[0] - v1[0], v2[1] - v1[1], v2[2] - v1[2]}
	edge2 := [3]float64{v3[0] - v1[0], v3[1] - v1[1], v3[2] - v1[2]}

	// Calculate the cross product of the edges to find the normal vector
	normal := [3]float64{
		edge1[1]*edge2[2] - edge1[2]*edge2[1],
		edge1[2]*edge2[0] - edge1[0]*edge2[2],
		edge1[0]*edge2[1] - edge1[1]*edge2[0],
	}

	// Calculate the magnitude of the normal vector
	normalMagnitude := Length(normal)

	// Calculate the area of the triangle using the magnitude of the normal vector
	area := 0.5 * normalMagnitude

	return area
}

// Length calculates the magnitude of a three-dimensional vector.
func Length(v [3]float64) float64 {
	// Calculate the magnitude using the Pythagorean theorem
	return math.Sqrt(v[0]*v[0] + v[1]*v[1] + v[2]*v[2])
}
