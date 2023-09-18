package mocks

import (
	"bytes"
	"testing"
)

func TestMockFileRead(t *testing.T) {
	// Create an instance of MockFile with test data
	data := []byte("Test data")
	mockFile := &MockFile{Reader: bytes.NewReader(data)}

	// Create a buffer to store the read output
	buffer := make([]byte, len(data))

	// Read data from MockFile into the buffer
	n, err := mockFile.Read(buffer)

	// Check for no errors
	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}

	// Check if the number of bytes read is correct
	if n != len(data) {
		t.Errorf("Expected to read %d bytes, but got %d", len(data), n)
	}

	// Check if the read data matches the original data
	if !bytes.Equal(buffer, data) {
		t.Errorf("Expected data %v, but got %v", data, buffer)
	}
}

func TestMockFileClose(t *testing.T) {
	// Create an instance of MockFile
	mockFile := &MockFile{Reader: nil}

	// Close the MockFile
	err := mockFile.Close()

	// Check for no errors
	if err != nil {
		t.Errorf("Expected no error when closing MockFile, but got %v", err)
	}
}
