package postgres

import (
	"context"

	"github.com/go-rel/rel"
	"testAPIYTS/domain"
)

//MySQLJKRepo is
type MySQLJKRepo struct {
	repo rel.Repository
}

// NewMySQLJKRepo will create an object that represent the JadwalKajian.Repository interface
func NewMySQLJKRepo(repo rel.Repository) domain.JadwalKajianRepository {
	return &MySQLJKRepo{repo}
}

//GetAll is
func (p *MySQLJKRepo) GetAll(ctx context.Context) ([]domain.JadwalKajian, error) {
	var jk []domain.JadwalKajian
	if err := p.repo.FindAll(ctx, &jk); err != nil {
		return []domain.JadwalKajian{}, err
	}

	return jk, nil
}

//Store is
func (p *MySQLJKRepo) Store(ctx context.Context, jk *domain.JadwalKajian) error {
	err := p.repo.Insert(ctx, jk)
	if err != nil {
		return err
	}

	return nil
}
