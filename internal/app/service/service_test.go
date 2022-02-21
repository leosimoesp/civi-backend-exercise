package service

import (
	"errors"
	"reflect"
	"testing"

	"github.com/leosimoesp/civi-backend-exercise/internal/app/datatype"
	"github.com/leosimoesp/civi-backend-exercise/internal/app/repo"
	"github.com/leosimoesp/civi-backend-exercise/internal/app/repo/mocks"
)

func Test_cartesianImpl_GetPointsWithinDistance(t *testing.T) {

	points1 := []datatype.Point{
		{X: -10, Y: -25}, {X: 10, Y: 25}, {X: -10, Y: 1},
	}
	points2 := []datatype.Point{
		{X: -100, Y: 250}, {X: 6, Y: 8}, {X: 10, Y: 254}, {X: -63, Y: 120}, {X: 1, Y: 99}, {X: -7, Y: -2},
	}

	fileNotExistErr := errors.New("File not exists")

	type fields struct {
		repo repo.PointRepo
	}
	type args struct {
		p1       datatype.Point
		distance int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []datatype.Point
		wantErr error
	}{
		{
			name: "[1]-Should return a orded list of points when they're within same Manhattan distance",
			fields: fields{
				repo: &mocks.PointRepo{},
			},
			args: args{
				p1: datatype.Point{
					X: -10,
					Y: -25,
				},
				distance: 25,
			},
			want: []datatype.Point{points1[0]},
		},
		{
			name: "[2]-Should return a orded list of points when they're within same Manhattan distance",
			fields: fields{
				repo: &mocks.PointRepo{},
			},
			args: args{
				p1: datatype.Point{
					X: 0,
					Y: 0,
				},
				distance: 100,
			},
			want: []datatype.Point{points2[5], points2[1], points2[4]},
		},
		{
			name: "[3]-Should return a empty list of points if there isn't loaded points",
			fields: fields{
				repo: &mocks.PointRepo{},
			},
			args: args{
				p1: datatype.Point{
					X: 0,
					Y: 0,
				},
				distance: 100,
			},
			want: []datatype.Point{},
		},
		{
			name: "[4]-Should return error when file is wrong",
			fields: fields{
				repo: &mocks.PointRepo{},
			},
			args: args{
				p1: datatype.Point{
					X: -10,
					Y: -25,
				},
				distance: 25,
			},
			wantErr: fileNotExistErr,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := cartesianImpl{
				repo: tt.fields.repo,
			}

			switch tt.name {
			case "[1]-Should return a orded list of points when they're within same Manhattan distance":
				c.repo.(*mocks.PointRepo).On("GetAllPoints").Return(points1, nil)
			case "[2]-Should return a orded list of points when they're within same Manhattan distance":
				c.repo.(*mocks.PointRepo).On("GetAllPoints").Return(points2, nil)
			case "[3]-Should return a empty list of points if there isn't loaded points":
				c.repo.(*mocks.PointRepo).On("GetAllPoints").Return([]datatype.Point{}, nil)
			case "[4]-Should return error when file is wrong":
				c.repo.(*mocks.PointRepo).On("GetAllPoints").Return(nil, tt.wantErr)
			}

			got, err := c.GetPointsWithinDistance(tt.args.p1, tt.args.distance)
			if (err != nil && tt.wantErr == nil) || (err == nil && tt.wantErr != nil) ||
				(err != nil && tt.wantErr != nil && err.Error() != tt.wantErr.Error()) {
				t.Errorf("cartesianImpl.GetPointsWithinDistance() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("cartesianImpl.GetPointsWithinDistance() = %v, want %v", got, tt.want)
			}
		})
	}
}
