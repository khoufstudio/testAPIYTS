package usecase_test

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"testAPIYTS/app/jadwalkajian/usecase"
	"testAPIYTS/domain"
	"testAPIYTS/domain/mocks"
)

func TestGetAll(t *testing.T) {
	mockjkRepo := new(mocks.JKMocks)
	mockJadwal := domain.JadwalKajian{
		Title:     "Riyadush Shalihin",
		Lecturer:  "Ustadz Abduh Tuasikal",
		Start:     time.Now(),
		End:       time.Now(),
		EventDate: time.Now(),
	}

	mockListJadwal := make([]domain.JadwalKajian, 0)
	mockListJadwal = append(mockListJadwal, mockJadwal)

	t.Run("success", func(t *testing.T) {
		mockjkRepo.On("GetAll", mock.Anything).Return(mockListJadwal, nil).Once()
		u := usecase.NewJKUsecase(mockjkRepo, time.Second*2)

		list, err := u.GetAll(context.TODO())
		assert.NoError(t, err)
		assert.Len(t, list, len(mockListJadwal))

		mockjkRepo.AssertExpectations(t)
	})

	t.Run("error-failed", func(t *testing.T) {
		mockjkRepo.On("GetAll", mock.Anything).Return(nil, errors.New("Unexpexted Error")).Once()
		u := usecase.NewJKUsecase(mockjkRepo, time.Second*2)

		list, err := u.GetAll(context.TODO())
		assert.Error(t, err)
		assert.Len(t, list, 0)

		mockjkRepo.AssertExpectations(t)
	})
}

func TestStore(t *testing.T) {
	mockjkRepo := new(mocks.JKMocks)
	mockJadwal := domain.JadwalKajian{
		Title:     "Riyadush Shalihin",
		Lecturer:  "Ustadz Abduh Tuasikal",
		Start:     time.Now(),
		End:       time.Now(),
		EventDate: time.Now(),
	}

	t.Run("success", func(t *testing.T) {
		tempMockJadwal := mockJadwal
		tempMockJadwal.ID = 0
		// mockjkRepo.On("Validate", mock.AnythingOfType("*domain.JadwalKajian")).Return(nil).Once()
		mockjkRepo.On("Store", mock.Anything, mock.AnythingOfType("*domain.JadwalKajian")).Return(nil).Once()

		u := usecase.NewJKUsecase(mockjkRepo, time.Second*2)

		err := u.Store(context.TODO(), &tempMockJadwal)

		assert.Nil(t, err)
		assert.Equal(t, mockJadwal.Title, tempMockJadwal.Title)
		mockjkRepo.AssertExpectations(t)
	})
	t.Run("fail-Validation", func(t *testing.T) {
		tempMockJadwal := mockJadwal
		tempMockJadwal.Title = ""
		// mockjkRepo.On("Validate", mock.AnythingOfType("*domain.JadwalKajian")).Return(errors.New("Jadwal kajian title is empty")).Once()
		// mockjkRepo.On("Store", mock.Anything, mock.AnythingOfType("*domain.JadwalKajian")).Return(errors.New("Jadwal kajian title is empty")).Once()

		u := usecase.NewJKUsecase(mockjkRepo, time.Second*2)

		err := u.Store(context.TODO(), &tempMockJadwal)

		assert.Error(t, err)
		mockjkRepo.AssertExpectations(t)
	})
}
