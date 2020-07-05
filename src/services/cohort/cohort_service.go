package cohort

import (
	"github.com/dzikrisyafi/kursusvirtual_gateway-api/src/domain/cohort"
	"github.com/dzikrisyafi/kursusvirtual_gateway-api/src/repository/rest"
	"github.com/dzikrisyafi/kursusvirtual_utils-go/rest_errors"
)

type Service interface {
	Create(cohort.Cohort, string) (*cohort.Cohort, rest_errors.RestErr)
	Update(cohort.Cohort, string) (*cohort.Cohort, rest_errors.RestErr)
	Delete(int, string) rest_errors.RestErr
}

type service struct {
	restUsersRepo   rest.RestUsersRepository
	restCoursesRepo rest.RestCoursesRepository
}

func NewService(usersRepo rest.RestUsersRepository, coursesRepo rest.RestCoursesRepository) Service {
	return &service{
		restUsersRepo:   usersRepo,
		restCoursesRepo: coursesRepo,
	}
}

func (s *service) Create(request cohort.Cohort, at string) (*cohort.Cohort, rest_errors.RestErr) {
	if err := request.Validate(); err != nil {
		return nil, err
	}

	if _, err := s.restUsersRepo.CreateCohort(request, at); err != nil {
		return nil, err
	}

	cohort, err := s.restCoursesRepo.CreateCohort(request, at)
	if err != nil {
		return nil, err
	}

	return cohort, nil
}

func (s *service) Update(request cohort.Cohort, at string) (*cohort.Cohort, rest_errors.RestErr) {
	if err := request.Validate(); err != nil {
		return nil, err
	}

	if _, err := s.restUsersRepo.UpdateCohort(request, at); err != nil {
		return nil, err
	}

	cohort, err := s.restCoursesRepo.UpdateCohort(request, at)
	if err != nil {
		return nil, err
	}

	return cohort, nil
}

func (s *service) Delete(cohortID int, at string) rest_errors.RestErr {
	if err := s.restUsersRepo.DeleteCohort(cohortID, at); err != nil {
		return err
	}

	if err := s.restCoursesRepo.DeleteCohort(cohortID, at); err != nil {
		return err
	}

	return nil
}
