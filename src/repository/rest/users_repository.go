package rest

import (
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/dzikrisyafi/kursusvirtual_gateway-api/src/domain/cohort"
	"github.com/dzikrisyafi/kursusvirtual_gateway-api/src/domain/enrolls"
	"github.com/dzikrisyafi/kursusvirtual_utils-go/rest_errors"
	"github.com/golang-restclient/rest"
)

var (
	usersRestClient = rest.RequestBuilder{
		BaseURL: "http://localhost:8001",
		Timeout: 100 * time.Millisecond,
	}
)

type RestUsersRepository interface {
	EnrollUserCourse(enrolls.Enroll, string) (*enrolls.Enroll, rest_errors.RestErr)
	UpdateEnroll(enrolls.Enroll, string) (*enrolls.Enroll, rest_errors.RestErr)
	DeleteEnroll(int, string) rest_errors.RestErr
	CreateCohort(cohort.Cohort, string) (*cohort.Cohort, rest_errors.RestErr)
	UpdateCohort(cohort.Cohort, string) (*cohort.Cohort, rest_errors.RestErr)
	DeleteCohort(cohortID int, at string) rest_errors.RestErr
}

type usersRepository struct{}

func NewRestUsersRepository() RestUsersRepository {
	return &usersRepository{}
}

func (r *usersRepository) EnrollUserCourse(request enrolls.Enroll, at string) (*enrolls.Enroll, rest_errors.RestErr) {
	response := usersRestClient.Post("/internal/enrolls?access_token="+at, request)

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

func (r *usersRepository) UpdateEnroll(request enrolls.Enroll, at string) (*enrolls.Enroll, rest_errors.RestErr) {
	response := usersRestClient.Put(fmt.Sprintf("/internal/enrolls/%d?access_token=%s", request.ID, at), request)

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

func (r *usersRepository) DeleteEnroll(enrollID int, at string) rest_errors.RestErr {
	response := usersRestClient.Delete(fmt.Sprintf("/internal/enrolls/%d?access_token=%s", enrollID, at))

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

func (r *usersRepository) CreateCohort(request cohort.Cohort, at string) (*cohort.Cohort, rest_errors.RestErr) {
	response := usersRestClient.Post("/internal/cohorts?access_token="+at, request)

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

func (r *usersRepository) UpdateCohort(request cohort.Cohort, at string) (*cohort.Cohort, rest_errors.RestErr) {
	response := usersRestClient.Put(fmt.Sprintf("/internal/cohorts/%d?access_token=%s", request.ID, at), request)

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

func (r *usersRepository) DeleteCohort(cohortID int, at string) rest_errors.RestErr {
	response := usersRestClient.Delete(fmt.Sprintf("/internal/cohorts/%d?access_token=%s", cohortID, at))

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
