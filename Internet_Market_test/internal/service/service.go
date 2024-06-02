package service

import (
	"github.com/Dimoonevs/Online_store/internal/auth"
	"github.com/Dimoonevs/Online_store/internal/models"
	"github.com/Dimoonevs/Online_store/internal/repository/postgresql"
)

// SecurityService: defines the interface for security-related operations.
type SecurityService interface {
	Test() string
	Authentication(admin *models.Admin) (*authResponse, error)
	ValidateToken(token string) error
}

// SellerService: defines the interface for seller-related operations
type SellerService interface {
	CreateSeller(seller *models.Seller) (int32, error)
	GetSellers() ([]models.SellerResponse, error)
	GetSellerByID(id int32) (models.SellerResponse, error)
	UpdateSeller(*models.Seller, int32) error
	DeleteSeller(id int32) error
}

// ProductService: defines the interface for product-related operations.
type ProductService interface {
	CreateProduct(model *models.Product, idSeller int32) (int32, error)
	GetProductsSeller(idSeller int32) ([]models.ProductResponse, error)
	GetProductByID(id int32) (models.ProductResponse, error)
	UpdateProduct(*models.Product, int32) error
	DeleteProduct(int32) error
	GetProducts() ([]models.ProductResponse, error)
}

// Service: struct holds references to all service interfaces.
type Service struct {
	Security SecurityService
	Seller   SellerService
	Product  ProductService
}

// NewService: initializes and returns a new Service instance.
func NewService(jwtWrapper auth.JwtWrapper, repositories *postgresql.Repositories) *Service {
	return &Service{
		Security: NewSecurityService(jwtWrapper, repositories.Security),
		Seller:   NewSellerService(repositories.Seller),
		Product:  NewProductServiceImpl(repositories.Product),
	}
}
