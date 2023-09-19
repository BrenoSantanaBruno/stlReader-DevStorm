package main

import (
	"encoding/binary"
	"math/rand"
	"os"
)

func main() {
	// Create an empty binary STL file
	file, err := os.Create("random_triangles.stl")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// Write the empty 80-byte header
	header := make([]byte, 80)
	if _, err := file.Write(header); err != nil {
		panic(err)
	}

	// Write the number of triangles (2 triangles) as a uint32 value in little-endian
	numTriangles := uint32(6)
	if err := binary.Write(file, binary.LittleEndian, numTriangles); err != nil {
		panic(err)
	}

	// Create random triangle data
	for i := 0; i < int(numTriangles); i++ {
		// Create random vertex coordinates (x, y, z)
		vertices := [3][3]float32{
			{rand.Float32(), rand.Float32(), rand.Float32()},
			{rand.Float32(), rand.Float32(), rand.Float32()},
			{rand.Float32(), rand.Float32(), rand.Float32()},
		}

		// Write vertex coordinates (float32) in little-endian
		for _, vertex := range vertices {
			for _, coord := range vertex {
				if err := binary.Write(file, binary.LittleEndian, coord); err != nil {
					panic(err)
				}
			}
		}

		// Write triangle attribute (0) as uint16 in little-endian
		attributeCount := uint16(0)
		if err := binary.Write(file, binary.LittleEndian, attributeCount); err != nil {
			panic(err)
		}
	}

	println("Binary STL file with random triangles and areas created successfully.")
}
