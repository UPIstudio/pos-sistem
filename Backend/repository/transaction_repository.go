package repository

import (
	"github.com/luthfi/pos/config"
	"github.com/luthfi/pos/models"
	"gorm.io/gorm"
)

type TransactionRepository interface {
	Create(t *models.Transaction) error
	Delete(id uint) error
	GetAll() ([]models.Transaction, error)
	GetByID(id uint) (*models.Transaction, error)
	UpdateStatus(id uint, status string) error
}

type transactionRepo struct{}

func NewTransactionRepository() TransactionRepository { return &transactionRepo{} }

func (r *transactionRepo) Create(t *models.Transaction) error {
	return config.DB.Create(t).Error
}

func (r *transactionRepo) Delete(id uint) error {
	return config.DB.Transaction(func(tx *gorm.DB) error {
		tx.Where("transaction_id = ?", id).Delete(&models.TransactionItem{})
		return tx.Delete(&models.Transaction{}, id).Error
	})
}

func (r *transactionRepo) GetAll() (t []models.Transaction, err error) {
	err = config.DB.Preload("Items").Preload("Member").Find(&t).Error
	return
}

func (r *transactionRepo) GetByID(id uint) (*models.Transaction, error) {
	var t models.Transaction
	err := config.DB.Preload("Items").Preload("Member").First(&t, id).Error
	return &t, err
}

func (r *transactionRepo) UpdateStatus(id uint, status string) error {
	return config.DB.Model(&models.Transaction{}).Where("id = ?", id).Update("status", status).Error
}
