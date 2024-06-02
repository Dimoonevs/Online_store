package service

import (
	"github.com/Dimoonevs/Online_store/internal/models"
	"github.com/Dimoonevs/Online_store/internal/repository/postgresql"
)

type ProductServiceImpl struct {
	repo postgresql.ProductRepository
}

func NewProductServiceImpl(repo postgresql.ProductRepository) *ProductServiceImpl {
	return &ProductServiceImpl{repo: repo}
}

func (p *ProductServiceImpl) CreateProduct(product *models.Product, idSeller int32) (int32, error) {
	return p.repo.CreateProduct(product, idSeller)
}
func (p *ProductServiceImpl) GetProductsSeller(idSeller int32) ([]models.ProductResponse, error) {
	return p.repo.GetProductsSeller(idSeller)
}
func (p *ProductServiceImpl) GetProductByID(id int32) (models.ProductResponse, error) {
	return p.repo.GetProductByID(id)
}
func (p *ProductServiceImpl) UpdateProduct(model *models.Product, idProduct int32) error {
	return p.repo.UpdateProduct(model, idProduct)
}
func (p *ProductServiceImpl) DeleteProduct(idProduct int32) error {
	return p.repo.DeleteProduct(idProduct)
}
func (p *ProductServiceImpl) GetProducts() ([]models.ProductResponse, error) {
	return p.repo.GetProducts()
}
