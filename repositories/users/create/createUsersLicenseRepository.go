package users

import (
	"context"
	"time"
)

const createLicense = `-- name: CreateLicense :one

INSERT INTO users.users_license
(usli_id, usli_license_code, usli_modified_date, usli_status, usli_entity_id)
VALUES($1,$2,$3,$4,$5)
RETURNING usli_id
`

type CreateLicenseParams struct {
	UsliID           int32          `db:"usli_id" json:"usliId"`
	UsliLicenseCode  string `db:"usli_license_code" json:"usliLicenseCode"`
	UsliModifiedDate time.Time   `db:"usli_modified_date" json:"usliModifiedDate"`
	UsliStatus       string `db:"usli_status" json:"usliStatus"`
	UsliEntityID     int32          `db:"usli_entity_id" json:"usliEntityId"`
}

func (q *Queries) CreateLicense(ctx context.Context, arg CreateLicenseParams) (int32, error) {
	row := q.db.QueryRowContext(ctx, createLicense,
		arg.UsliID,
		arg.UsliLicenseCode,
		arg.UsliModifiedDate,
		arg.UsliStatus,
		arg.UsliEntityID,
	)
	var usli_id int32
	err := row.Scan(&usli_id)
	return usli_id, err
}