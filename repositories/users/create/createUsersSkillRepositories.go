package users

import (
	"context"
	"time"
)

const createSkill = `-- name: CreateSkill :one

INSERT INTO users.users_skill
(uski_id, uski_entity_id, uski_modified_date, uski_skty_name)
VALUES($1,$2,$3,$4)
RETURNING uski_id
`

type CreateSkillParams struct {
	UskiID           int32          `db:"uski_id" json:"uskiId"`
	UskiEntityID     int32          `db:"uski_entity_id" json:"uskiEntityId"`
	UskiModifiedDate time.Time   `db:"uski_modified_date" json:"uskiModifiedDate"`
	UskiSktyName     string `db:"uski_skty_name" json:"uskiSktyName"`
}

func (q *Queries) CreateSkill(ctx context.Context, arg CreateSkillParams) (int32, error) {
	row := q.db.QueryRowContext(ctx, createSkill,
		arg.UskiID,
		arg.UskiEntityID,
		arg.UskiModifiedDate,
		arg.UskiSktyName,
	)
	var uski_id int32
	err := row.Scan(&uski_id)
	return uski_id, err
}