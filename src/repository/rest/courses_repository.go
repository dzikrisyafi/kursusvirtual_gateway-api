package rest

import (
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/dzikrisyafi/kursusvirtual_gateway-api/src/domain/cohort"
	"github.com/dzikrisyafi/kursusvirtual_gateway-api/src/domain/enrolls"
	"github.com/dzikrisyafi/kursusvirtual_utils-go/rest_errors"
	"github.com/mercadolibre/golang-restclient/rest"
)

var (
	coursesRestClient = rest.RequestBuilder{
		BaseURL: "http://localhost:8002",
		Timeout: 100 * time.Millisecond,
	}
)

type RestCoursesRepository interface {
	EnrollUserCourse(enrolls.Enroll, string) (*enrolls.Enroll, rest_errors.RestErr)
	UpdateEnroll(enrolls.Enroll, string) (*enrolls.Enroll, rest_errors.RestErr)
	DeleteEnroll(int, string) rest_errors.RestErr
	CreateCohort(cohort.Cohort, string) (*cohort.Cohort, rest_errors.RestErr)
	UpdateCohort(cohort.Cohort, string) (*cohort.Cohort, rest_errors.RestErr)
	DeleteCohort(cohortID int, at string) rest_errors.RestErr
}

type coursesRepository struct{}

func NewRestCoursesRepository() RestCoursesRepository {
	return &coursesRepository{}
}

func (r *coursesRepository) EnrollUserCourse(request enrolls.Enroll, at string) (*enrolls.Enroll, rest_errors.RestErr) {
	response := coursesRestClient.Post("/internal/enrolls?access_token="+at, request)

	if response == nil || response.Response == nil {
		return nil, rest_errors.NewInternalServerError("invalid rest client response when trying to enroll user course", errors.New("rest client error"))
	}

	if response.StatusCode > 299 {
		apiErr, err := rest_errors.NewRestErrorFromBytes(response.Bytes())
		if err != nil {
			return nil, rest_errors.NewInternalServerError("invalid error interface when trying to enroll user course", err)
		}
		return nil, apiErr
	}

	var enroll enrolls.Enroll
	if err := json.Unmarshal(response.Bytes(), &enroll); err != nil {
		return nil, rest_errors.NewInternalServerError("error when trying to unmarshal enroll user course response", errors.New("json parsing error"))
	}

	return &enroll, nil
}

func (r *coursesRepository) UpdateEnroll(request enrolls.Enroll, at string) (*enrolls.Enroll, rest_errors.RestErr) {
	response := coursesRestClient.Put(fmt.Sprintf("/internal/enrolls/%d?access_token=%s", request.ID, at), request)

	if response == nil || response.Response == nil {
		return nil, rest_errors.NewInternalServerError("invalid rest client response when trying to enroll user course", errors.New("rest client error"))
	}

	if response.StatusCode > 299 {
		apiErr, err := rest_errors.NewRestErrorFromBytes(response.Bytes())
		if err != nil {
			return nil, rest_errors.NewInternalServerError("invalid error interface when trying to enroll user course", err)
		}
		return nil, apiErr
	}

	var enroll enrolls.Enroll
	if err := json.Unmarshal(response.Bytes(), &enroll); err != nil {
		return nil, rest_errors.NewInternalServerError("error when trying to unmarshal enroll user course response", errors.New("json parsing error"))
	}

	return &enroll, nil
}

func (r *coursesRepository) DeleteEnroll(enrollID int, at string) rest_errors.RestErr {
	response := coursesRestClient.Delete(fmt.Sprintf("/internal/enrolls/%d?access_token=%s", enrollID, at))

	if response == nil || response.Response == nil {
		return rest_errors.NewInternalServerError("invalid rest client response when trying to delete user enroll", errors.New("rest client error"))
	}

	if response.StatusCode > 299 {
		apiErr, err := rest_errors.NewRestErrorFromBytes(response.Bytes())
		if err != nil {
			return rest_errors.NewInternalServerError("invalid error interface when trying to delete user enroll", err)
		}
		return apiErr
	}

	return nil
}

func (r *coursesRepository) CreateCohort(request cohort.Cohort, at string) (*cohort.Cohort, rest_errors.RestErr) {
	response := coursesRestClient.Post("/internal/cohorts?access_token="+at, request)

	if response == nil || response.Response == nil {
		return nil, rest_errors.NewInternalServerError("invalid rest client response when trying save cohort", errors.New("rest client error"))
	}

	if response.StatusCode > 299 {
		apiErr, err := rest_errors.NewRestErrorFromBytes(response.Bytes())
		if err != nil {
			return nil, rest_errors.NewInternalServerError("invalid error interface when trying to save cohort", err)
		}

		return nil, apiErr
	}

	var cohort cohort.Cohort
	if err := json.Unmarshal(response.Bytes(), &cohort); err != nil {
		return nil, rest_errors.NewInternalServerError("error when trying to unmarshal cohort response", errors.New("json parsing error"))
	}

	return &cohort, nil
}

func (r *coursesRepository) UpdateCohort(request cohort.Cohort, at string) (*cohort.Cohort, rest_errors.RestErr) {
	response := coursesRestClient.Put(fmt.Sprintf("/internal/cohorts/%d?access_token=%s", request.ID, at), request)

	if response == nil || response.Response == nil {
		return nil, rest_errors.NewInternalServerError("invalid rest client response when trying save cohort", errors.New("rest client error"))
	}

	if response.StatusCode > 299 {
		apiErr, err := rest_errors.NewRestErrorFromBytes(response.Bytes())
		if err != nil {
			return nil, rest_errors.NewInternalServerError("invalid error interface when trying to save cohort", err)
		}
		return nil, apiErr
	}

	var cohort cohort.Cohort
	if err := json.Unmarshal(response.Bytes(), &cohort); err != nil {
		return nil, rest_errors.NewInternalServerError("error when trying to unmarshal cohort response", errors.New("json parsing error"))
	}

	return &cohort, nil
}

func (r *coursesRepository) DeleteCohort(cohortID int, at string) rest_errors.RestErr {
	response := coursesRestClient.Delete(fmt.Sprintf("/internal/cohorts/%d?access_token=%s", cohortID, at))

	if response == nil || response.Response == nil {
		return rest_errors.NewInternalServerError("invalid rest client response when trying to delete cohort", errors.New("rest client error"))
	}

	if response.StatusCode > 299 {
		apiErr, err := rest_errors.NewRestErrorFromBytes(response.Bytes())
		if err != nil {
			return rest_errors.NewInternalServerError("invalid error interface when trying to delete cohort", err)
		}
		return apiErr
	}

	return nil
}
