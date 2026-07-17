package service

import (
	"github.com/luthfi/pos/models"
	"github.com/luthfi/pos/repository"
)

type MemberService struct {
	repo repository.MemberRepository
}

func NewMemberService(r repository.MemberRepository) *MemberService {
	return &MemberService{repo: r}
}

func (s *MemberService) GetAll() ([]models.Member, error) {
	return s.repo.GetAll()
}

func (s *MemberService) Search(query string) ([]models.Member, error) {
	return s.repo.Search(query)
}

func (s *MemberService) Create(m *models.Member) error {
	// Memastikan diskon selalu 10 saat pembuatan, berjaga-jaga jika input JSON kosong/berbeda
	m.Discount = 10
	return s.repo.Create(m)
}

func (s *MemberService) Update(id uint, m *models.Member) error {
	// Mencegah diskon diubah secara manual melalui update
	m.Discount = 10
	return s.repo.Update(id, m)
}

func (s *MemberService) Delete(id uint) error {
	return s.repo.Delete(id)
}
