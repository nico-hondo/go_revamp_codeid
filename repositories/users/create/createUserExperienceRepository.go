package users

import (
	"context"
	"time"
)

const createExperience = `-- name: CreateExperience :one

INSERT INTO users.users_experiences
(usex_id, usex_entity_id, usex_title, usex_profile_headline, usex_employment_type,
usex_company_name, usex_is_current, usex_start_date, usex_end_date, usex_industry,
usex_description, usex_experience_type, usex_city_id)
VALUES($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13)
RETURNING usex_id
`

type CreateExperienceParams struct {
	UsexID              int32          `db:"usex_id" json:"usexId"`
	UsexEntityID        int32          `db:"usex_entity_id" json:"usexEntityId"`
	UsexTitle           string `db:"usex_title" json:"usexTitle"`
	UsexProfileHeadline string `db:"usex_profile_headline" json:"usexProfileHeadline"`
	UsexEmploymentType  string `db:"usex_employment_type" json:"usexEmploymentType"`
	UsexCompanyName     string `db:"usex_company_name" json:"usexCompanyName"`
	UsexIsCurrent       string `db:"usex_is_current" json:"usexIsCurrent"`
	UsexStartDate       time.Time   `db:"usex_start_date" json:"usexStartDate"`
	UsexEndDate         time.Time   `db:"usex_end_date" json:"usexEndDate"`
	UsexIndustry        string `db:"usex_industry" json:"usexIndustry"`
	UsexDescription     string `db:"usex_description" json:"usexDescription"`
	UsexExperienceType  string `db:"usex_experience_type" json:"usexExperienceType"`
	UsexCityID          int32  `db:"usex_city_id" json:"usexCityId"`
}

func (q *Queries) CreateExperience(ctx context.Context, arg CreateExperienceParams) (int32, error) {
	row := q.db.QueryRowContext(ctx, createExperience,
		arg.UsexID,
		arg.UsexEntityID,
		arg.UsexTitle,
		arg.UsexProfileHeadline,
		arg.UsexEmploymentType,
		arg.UsexCompanyName,
		arg.UsexIsCurrent,
		arg.UsexStartDate,
		arg.UsexEndDate,
		arg.UsexIndustry,
		arg.UsexDescription,
		arg.UsexExperienceType,
		arg.UsexCityID,
	)
	var usex_id int32
	err := row.Scan(&usex_id)
	return usex_id, err
}
