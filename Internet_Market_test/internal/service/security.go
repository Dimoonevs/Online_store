package service

import (
	"errors"
	"github.com/Dimoonevs/Online_store/internal/auth"
	"github.com/Dimoonevs/Online_store/internal/models"
	"github.com/Dimoonevs/Online_store/internal/repository/postgresql"
)

type authResponse struct {
	AccessToken string `json:"access_token"`
}

type SecurityServiceImpl struct {
	jwtWrapper auth.JwtWrapper
	repo       postgresql.SecurityRepository
}

func NewSecurityService(jwtWrapper auth.JwtWrapper, repository postgresql.SecurityRepository) SecurityService {
	return &SecurityServiceImpl{
		jwtWrapper: jwtWrapper,
		repo:       repository,
	}
}

func (s *SecurityServiceImpl) Test() string {
	return "Hello world"
}

func (s *SecurityServiceImpl) Authentication(admin *models.Admin) (*authResponse, error) {
	adminFromDb, err := s.repo.GetAdmin()
	if err != nil {
		return nil, err
	}
	if admin.Username != adminFromDb.Username {
		return nil, errors.New("invalid username")
	}
	if !auth.CheckPasswordHash(admin.Password, adminFromDb.Password) {
		return nil, errors.New("Password is incorrect")
	}
	token, err := s.jwtWrapper.GenerateToken(admin.Username)
	if err != nil {
		return nil, err
	}
	return &authResponse{
		token,
	}, nil
}

func (s *SecurityServiceImpl) ValidateToken(token string) error {
	admin, err := s.repo.GetAdmin()
	if err != nil {
		return err
	}
	clime, err := s.jwtWrapper.ValidateToken(token)
	if err != nil {
		return err
	}
	if admin.Username != clime.Usrname {
		return errors.New("Error token")
	}

	return nil
}
