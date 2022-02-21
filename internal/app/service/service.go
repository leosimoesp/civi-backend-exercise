package service

import (
	"sort"

	"github.com/leosimoesp/civi-backend-exercise/internal/app/datatype"
	"github.com/leosimoesp/civi-backend-exercise/internal/app/repo"
	"github.com/leosimoesp/civi-backend-exercise/pkg/cartesian"
)

type CartesianService interface {
	GetPointsWithinDistance(p1 datatype.Point, distance int) ([]datatype.Point, error)
}

type cartesianImpl struct {
	repo repo.PointRepo
}

func NewCartesianService(repo repo.PointRepo) CartesianService {
	return cartesianImpl{
		repo: repo,
	}
}

func (c cartesianImpl) GetPointsWithinDistance(p1 datatype.Point, distance int) ([]datatype.Point, error) {

	type filterPoint struct {
		point    datatype.Point
		distance int
	}

	pointsFilter := []filterPoint{}

	points, err := c.repo.GetAllPoints()

	if err != nil {
		return nil, err
	}

	for _, point := range points {
		if ok, distance := cartesian.CheckPointInManhattanDistance(p1, point, distance); ok {
			pointsFilter = append(pointsFilter, filterPoint{point: point, distance: distance})
		}
	}

	sort.SliceStable(pointsFilter, func(i, j int) bool {
		return pointsFilter[i].distance < pointsFilter[j].distance
	})

	resultPoints := []datatype.Point{}

	for _, filtered := range pointsFilter {
		resultPoints = append(resultPoints, filtered.point)
	}

	return resultPoints, nil
}
