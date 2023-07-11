package users

import "context"

const createExperienceSkill = `-- name: CreateExperienceSkill :one

INSERT INTO users.users_experiences_skill
(uesk_usex_id, uesk_uski_id)
VALUES($1,$2)
RETURNING uesk_usex_id
`

type CreateExperienceSkillParams struct {
	UeskUsexID int32 `db:"uesk_usex_id" json:"ueskUsexId"`
	UeskUskiID int32 `db:"uesk_uski_id" json:"ueskUskiId"`
}

func (q *Queries) CreateExperienceSkill(ctx context.Context, arg CreateExperienceSkillParams) (int32, error) {
	row := q.db.QueryRowContext(ctx, createExperienceSkill, arg.UeskUsexID, arg.UeskUskiID)
	var uesk_usex_id int32
	err := row.Scan(&uesk_usex_id)
	return uesk_usex_id, err
}