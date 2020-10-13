package controller

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/sirupsen/logrus"
	validator "gopkg.in/go-playground/validator.v9"

	"testAPIYTS/domain"
)

// ResponseError represent the reseponse error struct
type ResponseError struct {
	Message string `json:"message"`
}

//JKHandler is
type JKHandler struct {
	Ucase domain.JadwalKajianUsecase
}

//NewJKHandler is
func NewJKHandler(e *echo.Echo, jku domain.JadwalKajianUsecase) {
	handler := &JKHandler{
		Ucase: jku,
	}
	e.GET("/jadwal", handler.GetAll)
	e.POST("/jadwal", handler.Store)
}

//GetAll is
func (a *JKHandler) GetAll(c echo.Context) error {
	ctx := c.Request().Context()

	jk, err := a.Ucase.GetAll(ctx)
	if err != nil {
		return c.JSON(getStatusCode(err), ResponseError{Message: err.Error()})
	}

	return c.JSONPretty(http.StatusOK, jk, " ")
}

//Store is
func (a *JKHandler) Store(c echo.Context) (err error) {
	var jadwalkajian domain.JadwalKajian
	err = c.Bind(&jadwalkajian)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err.Error())
	}

	var ok bool
	if ok, err = isRequestValid(&jadwalkajian); !ok {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	ctx := c.Request().Context()
	err = a.Ucase.Store(ctx, &jadwalkajian)
	if err != nil {
		return c.JSON(getStatusCode(err), ResponseError{Message: err.Error()})
	}

	return c.JSONPretty(http.StatusCreated, jadwalkajian, " ")
}

func isRequestValid(m *domain.JadwalKajian) (bool, error) {
	validate := validator.New()
	err := validate.Struct(m)
	if err != nil {
		return false, err
	}
	return true, nil
}

func getStatusCode(err error) int {
	if err == nil {
		return http.StatusOK
	}

	logrus.Error(err)
	switch err {
	case domain.ErrInternalServerError:
		return http.StatusInternalServerError
	case domain.ErrNotFound:
		return http.StatusNotFound
	case domain.ErrConflict:
		return http.StatusConflict
	default:
		return http.StatusInternalServerError
	}
}
