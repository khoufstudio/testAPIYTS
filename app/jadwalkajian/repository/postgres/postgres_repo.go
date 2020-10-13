package postgres

import (
	"context"

	"github.com/go-rel/rel"
	"testAPIYTS/domain"
)

//PostgreJKRepo is
type PostgreJKRepo struct {
	repo rel.Repository
}

// NewPostgreJKRepo will create an object that represent the JadwalKajian.Repository interface
func NewPostgreJKRepo(repo rel.Repository) domain.JadwalKajianRepository {
	return &PostgreJKRepo{repo}
}

//GetAll is
func (p *PostgreJKRepo) GetAll(ctx context.Context) ([]domain.JadwalKajian, error) {
	var jk []domain.JadwalKajian
	if err := p.repo.FindAll(ctx, &jk); err != nil {
		return []domain.JadwalKajian{}, err
	}

	return jk, nil
}

//Store is
func (p *PostgreJKRepo) Store(ctx context.Context, jk *domain.JadwalKajian) error {
	err := p.repo.Insert(ctx, jk)
	if err != nil {
		return err
	}

	return nil
}
