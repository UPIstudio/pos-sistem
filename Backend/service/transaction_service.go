package service

import (
	"errors"

	"github.com/luthfi/pos/models"
	"github.com/luthfi/pos/repository"
)

type TransactionService struct {
	repo repository.TransactionRepository
}

func NewTransactionService(r repository.TransactionRepository) *TransactionService {
	return &TransactionService{repo: r}
}

func (s *TransactionService) Create(t *models.Transaction) error {
	if t.Status == "Pending" && t.CustomerName == "" {
		return errors.New("customer name wajib diisi untuk transaksi pending")
	}
	if t.MemberID != nil {
		t.Discount = t.TotalAmount * 0.10
	}
	t.GrandTotal = t.TotalAmount - t.Discount
	return s.repo.Create(t)
}

func (s *TransactionService) UpdateStatus(id uint, status string) error {
	trans, err := s.repo.GetByID(id)
	if err != nil {
		return err
	}

	if trans.Status == "Success" {
		return errors.New("transaksi sudah sukses, tidak bisa diubah")
	}

	if status == "Success" {
		// Nanti di sini kita panggil Logika FEFO
	}

	return s.repo.UpdateStatus(id, status)
}

func (s *TransactionService) GetAll() ([]models.Transaction, error) {
	return s.repo.GetAll()
}
func (s *TransactionService) GetByID(id uint) (*models.Transaction, error) {
	return s.repo.GetByID(id)
}
func (s *TransactionService) Delete(id uint) error {
	return s.repo.Delete(id)
}
