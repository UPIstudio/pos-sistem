package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/luthfi/pos/service"
)

type ReportHandler struct{ svc *service.ReportService }

func NewReportHandler(s *service.ReportService) *ReportHandler { return &ReportHandler{svc: s} }

func (h *ReportHandler) ExportExcel(c *fiber.Ctx) error {
	buffer, err := h.svc.ExportToExcel()
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Gagal generate excel"})
	}

	c.Set("Content-Disposition", "attachment; filename=Laporan_Penjualan.xlsx")
	c.Set("Content-Type", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")

	return c.Send(buffer.Bytes())
}
