package postgresql

import (
	"github.com/Dimoonevs/Online_store/internal/models"
	"github.com/jackc/pgx/v4/pgxpool"
)

// SecurityRepository: defines the interface for security-related database operations.
type SecurityRepository interface {
	TestRepo() string
	GetAdmin() (models.Admin, error)
}

// SellerRepository: defines the interface for seller-related database operations.
type SellerRepository interface {
	CreateSeller(response *models.Seller) (int32, error)
	GetSellers() ([]models.SellerResponse, error)
	GetSellerByID(id int32) (models.SellerResponse, error)
	UpdateSeller(*models.Seller, int32) error
	DeleteSeller(id int32) error
}

// ProductRepository: defines the interface for product-related database operations.
type ProductRepository interface {
	CreateProduct(model *models.Product, idSeller int32) (int32, error)
	GetProductsSeller(idSeller int32) ([]models.ProductResponse, error)
	GetProductByID(id int32) (models.ProductResponse, error)
	UpdateProduct(*models.Product, int32) error
	DeleteProduct(int32) error
	GetProducts() ([]models.ProductResponse, error)
}

// Repositories: struct holds references to all repository interfaces.
type Repositories struct {
	Security SecurityRepository
	Seller   SellerRepository
	Product  ProductRepository
}

// NewRepositories: initializes and returns a new Repositories instance.
func NewRepositories(db *pgxpool.Pool) *Repositories {
	return &Repositories{
		Security: NewSecurityRepository(db),
		Seller:   NewSellerRepository(db),
		Product:  NewProductRepositoryImpl(db),
	}
}
