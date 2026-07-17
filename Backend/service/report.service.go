package service

import (
	"bytes"
	"fmt"

	"github.com/luthfi/pos/repository"
	"github.com/xuri/excelize/v2"
)

type ReportService struct {
	repo repository.TransactionRepository
}

func NewReportService(r repository.TransactionRepository) *ReportService {
	return &ReportService{repo: r}
}

func (s *ReportService) ExportToExcel() (*bytes.Buffer, error) {
	transactions, err := s.repo.GetAll()
	if err != nil {
		return nil, err
	}

	f := excelize.NewFile()
	sheet := "Data Penjualan"
	f.SetSheetName("Sheet1", sheet)

	// Header Style
	f.SetCellValue(sheet, "A1", "ID")
	f.SetCellValue(sheet, "B1", "Customer Name")
	f.SetCellValue(sheet, "C1", "Total Amount")
	f.SetCellValue(sheet, "D1", "Discount")
	f.SetCellValue(sheet, "E1", "Grand Total")
	f.SetCellValue(sheet, "F1", "Status")

	// Isi Data
	for i, t := range transactions {
		row := i + 2
		f.SetCellValue(sheet, fmt.Sprintf("A%d", row), t.ID)
		f.SetCellValue(sheet, fmt.Sprintf("B%d", row), t.CustomerName)
		f.SetCellValue(sheet, fmt.Sprintf("C%d", row), t.TotalAmount)
		f.SetCellValue(sheet, fmt.Sprintf("D%d", row), t.Discount)
		f.SetCellValue(sheet, fmt.Sprintf("E%d", row), t.GrandTotal)
		f.SetCellValue(sheet, fmt.Sprintf("F%d", row), t.Status)
	}

	return f.WriteToBuffer()
}
