package postgres_test

import (
	"context"
	"testing"
	"time"

	"github.com/go-rel/rel/reltest"
	"github.com/stretchr/testify/assert"
	postgres_repo "testAPIYTS/app/jadwalkajian/repository/postgres"
	"testAPIYTS/domain"
)

func TestGetAll(t *testing.T) {
	assert := assert.New(t)
	var (
		repo         = reltest.New()
		jadwalkajian = []domain.JadwalKajian{
			{
				ID:        1,
				Title:     "Riyadush Shalihin",
				Lecturer:  "Ustadz Abduh Tuasikal",
				Start:     time.Now(),
				End:       time.Now(),
				EventDate: time.Now(),
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			},
			{
				ID:        2,
				Title:     "Tadzkiratus Saami'",
				Lecturer:  "Ustadz Nuzul",
				Start:     time.Now(),
				End:       time.Now(),
				EventDate: time.Now(),
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			},
		}
	)

	repo.ExpectFindAll().Result(jadwalkajian)
	a := postgres_repo.NewPostgreJKRepo(repo)

	anJadwal, err := a.GetAll(context.TODO())
	assert.NoError(err)
	assert.NotNil(anJadwal)
	assert.Equal(anJadwal, jadwalkajian)
}

func TestStore_Success(t *testing.T) {
	now := time.Now()
	repo := reltest.New()
	jk := &domain.JadwalKajian{
		Title:     "Riyadush Shalihin",
		Lecturer:  "Ustadz Abduh Tuasikal",
		Start:     now,
		End:       now,
		EventDate: now,
		CreatedAt: now,
		UpdatedAt: now,
	}

	repo.ExpectInsert().For(jk)
	a := postgres_repo.NewPostgreJKRepo(repo)

	err := a.Store(context.TODO(), jk)
	assert.NoError(t, err)
	assert.NotEmpty(t, jk.ID)
}
