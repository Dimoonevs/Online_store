package postgresql

import (
	"context"
	"github.com/Dimoonevs/Online_store/internal/models"
	"github.com/jackc/pgx/v4/pgxpool"
)

type SecurityRepositoryImpl struct {
	db *pgxpool.Pool
}

func NewSecurityRepository(db *pgxpool.Pool) SecurityRepository {
	return &SecurityRepositoryImpl{
		db: db,
	}
}

func (r *SecurityRepositoryImpl) TestRepo() string {
	return "string()"
}

func (r *SecurityRepositoryImpl) GetAdmin() (models.Admin, error) {
	conn, err := r.db.Acquire(context.Background())
	if err != nil {
		return models.Admin{}, err
	}
	defer conn.Release()
	var admin models.Admin
	err = conn.QueryRow(context.Background(), "SELECT username, password FROM admin LIMIT 1").Scan(&admin.Username, &admin.Password)
	if err != nil {
		return models.Admin{}, err
	}
	return admin, nil
}
