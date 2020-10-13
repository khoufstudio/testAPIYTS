package domain

import (
	"context"
	"time"
)

//JadwalKajian is ...
type JadwalKajian struct {
	ID        int64     `json:"id"`
	Title     string    `json:"title"`
	Lecturer  string    `json:"lecturer"`
	StartDate time.Time `json:"start"`
	EndDate   time.Time `json:"end"`
	EventDate time.Time `json:"event_date"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

//JadwalKajianUsecase is
type JadwalKajianUsecase interface {
	GetAll(ctx context.Context) ([]JadwalKajian, error)
	Store(ctx context.Context, jk *JadwalKajian) error
}

//JadwalKajianRepository is
type JadwalKajianRepository interface {
	GetAll(ctx context.Context) ([]JadwalKajian, error)
	Store(ctx context.Context, jk *JadwalKajian) error
}

//Table rename table for REL
func (jk JadwalKajian) Table() string {
	return "jadwal_kajian"
}
