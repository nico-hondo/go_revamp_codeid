package users

import (
	"context"
	"time"
)

const createUsersPhoneNumberType = `-- name: CreateUsersPhoneNumberType :one

INSERT INTO users.phone_number_type
(ponty_code, ponty_modified_date)
VALUES($1,$2)
RETURNING ponty_code
`

type CreateUsersPhoneNumberTypeParams struct {
	PontyCode         string       `db:"ponty_code" json:"pontyCode"`
	PontyModifiedDate time.Time `db:"ponty_modified_date" json:"pontyModifiedDate"`
}

func (q *Queries) CreateUsersPhoneNumberType(ctx context.Context, arg CreateUsersPhoneNumberTypeParams) (string, error) {
	row := q.db.QueryRowContext(ctx, createUsersPhoneNumberType, arg.PontyCode, arg.PontyModifiedDate)
	var ponty_code string
	err := row.Scan(&ponty_code)
	return ponty_code, err
}
