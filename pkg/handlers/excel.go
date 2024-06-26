package handlers

import (
	"log"

	"github.com/xuri/excelize/v2"
)

const filePath = "C:\\Users\\vicev\\Documents\\projects\\enubeTeste\\pkg\\data\\alunos.xlsx"
const sheetName = "Sheet1"

const truefilePath = "C:\\Users\\vicev\\Documents\\projects\\enubeTeste\\pkg\\data\\Reconfile fornecedores.xlsx"
const trueSheetName = "Planilha1"

func readColumn(column string) ([]string, error) {
	f, err := excelize.OpenFile(filePath)
	if err != nil {
		return nil, err
	}
	defer func() {
		if err := f.Close(); err != nil {
			log.Println("Erro ao fechar o arquivo:", err)
		}
	}()

	rows, err := f.GetRows(sheetName)
	if err != nil {
		return nil, err
	}

	var values []string
	var colIndex int
	switch column {
	case "Nome":
		colIndex = 1
	case "Idade":
		colIndex = 2
	case "Escola":
		colIndex = 3
	default:
		return nil, nil // Tratar caso a coluna não seja encontrada
	}

	for _, row := range rows {
		if len(row) >= colIndex {
			values = append(values, row[colIndex-1])
		}
	}
	return values, nil
}

func readTrueColumn(column string) ([]string, error) {
	f, err := excelize.OpenFile(truefilePath)
	if err != nil {
		return nil, err
	}
	defer func() {
		if err := f.Close(); err != nil {
			log.Println("Erro ao fechar o arquivo:", err)
		}
	}()

	rows, err := f.GetRows(trueSheetName)
	if err != nil {
		return nil, err
	}

	var values []string
	var colIndex int
	switch column {
	case "PartnerId":
		colIndex = 1
	case "PartnerName":
		colIndex = 2
	case "CustomerId":
		colIndex = 3
	default:
		return nil, nil // Tratar caso a coluna não seja encontrada
	}

	for _, row := range rows {
		if len(row) >= colIndex {
			values = append(values, row[colIndex-1])
		}
	}
	return values, nil
}
