package users

import (
	"context"
	"time"
)

const createPhones = `-- name: CreatePhones :one

INSERT INTO users.users_phones
(uspo_entity_id, uspo_number, uspo_modified_date, uspo_ponty_code)
VALUES($1,$2,$3,$4)
RETURNING uspo_entity_id
`

type CreatePhonesParams struct {
	UspoEntityID     int32          `db:"uspo_entity_id" json:"uspoEntityId"`
	UspoNumber       string         `db:"uspo_number" json:"uspoNumber"`
	UspoModifiedDate time.Time   `db:"uspo_modified_date" json:"uspoModifiedDate"`
	UspoPontyCode    string `db:"uspo_ponty_code" json:"uspoPontyCode"`
}

func (q *Queries) CreatePhones(ctx context.Context, arg CreatePhonesParams) (int32, error) {
	row := q.db.QueryRowContext(ctx, createPhones,
		arg.UspoEntityID,
		arg.UspoNumber,
		arg.UspoModifiedDate,
		arg.UspoPontyCode,
	)
	var uspo_entity_id int32
	err := row.Scan(&uspo_entity_id)
	return uspo_entity_id, err
}
