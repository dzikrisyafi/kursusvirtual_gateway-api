package enrolls

import "encoding/json"

type PublicEnroll struct {
	ID int `json:"id"`
}

type PrivateEnroll struct {
	ID       int `json:"id"`
	UserID   int `json:"user_id"`
	CourseID int `json:"course_id"`
	CohortID int `json:"cohort_id"`
}

func (enroll *Enroll) Marshall(isPublic bool) interface{} {
	if isPublic {
		return PublicEnroll{
			ID: enroll.ID,
		}
	}

	enrollJson, _ := json.Marshal(enroll)
	var privateEnroll PrivateEnroll
	json.Unmarshal(enrollJson, &privateEnroll)
	return privateEnroll
}
