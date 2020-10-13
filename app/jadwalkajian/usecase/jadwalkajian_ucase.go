package usecase

import (
	"context"
	"time"

	"testAPIYTS/domain"
)

//JKUsecase is ...
type JKUsecase struct {
	jkRepo         domain.JadwalKajianRepository
	contextTimeout time.Duration
}

//NewJKUsecase is
func NewJKUsecase(jkr domain.JadwalKajianRepository, timeout time.Duration) domain.JadwalKajianUsecase {
	return &JKUsecase{
		jkRepo:         jkr,
		contextTimeout: timeout,
	}
}

//GetAll is ...
func (jku *JKUsecase) GetAll(c context.Context) ([]domain.JadwalKajian, error) {
	ctx, cancel := context.WithTimeout(c, jku.contextTimeout)
	defer cancel()

	res, err := jku.jkRepo.GetAll(ctx)
	if err != nil {
		return nil, err
	}
	return res, nil
}

//Store is ...
func (jku *JKUsecase) Store(c context.Context, jk *domain.JadwalKajian) error {
	ctx, cancel := context.WithTimeout(c, jku.contextTimeout)
	defer cancel()

	errRepo := jku.jkRepo.Store(ctx, jk)
	if errRepo != nil {
		return errRepo
	}
	return nil
}
