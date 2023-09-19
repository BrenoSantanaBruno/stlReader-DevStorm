package stl

import (
	"bufio"
	"encoding/binary"
	"fmt"
	"io"
	"math"
	"strconv"
	"strings"
)

// Triangle represents a triangle in an STL file.
type Triangle struct {
	Vertices [3][3]float64
}

// ProcessSTLFile processes the STL file and returns the total area and the number of triangles.
func ProcessSTLFile(file io.Reader) (float64, int, error) {
	isBinary, err := IsBinarySTL(file)
	if err != nil {
		return 0, 0, err
	}

	if isBinary {
		return ProcessBinarySTL(file)
	}

	return ProcessASCIISTL(file)
}

// ProcessBinarySTL processes a binary STL file and returns the total area and the number of triangles.
func ProcessBinarySTL(file io.Reader) (float64, int, error) {
	// Read the 80-byte header (which can be ignored)
	header := make([]byte, 80)
	_, err := io.ReadFull(file, header)
	if err != nil {
		return 0, 0, err
	}

	// Read the 4-byte triangle count (little-endian)
	var numTriangles uint32
	err = binary.Read(file, binary.LittleEndian, &numTriangles)
	if err != nil {
		return 0, 0, err
	}

	// Iterate through each triangle in the binary STL file
	areaTotal := 0.0
	for i := uint32(0); i < numTriangles; i++ {
		// Read the 12 floats (3 vertices, each with X, Y, and Z) for each triangle
		var vertices [3][3]float64
		for j := 0; j < 3; j++ {
			for k := 0; k < 3; k++ {
				var coord float32
				err := binary.Read(file, binary.LittleEndian, &coord)
				if err != nil {
					return 0, 0, err
				}
				// Convert from float32 to float64
				vertices[j][k] = float64(coord)
			}
		}

		// Calculate the area of the triangle and add it to the total area
		area := CalculateTriangleArea(vertices)
		areaTotal += area

		// Skip the 2-byte attribute count
		var attributeCount uint16
		err := binary.Read(file, binary.LittleEndian, &attributeCount)
		if err != nil {
			return 0, 0, err
		}
	}

	return areaTotal, int(numTriangles), nil
}

// ProcessASCIISTL processes an ASCII STL file and returns the total area and the number of triangles.
func ProcessASCIISTL(file io.Reader) (float64, int, error) {
	areaTotal := 0.0
	numTriangles := 0

	// Create a scanner to read lines from the input file
	scanner := bufio.NewScanner(file)

	var currentTriangle Triangle

	for scanner.Scan() {
		line := scanner.Text()
		line = trimSpaceAndToLower(line)

		if strings.HasPrefix(line, "facet normal") {
			numTriangles++
		} else if strings.HasPrefix(line, "outer loop") {
			// Read and parse each line of vertex coordinates
			for i := 0; i < 3; i++ {
				if !scanner.Scan() {
					return 0, 0, fmt.Errorf("unexpected end of file")
				}

				line := scanner.Text()
				line = trimSpaceAndToLower(line)
				parts := strings.Fields(line)[1:] // Skip the first field
				if len(parts) != 3 {
					return 0, 0, fmt.Errorf("invalid vertex coordinates: %s", line)
				}

				for j, part := range parts {
					// Parse and store the vertex coordinates
					coord, err := strconv.ParseFloat(strings.Replace(part, "e", "e+", 1), 64)
					if err != nil {
						return 0, 0, err
					}
					currentTriangle.Vertices[i][j] = coord
				}
			}

			// Calculate the area of the triangle and add it to the total area
			area := CalculateTriangleArea(currentTriangle.Vertices)
			areaTotal += area
		}
	}

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

// Helper function to trim spaces and convert to lowercase
func trimSpaceAndToLower(s string) string {
	return strings.ToLower(strings.TrimSpace(s))
}

// IsBinarySTL checks if a given STL file is in binary format.
func IsBinarySTL(file io.Reader) (bool, error) {
	// Read the first 5 bytes from the file
	header := make([]byte, 5)
	_, err := io.ReadAtLeast(file, header, 5)
	if err != nil {
		return false, err
	}

	// Check if the first 5 bytes contain "solid" keyword (ASCII STL) or a binary header
	if string(header) == "solid" {
		return false, nil // ASCII STL
	}

	// Check if the first 80 bytes contain an ASCII comment (text after "solid")
	comment := make([]byte, 75) // 80 bytes minus "solid"
	_, err = io.ReadAtLeast(file, comment, 75)
	if err != nil {
		return false, err
	}

	// Check if the comment section is empty or contains only whitespaces
	for _, char := range comment {
		if char != ' ' && char != '\t' && char != '\n' && char != '\r' {
			return true, nil // Binary STL
		}
	}

	return false, nil // ASCII STL (no comment)
}
