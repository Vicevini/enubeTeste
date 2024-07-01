package main

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"
)

func TestFilePath(t *testing.T) {
	basepath, err := os.Getwd()
	if err != nil {
		t.Fatalf("Erro ao obter o diretório atual: %v", err)
	}

	relativePath := "../pkg/data/Reconfile fornecedores.xlsx"
	absPath := filepath.Join(basepath, relativePath)
	fmt.Println("Caminho Absoluto:", absPath)

	if _, err := os.Stat(absPath); os.IsNotExist(err) {
		t.Fatalf("Arquivo não encontrado: %s", absPath)
	} else {
		fmt.Println("Arquivo encontrado:", absPath)
	}
}
