package excel

import (
	"fmt"
	"log"

	"github.com/xuri/excelize/v2"
	"moritzgoedicke.com/bewaesserung/models"
)

type ExcelFile struct {
	file *excelize.File
}

func NewExcelFile() *ExcelFile {
	f := excelize.NewFile()
	_, err := f.NewSheet("Sheet1")
	if err != nil {
		log.Fatalf("Failed to create new sheet: %v", err)
	}
	return &ExcelFile{file: f}
}

func (e *ExcelFile) SetHeaders(headers []string) {
	for i, header := range headers {
		cell := fmt.Sprintf("%s1", string('A'+i))
		e.file.SetCellValue("Sheet1", cell, header)
	}
}

func (e *ExcelFile) WriteRow(row int, data models.DeviceData, payload models.DecodedPayload) {
	e.file.SetCellValue("Sheet1", fmt.Sprintf("A%d", row), data.Identifiers[0].DeviceIds.DeviceID)
	e.file.SetCellValue("Sheet1", fmt.Sprintf("B%d", row), data.Identifiers[0].DeviceIds.DeviceEUI)
	e.file.SetCellValue("Sheet1", fmt.Sprintf("C%d", row), data.Data.ReceivedAt)
	e.file.SetCellValue("Sheet1", fmt.Sprintf("D%d", row), payload.BAT)
	e.file.SetCellValue("Sheet1", fmt.Sprintf("E%d", row), payload.H1)
	e.file.SetCellValue("Sheet1", fmt.Sprintf("F%d", row), payload.H2)
	e.file.SetCellValue("Sheet1", fmt.Sprintf("G%d", row), payload.T1)
}

func (e *ExcelFile) Save(outputFile string) {
	if err := e.file.SaveAs(outputFile); err != nil {
		log.Fatalf("Failed to save Excel file: %v", err)
	}
}
