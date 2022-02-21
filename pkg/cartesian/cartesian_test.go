package cartesian

import (
	"testing"

	"github.com/leosimoesp/civi-backend-exercise/internal/app/datatype"
)

func TestAbs(t *testing.T) {
	type args struct {
		value int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "[1]-Should return absolute value to negative number",
			args: args{
				value: -5,
			},
			want: 5,
		},
		{
			name: "[2]-Should return absolute value to positive number",
			args: args{
				value: 5,
			},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Abs(tt.args.value); got != tt.want {
				t.Errorf("Abs() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestManhattanDistance(t *testing.T) {
	type args struct {
		p1 datatype.Point
		p2 datatype.Point
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{

			name: "[1]-Should calculate Manhattan distance between two cartesian points",
			args: args{
				p1: datatype.Point{
					X: -4,
					Y: 6,
				},
				p2: datatype.Point{
					X: 3,
					Y: -4,
				},
			},
			want: 17,
		},
		{

			name: "[2]-Should calculate Manhattan distance between two cartesian points negatives",
			args: args{
				p1: datatype.Point{
					X: -4,
					Y: -6,
				},
				p2: datatype.Point{
					X: -3,
					Y: -4,
				},
			},
			want: 3,
		},
		{

			name: "[3]-Should calculate Manhattan distance between two cartesian points positives",
			args: args{
				p1: datatype.Point{
					X: 4,
					Y: 6,
				},
				p2: datatype.Point{
					X: 3,
					Y: 4,
				},
			},
			want: 3,
		},
		{

			name: "[4]-Should calculate Manhattan distance between two cartesian points",
			args: args{
				p1: datatype.Point{
					X: 4,
					Y: -6,
				},
				p2: datatype.Point{
					X: -3,
					Y: 4,
				},
			},
			want: 17,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ManhattanDistance(tt.args.p1, tt.args.p2); got != tt.want {
				t.Errorf("ManhattanDistance() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCheckPointInManhattanDistance(t *testing.T) {
	type args struct {
		origin         datatype.Point
		p2             datatype.Point
		targetDistance int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "[1]-Should return success when point is within Manhattan distance",
			args: args{
				origin: datatype.Point{
					X: 3,
					Y: 0,
				},
				p2: datatype.Point{
					X: 4,
					Y: 6,
				},
				targetDistance: 100,
			},
			want: true,
		},
		{
			name: "[2]-Should return success when point is within Manhattan distance",
			args: args{
				origin: datatype.Point{
					X: 3,
					Y: 0,
				},
				p2: datatype.Point{
					X: 100,
					Y: 3,
				},
				targetDistance: 100,
			},
			want: true,
		},
		{
			name: "[3]-Should return success when point is within Manhattan distance",
			args: args{
				origin: datatype.Point{
					X: 3,
					Y: 0,
				},
				p2: datatype.Point{
					X: -97,
					Y: 0,
				},
				targetDistance: 100,
			},
			want: true,
		},
		{
			name: "[4]-Should return success when point is within Manhattan distance",
			args: args{
				origin: datatype.Point{
					X: 3,
					Y: 0,
				},
				p2: datatype.Point{
					X: -87,
					Y: -10,
				},
				targetDistance: 100,
			},
			want: true,
		},
		{
			name: "[5]-Should return success when point is within Manhattan distance",
			args: args{
				origin: datatype.Point{
					X: 3,
					Y: 0,
				},
				p2: datatype.Point{
					X: -3,
					Y: 0,
				},
				targetDistance: 100,
			},
			want: true,
		},
		{
			name: "[6]-Should return false when point isn't within Manhattan distance",
			args: args{
				origin: datatype.Point{
					X: 3,
					Y: 0,
				},
				p2: datatype.Point{
					X: 4,
					Y: 106,
				},
				targetDistance: 100,
			},
			want: false,
		},
		{
			name: "[7]-Should return true when origin and target are equal within Manhattan distance",
			args: args{
				origin: datatype.Point{
					X: 3,
					Y: 0,
				},
				p2: datatype.Point{
					X: 3,
					Y: 0,
				},
				targetDistance: 100,
			},
			want: true,
		},
		{
			name: "[8]-Should return false when point isn't within Manhattan distance",
			args: args{
				origin: datatype.Point{
					X: 3,
					Y: 0,
				},
				p2: datatype.Point{
					X: -98,
					Y: -1,
				},
				targetDistance: 100,
			},
			want: false,
		},
		{
			name: "[9]-Should return false when point isn't within Manhattan distance",
			args: args{
				origin: datatype.Point{
					X: 3,
					Y: 0,
				},
				p2: datatype.Point{
					X: -1,
					Y: -103,
				},
				targetDistance: 100,
			},
			want: false,
		},
		{
			name: "[10]-Should return false when point isn't within Manhattan distance",
			args: args{
				origin: datatype.Point{
					X: 3,
					Y: 0,
				},
				p2: datatype.Point{
					X: -98,
					Y: -1,
				},
				targetDistance: 100,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got, _ := CheckPointInManhattanDistance(tt.args.origin, tt.args.p2, tt.args.targetDistance); got != tt.want {
				t.Errorf("CheckPointInManhattanDistance() = %v, want %v", got, tt.want)
			}
		})
	}
}
