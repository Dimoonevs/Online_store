package postgresql

import (
	"context"
	"github.com/Dimoonevs/Online_store/internal/models"
	"github.com/jackc/pgx/v4/pgxpool"
	"log"
)

type ProductRepositoryImpl struct {
	db *pgxpool.Pool
}

func NewProductRepositoryImpl(db *pgxpool.Pool) *ProductRepositoryImpl {
	return &ProductRepositoryImpl{
		db: db,
	}
}

func (p *ProductRepositoryImpl) CreateProduct(model *models.Product, idSeller int32) (int32, error) {
	var idProduct int32
	conn, err := p.db.Acquire(context.Background())
	if err != nil {
		return 0, err
	}
	defer conn.Release()

	err = conn.QueryRow(context.Background(), "INSERT INTO products (name, price, sellers_id) VALUES ($1, $2, $3) RETURNING id",
		model.Name, model.Price, idSeller).Scan(&idProduct)
	if err != nil {
		return 0, err
	}
	return idProduct, nil
}
func (p *ProductRepositoryImpl) GetProductsSeller(idSeller int32) ([]models.ProductResponse, error) {
	var products []models.ProductResponse
	conn, err := p.db.Acquire(context.Background())
	if err != nil {
		return nil, err
	}
	defer conn.Release()
	rows, err := conn.Query(context.Background(),
		`SELECT p.id, p.name, p.price
		FROM products p
		JOIN sellers s ON p.sellers_id = s.id
		WHERE p.sellers_id = $1`, idSeller)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var product models.ProductResponse
		if err := rows.Scan(&product.Id, &product.Name, &product.Price); err != nil {
			return nil, err
		}
		product.SellerId = idSeller
		products = append(products, product)
	}
	return products, nil
}
func (p *ProductRepositoryImpl) GetProductByID(id int32) (models.ProductResponse, error) {
	conn, err := p.db.Acquire(context.Background())
	if err != nil {
		return models.ProductResponse{}, err
	}
	defer conn.Release()
	var product models.ProductResponse
	err = conn.QueryRow(context.Background(), "SELECT id, name, price, sellers_id FROM products WHERE id = $1", id).
		Scan(&product.Id, &product.Name, &product.Price, &product.SellerId)
	if err != nil {
		return models.ProductResponse{}, err
	}
	return product, nil
}

func (p *ProductRepositoryImpl) UpdateProduct(model *models.Product, idProduct int32) error {
	conn, err := p.db.Acquire(context.Background())
	if err != nil {
		return err
	}
	defer conn.Release()
	_, err = conn.Exec(context.Background(), "UPDATE products SET name = $1, price = $2 WHERE id = $3",
		model.Name, model.Price, idProduct)
	if err != nil {
		return err
	}
	return nil
}
func (p *ProductRepositoryImpl) DeleteProduct(idProduct int32) error {
	conn, err := p.db.Acquire(context.Background())
	if err != nil {
		return err
	}
	defer conn.Release()
	_, err = conn.Exec(context.Background(), "DELETE FROM products WHERE id = $1", idProduct)
	if err != nil {
		return err
	}
	return nil
}
func (p *ProductRepositoryImpl) GetProducts() ([]models.ProductResponse, error) {
	var products []models.ProductResponse
	conn, err := p.db.Acquire(context.Background())
	if err != nil {
		return nil, err
	}
	defer conn.Release()
	rows, err := conn.Query(context.Background(), `
        SELECT p.id, p.name, p.price, s.sellers_id
        FROM products p
    `)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var product models.ProductResponse
		if err := rows.Scan(&product.Id, &product.Name, &product.Price, &product.SellerId); err != nil {
			return nil, err
		}
		products = append(products, product)
	}
	log.Println(products)
	return products, nil
}
