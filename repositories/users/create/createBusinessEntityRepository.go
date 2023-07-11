package users

import "context"

const createBusinessEntity = `-- name: CreateBusinessEntity :one
INSERT INTO users.business_entity 
(entity_id)
VALUES($1)
RETURNING entity_id
`

type CreateBusinessEntityParams struct {
	EntityID           int32    `db:"entity_id" json:"entityId"`
}

func (q *Queries) CreateBusinessEntity(ctx context.Context, entityID int32) (int32, error) {
	row := q.db.QueryRowContext(ctx, createBusinessEntity, entityID)
	var entity_id int32
	err := row.Scan(&entity_id)
	return entity_id, err
}