package repositories

import (
	"context"
	"time"
)

const createPaymentBank = `-- name: CreatePaymentBank :one

INSERT INTO
    payment.bank(
        bank_entity_id,
        bank_code,
        bank_name,
        bank_modified_date
    )
VALUES ($1, $2, $3, $4) RETURNING bank_entity_id
`

type CreatePaymentBankParams struct {
	BankEntityID     int32     `db:"bank_entity_id" json:"bankEntityId"`
	BankCode         string    `db:"bank_code" json:"bankCode"`
	BankName         string    `db:"bank_name" json:"bankName"`
	BankModifiedDate time.Time `db:"bank_modified_date" json:"bankModifiedDate"`
}

func (q *Queries) CreatePaymentBank(ctx context.Context, arg CreatePaymentBankParams) (int32, error) {
	row := q.db.QueryRowContext(ctx, createPaymentBank,
		arg.BankEntityID,
		arg.BankCode,
		arg.BankName,
		arg.BankModifiedDate,
	)
	var bank_entity_id int32
	err := row.Scan(&bank_entity_id)
	return bank_entity_id, err
}

const deletePaymentBank = `-- name: DeletePaymentBank :exec

DELETE FROM payment.bank WHERE bank_entity_id = $1
`

func (q *Queries) DeletePaymentBank(ctx context.Context, bankEntityID int32) error {
	_, err := q.db.ExecContext(ctx, deletePaymentBank, bankEntityID)
	return err
}
