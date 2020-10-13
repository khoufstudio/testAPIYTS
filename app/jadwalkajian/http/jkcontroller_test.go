package controller_test

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	_jadwalhttp "testAPIYTS/app/jadwalkajian/http"
	"testAPIYTS/domain"
	"testAPIYTS/domain/mocks"
)

func TestGetAll(t *testing.T) {
	mockjkUcase := new(mocks.JKMocks)
	mockJadwal := domain.JadwalKajian{
		ID:        1,
		Title:     "Riyadush Shalihin",
		Lecturer:  "Ustadz Abduh Tuasikal",
		Start:     time.Now(),
		End:       time.Now(),
		EventDate: time.Now(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	mockListJadwal := make([]domain.JadwalKajian, 0)
	mockListJadwal = append(mockListJadwal, mockJadwal)

	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/jadwal", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	mockjkUcase.On("GetAll", mock.Anything).Return(mockListJadwal, nil).Once()

	handler := _jadwalhttp.JKHandler{
		Ucase: mockjkUcase,
	}

	err := handler.GetAll(c)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)

	mockjkUcase.AssertExpectations(t)
}
