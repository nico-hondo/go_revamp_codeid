package users

import (
	"context"
	"time"
)

const createAddrees = `-- name: CreateAddrees :one

INSERT INTO users.users_address
(etad_addr_id, etad_modified_date, etad_entity_id, etad_adty_id)
VALUES($1,$2,$3,$4)
RETURNING etad_addr_id
`

type CreateAddreesParams struct {
	EtadAddrID       int32         `db:"etad_addr_id" json:"etadAddrId"`
	EtadModifiedDate time.Time  `db:"etad_modified_date" json:"etadModifiedDate"`
	EtadEntityID     int32 `db:"etad_entity_id" json:"etadEntityId"`
	EtadAdtyID       int32 `db:"etad_adty_id" json:"etadAdtyId"`
}

func (q *Queries) CreateAddrees(ctx context.Context, arg CreateAddreesParams) (int32, error) {
	row := q.db.QueryRowContext(ctx, createAddrees,
		arg.EtadAddrID,
		arg.EtadModifiedDate,
		arg.EtadEntityID,
		arg.EtadAdtyID,
	)
	var etad_addr_id int32
	err := row.Scan(&etad_addr_id)
	return etad_addr_id, err
}