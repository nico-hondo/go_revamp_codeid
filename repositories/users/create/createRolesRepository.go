package users

import (
	"context"
	"time"
)

const createRoles = `-- name: CreateRoles :one

INSERT INTO users.roles 
(role_id, role_name, role_type, role_modified_date)
VALUES($1,$2,$3,$4)
RETURNING role_id
`

type CreateRolesParams struct {
	RoleID           int32    `db:"role_id" json:"roleId"`
	RoleName         string   `db:"role_name" json:"roleName"`
	RoleType         string   `db:"role_type" json:"roleType"`
	RoleModifiedDate time.Time `db:"role_modified_date" json:"roleModifiedDate"`
}

func (q *Queries) CreateRoles(ctx context.Context, arg CreateRolesParams) (int32, error) {
	row := q.db.QueryRowContext(ctx, createRoles,
		arg.RoleID,
		arg.RoleName,
		arg.RoleType,
		arg.RoleModifiedDate,
	)
	var role_id int32
	err := row.Scan(&role_id)
	return role_id, err
}