package users

import (
	"context"
	"time"
)

const createEducation = `-- name: CreateEducation :one

INSERT INTO users.users_education
(usdu_id, usdu_entity_id, usdu_school, usdu_degree, usdu_field_study,
usdu_graduate_year, usdu_start_date, usdu_end_date, usdu_grade,
usdu_activities, usdu_description, usdu_modified_date)
VALUES($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12)
RETURNING usdu_id
`

type CreateEducationParams struct {
	UsduID           int32          `db:"usdu_id" json:"usduId"`
	UsduEntityID     int32          `db:"usdu_entity_id" json:"usduEntityId"`
	UsduSchool       string `db:"usdu_school" json:"usduSchool"`
	UsduDegree       string `db:"usdu_degree" json:"usduDegree"`
	UsduFieldStudy   string `db:"usdu_field_study" json:"usduFieldStudy"`
	UsduGraduateYear string `db:"usdu_graduate_year" json:"usduGraduateYear"`
	UsduStartDate    time.Time   `db:"usdu_start_date" json:"usduStartDate"`
	UsduEndDate      time.Time   `db:"usdu_end_date" json:"usduEndDate"`
	UsduGrade        string `db:"usdu_grade" json:"usduGrade"`
	UsduActivities   string `db:"usdu_activities" json:"usduActivities"`
	UsduDescription  string `db:"usdu_description" json:"usduDescription"`
	UsduModifiedDate time.Time   `db:"usdu_modified_date" json:"usduModifiedDate"`
}

func (q *Queries) CreateEducation(ctx context.Context, arg CreateEducationParams) (int32, error) {
	row := q.db.QueryRowContext(ctx, createEducation,
		arg.UsduID,
		arg.UsduEntityID,
		arg.UsduSchool,
		arg.UsduDegree,
		arg.UsduFieldStudy,
		arg.UsduGraduateYear,
		arg.UsduStartDate,
		arg.UsduEndDate,
		arg.UsduGrade,
		arg.UsduActivities,
		arg.UsduDescription,
		arg.UsduModifiedDate,
	)
	var usdu_id int32
	err := row.Scan(&usdu_id)
	return usdu_id, err
}