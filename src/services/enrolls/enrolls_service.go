package enrolls

import (
	"github.com/dzikrisyafi/kursusvirtual_gateway-api/src/domain/enrolls"
	"github.com/dzikrisyafi/kursusvirtual_gateway-api/src/repository/rest"
	"github.com/dzikrisyafi/kursusvirtual_utils-go/rest_errors"
)

type Service interface {
	Create(enrolls.Enroll, string) (*enrolls.Enroll, rest_errors.RestErr)
	Update(enrolls.Enroll, string) (*enrolls.Enroll, rest_errors.RestErr)
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

func (s *service) Create(request enrolls.Enroll, at string) (*enrolls.Enroll, rest_errors.RestErr) {
	if err := request.Validate(); err != nil {
		return nil, err
	}

	_, err := s.restUsersRepo.EnrollUserCourse(request, at)
	if err != nil {
		return nil, err
	}

	enrollCourse, err := s.restCoursesRepo.EnrollUserCourse(request, at)
	if err != nil {
		return nil, err
	}

	return enrollCourse, nil
}

func (s *service) Update(request enrolls.Enroll, at string) (*enrolls.Enroll, rest_errors.RestErr) {
	if err := request.Validate(); err != nil {
		return nil, err
	}

	_, err := s.restUsersRepo.UpdateEnroll(request, at)
	if err != nil {
		return nil, err
	}

	enrollCourse, err := s.restCoursesRepo.UpdateEnroll(request, at)
	if err != nil {
		return nil, err
	}

	return enrollCourse, nil
}

func (s *service) Delete(enrollID int, at string) rest_errors.RestErr {
	if err := s.restUsersRepo.DeleteEnroll(enrollID, at); err != nil {
		return err
	}

	if err := s.restCoursesRepo.DeleteEnroll(enrollID, at); err != nil {
		return err
	}

	return nil
}
