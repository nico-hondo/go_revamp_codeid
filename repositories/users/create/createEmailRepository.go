package users

import (
	"context"
	"time"
)

const createEmail = `-- name: CreateEmail :one

INSERT INTO users.users_email
(pmail_entity_id, pmail_id, pmail_address, pmail_modified_date)
VALUES($1,$2,$3,$4)
RETURNING pmail_id
`

type CreateEmailParams struct {
	PmailEntityID     int32          `db:"pmail_entity_id" json:"pmailEntityId"`
	PmailID           int32          `db:"pmail_id" json:"pmailId"`
	PmailAddress      string `db:"pmail_address" json:"pmailAddress"`
	PmailModifiedDate time.Time   `db:"pmail_modified_date" json:"pmailModifiedDate"`
}

func (q *Queries) CreateEmail(ctx context.Context, arg CreateEmailParams) (int32, error) {
	row := q.db.QueryRowContext(ctx, createEmail,
		arg.PmailEntityID,
		arg.PmailID,
		arg.PmailAddress,
		arg.PmailModifiedDate,
	)
	var pmail_id int32
	err := row.Scan(&pmail_id)
	return pmail_id, err
}