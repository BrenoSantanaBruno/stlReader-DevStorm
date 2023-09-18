package mocks

import (
	"bytes"
	"testing"
)

func TestMockFileRead(t *testing.T) {
	// Crie uma instância de MockFile com dados de teste
	data := []byte("Test data")
	mockFile := &MockFile{Reader: bytes.NewReader(data)}

	// Crie um buffer para armazenar a saída lida
	buffer := make([]byte, len(data))

	// Leia os dados do MockFile para o buffer
	n, err := mockFile.Read(buffer)

	// Verifique se não há erros
	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}

	// Verifique se o número de bytes lidos está correto
	if n != len(data) {
		t.Errorf("Expected to read %d bytes, but got %d", len(data), n)
	}

	// Verifique se os dados lidos correspondem aos dados originais
	if !bytes.Equal(buffer, data) {
		t.Errorf("Expected data %v, but got %v", data, buffer)
	}
}

func TestMockFileClose(t *testing.T) {
	// Crie uma instância de MockFile
	mockFile := &MockFile{Reader: nil}

	// Feche o MockFile
	err := mockFile.Close()

	// Verifique se não há erros
	if err != nil {
		t.Errorf("Expected no error when closing MockFile, but got %v", err)
	}
}
