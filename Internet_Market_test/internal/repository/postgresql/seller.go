package postgresql

import (
	"context"
	"github.com/Dimoonevs/Online_store/internal/models"
	"github.com/jackc/pgx/v4/pgxpool"
)

type SellerRepositoryImpl struct {
	db *pgxpool.Pool
}

func NewSellerRepository(db *pgxpool.Pool) SellerRepository {
	return &SellerRepositoryImpl{
		db: db,
	}
}
func (s *SellerRepositoryImpl) CreateSeller(seller *models.Seller) (int32, error) {
	conn, err := s.db.Acquire(context.Background())
	var id int32
	if err != nil {
		return 0, err
	}
	defer conn.Release()
	err = conn.QueryRow(context.Background(), "INSERT INTO sellers (name, phone) VALUES ($1, $2) RETURNING id",
		seller.Name, seller.Phone).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (s *SellerRepositoryImpl) UpdateSeller(seller *models.Seller, id int32) error {
	conn, err := s.db.Acquire(context.Background())
	if err != nil {
		return err
	}
	defer conn.Release()
	_, err = conn.Query(context.Background(), "UPDATE sellers SET name = $1, phone = $2 WHERE id = $3",
		seller.Name, seller.Phone, id)
	if err != nil {
		return err
	}
	return nil
}
func (s *SellerRepositoryImpl) DeleteSeller(id int32) error {
	conn, err := s.db.Acquire(context.Background())
	if err != nil {
		return err
	}
	defer conn.Release()
	_, err = conn.Query(context.Background(), "DELETE FROM sellers WHERE id = $1", id)
	if err != nil {
		return err
	}

	return nil
}

func (s *SellerRepositoryImpl) GetSellerByID(id int32) (models.SellerResponse, error) {
	conn, err := s.db.Acquire(context.Background())
	if err != nil {
		return models.SellerResponse{}, err
	}
	defer conn.Release()
	var seller models.SellerResponse
	err = conn.QueryRow(context.Background(), "SELECT id, name, phone FROM sellers WHERE id = $1", id).Scan(&seller.Id, &seller.Name, &seller.Phone)
	if err != nil {
		return models.SellerResponse{}, err
	}
	return seller, nil
}
func (s *SellerRepositoryImpl) GetSellers() ([]models.SellerResponse, error) {
	conn, err := s.db.Acquire(context.Background())
	if err != nil {
		return nil, err
	}
	defer conn.Release()
	var sellers []models.SellerResponse
	rows, err := conn.Query(context.Background(), "SELECT id, name, phone FROM sellers")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var seller models.SellerResponse
		if err := rows.Scan(&seller.Id, &seller.Name, &seller.Phone); err != nil {
			return nil, err
		}
		sellers = append(sellers, seller)
	}
	return sellers, nil
}
