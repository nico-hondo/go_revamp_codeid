package repositories

import (
	"context"
	"time"
)

const createPaymentFintech = `-- name: CreatePaymentFintech :one

INSERT INTO
    payment.fintech (
        fint_entity_id,
        fint_code,
        fint_name,
        fint_modified_date
    )
VALUES ($1, $2, $3, $4) RETURNING fint_entity_id
`

type CreatePaymentFintechParams struct {
	FintEntityID     int32     `db:"fint_entity_id" json:"fintEntityId"`
	FintCode         string    `db:"fint_code" json:"fintCode"`
	FintName         string    `db:"fint_name" json:"fintName"`
	FintModifiedDate time.Time `db:"fint_modified_date" json:"fintModifiedDate"`
}

func (q *Queries) CreatePaymentFintech(ctx context.Context, arg CreatePaymentFintechParams) (int32, error) {
	row := q.db.QueryRowContext(ctx, createPaymentFintech,
		arg.FintEntityID,
		arg.FintCode,
		arg.FintName,
		arg.FintModifiedDate,
	)
	var fint_entity_id int32
	err := row.Scan(&fint_entity_id)
	return fint_entity_id, err
}

const deletePaymentFintech = `-- name: DeletePaymentFintech :exec

DELETE FROM payment.fintech WHERE fint_entity_id = $1
`

func (q *Queries) DeletePaymentFintech(ctx context.Context, fintEntityID int32) error {
	_, err := q.db.ExecContext(ctx, deletePaymentFintech, fintEntityID)
	return err
}
