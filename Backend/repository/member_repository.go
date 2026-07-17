package repository

import (
	"github.com/luthfi/pos/config"
	"github.com/luthfi/pos/models"
)

type MemberRepository interface {
	GetAll() ([]models.Member, error)
	Search(query string) ([]models.Member, error)
	Create(m *models.Member) error
	Update(id uint, m *models.Member) error
	Delete(id uint) error
}

type memberRepo struct{}

func NewMemberRepository() MemberRepository { return &memberRepo{} }

func (r *memberRepo) GetAll() (members []models.Member, err error) {
	err = config.DB.Find(&members).Error
	return
}

func (r *memberRepo) Search(query string) (members []models.Member, err error) {
	err = config.DB.Where("name ILIKE ? OR phone LIKE ?", "%"+query+"%", "%"+query+"%").Find(&members).Error
	return
}

func (r *memberRepo) Create(m *models.Member) error {
	return config.DB.Create(m).Error
}

func (r *memberRepo) Update(id uint, m *models.Member) error {
	return config.DB.Model(&models.Member{}).Where("id = ?", id).Updates(m).Error
}

func (r *memberRepo) Delete(id uint) error {
	return config.DB.Delete(&models.Member{}, id).Error
}
