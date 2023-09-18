package mocks

import (
	"io"
)

// MockFile is a structure that implements the multipart.File interface
type MockFile struct {
	io.Reader
}

// Close closes the MockFile (in this case, it does nothing)
func (mf *MockFile) Close() error {
	return nil
}
