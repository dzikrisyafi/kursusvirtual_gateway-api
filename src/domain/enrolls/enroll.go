package enrolls

import (
	"github.com/dzikrisyafi/kursusvirtual_utils-go/rest_errors"
)

type Enroll struct {
	ID       int `json:"id"`
	UserID   int `json:"user_id"`
	CourseID int `json:"course_id"`
	Cohort   int `json:"cohort_id"`
}

func (enroll *Enroll) Validate() rest_errors.RestErr {
	if enroll.UserID <= 0 {
		return rest_errors.NewBadRequestError("invalid user id")
	}

	if enroll.CourseID <= 0 {
		return rest_errors.NewBadRequestError("invalid course id")
	}

	if enroll.Cohort <= 0 {
		return rest_errors.NewBadRequestError("invalid cohort id")
	}

	return nil
}
