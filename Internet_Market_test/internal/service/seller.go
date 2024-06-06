package service

import (
	"github.com/Dimoonevs/Online_store/internal/models"
	"github.com/Dimoonevs/Online_store/internal/repository/postgresql"
)

type SellerServiceImpl struct {
	repo postgresql.SellerRepository
}

func NewSellerService(repo postgresql.SellerRepository) SellerService {
	return &SellerServiceImpl{
		repo: repo,
	}

}

func (s *SellerServiceImpl) CreateSeller(seller *models.Seller) (int32, error) {
	id, err := s.repo.CreateSeller(seller)
	if err != nil {
		return 0, err
	}
	return id, nil
}
func (s *SellerServiceImpl) GetSellers() ([]models.SellerResponse, error) {
	return s.repo.GetSellers()
}
func (s *SellerServiceImpl) GetSellerByID(id int32) (models.SellerResponse, error) {
	return s.repo.GetSellerByID(id)
}
func (s *SellerServiceImpl) UpdateSeller(model *models.Seller, id int32) error {
	return s.repo.UpdateSeller(model, id)
}
func (s *SellerServiceImpl) DeleteSeller(id int32) error {
	return s.repo.DeleteSeller(id)
}
