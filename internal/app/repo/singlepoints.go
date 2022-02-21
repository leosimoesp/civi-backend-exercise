package repo

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"sync"

	"log"

	"github.com/leosimoesp/civi-backend-exercise/internal/app/datatype"
)

var points []datatype.Point

func loadPoints(path string) []datatype.Point {

	bytesRead, err := ioutil.ReadFile(fmt.Sprintf("./%s", path))

	if err != nil {
		log.Printf("LoadPoints read file error %v\n", err)
	}

	var resultPoints []datatype.Point

	err = json.Unmarshal(bytesRead, &resultPoints)

	if err != nil {
		log.Printf("LoadPoints unmarshal error %v\n", err)
	}

	return resultPoints
}

func LoadPoints() []datatype.Point {
	return getPoints()
}

var once sync.Once

func getPoints() []datatype.Point {

	if len(points) == 0 {
		path := os.Getenv("POINTS_PATH")
		once.Do(
			func() {
				points = loadPoints(path)
			})
	}

	return points
}
