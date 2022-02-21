package repo

import (
	"github.com/leosimoesp/civi-backend-exercise/internal/app/datatype"
)

type PointRepo interface {
	GetAllPoints() ([]datatype.Point, error)
}

type repoImpl struct {
}

func NewRepo() PointRepo {
	return repoImpl{}
}

func (r repoImpl) GetAllPoints() ([]datatype.Point, error) {
	return LoadPoints(), nil
}
