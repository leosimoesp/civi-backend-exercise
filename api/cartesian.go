package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strconv"

	"github.com/leosimoesp/civi-backend-exercise/internal/app/datatype"
	"github.com/leosimoesp/civi-backend-exercise/internal/app/service"
)

const (
	ErrMandatoryMessage = "URL parameter is missing in the request."
)

type ErrorMessage struct {
	Message string `json:"message"`
}

type ResponseError struct {
	Errors []ErrorMessage `json:"errors"`
}

type apiRoutes struct {
	service service.CartesianService
}

func NewApiRoutes(service service.CartesianService) apiRoutes {
	return apiRoutes{
		service: service,
	}
}

//HandleGetPointsWithinDistance returns list of points that are within distance from x,y, using the Manhattan distance
func (api apiRoutes) HandleGetPointsWithinDistance() {

	http.HandleFunc("/api/points", func(w http.ResponseWriter, r *http.Request) {

		queryParams := r.URL.Query()

		//validate query params
		validateErrors := checkQueryParams(queryParams)
		if len(validateErrors) > 0 {

			errorsAsJson, err := json.Marshal(ResponseError{Errors: validateErrors})

			if err != nil {
				log.Printf("Error when marshal validate errors %v\n", err)
			}

			w.WriteHeader(http.StatusBadRequest)
			_, err = fmt.Fprintf(w, "%v\n", string(errorsAsJson))

			if err != nil {
				log.Printf("Error when marshal validate errors %v\n", err)
			}
		} else {
			w.WriteHeader(http.StatusOK)
			distance, err := strconv.Atoi(queryParams.Get("distance"))

			if err != nil {
				errorsAsJson, err := json.Marshal(ResponseError{Errors: []ErrorMessage{{Message: err.Error()}}})
				log.Printf("Error parse param %v\n", err)
				fmt.Fprintf(w, "%v\n", string(errorsAsJson))
			}

			x, err := strconv.Atoi(queryParams.Get("x"))

			if err != nil {
				errorsAsJson, err := json.Marshal(ResponseError{Errors: []ErrorMessage{{Message: err.Error()}}})
				log.Printf("Error parse param %v\n", err)
				fmt.Fprintf(w, "%v\n", string(errorsAsJson))
			}

			y, err := strconv.Atoi(queryParams.Get("y"))

			if err != nil {
				errorsAsJson, err := json.Marshal(ResponseError{Errors: []ErrorMessage{{Message: err.Error()}}})
				log.Printf("Error parse param %v\n", err)
				fmt.Fprintf(w, "%v\n", string(errorsAsJson))
			}

			filteredPoints, err := api.service.GetPointsWithinDistance(datatype.Point{X: x, Y: y}, distance)

			if err != nil {
				errorsAsJson, err := json.Marshal(ResponseError{Errors: []ErrorMessage{{Message: err.Error()}}})
				log.Printf("Error when call service GetPointsWithinDistance %v\n", err)
				fmt.Fprintf(w, "%v\n", string(errorsAsJson))
			}

			pointsAsJson, _ := json.Marshal(filteredPoints)

			_, err = fmt.Fprintf(w, "%v\n", string(pointsAsJson))

			if err != nil {
				log.Printf("Error when marshal response list %v\n", err)
			}
		}
	})
}

func checkQueryParams(params url.Values) []ErrorMessage {
	hasErrors := []ErrorMessage{}
	if _, exist := params["x"]; !exist {
		hasErrors = append(hasErrors, ErrorMessage{
			Message: fmt.Sprintf("x %s", ErrMandatoryMessage),
		})
	}
	if _, exist := params["y"]; !exist {
		hasErrors = append(hasErrors, ErrorMessage{
			Message: fmt.Sprintf("y %s", ErrMandatoryMessage),
		})
	}
	if _, exist := params["distance"]; !exist {
		hasErrors = append(hasErrors, ErrorMessage{
			Message: fmt.Sprintf("distance %s", ErrMandatoryMessage),
		})
	}
	return hasErrors
}
