package users

import (
	"context"
	"time"
)

const createUsersRoles = `-- name: CreateUsersRoles :one

INSERT INTO users.users_roles
(usro_entity_id, usro_role_id, usro_modified_date)
VALUES($1,$2,$3)
RETURNING usro_entity_id
`

type CreateUsersRolesParams struct {
	UsroEntityID     int32        `db:"usro_entity_id" json:"usroEntityId"`
	UsroRoleID       int32        `db:"usro_role_id" json:"usroRoleId"`
	UsroModifiedDate time.Time `db:"usro_modified_date" json:"usroModifiedDate"`
}

func (q *Queries) CreateUsersRoles(ctx context.Context, arg CreateUsersRolesParams) (int32, error) {
	row := q.db.QueryRowContext(ctx, createUsersRoles, arg.UsroEntityID, arg.UsroRoleID, arg.UsroModifiedDate)
	var usro_entity_id int32
	err := row.Scan(&usro_entity_id)
	return usro_entity_id, err
}