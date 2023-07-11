package repositories

import (
	"context"
	"time"
)

const createPaymentTransaction_payment = `-- name: CreatePaymentTransaction_payment :one

INSERT INTO
    payment.transaction_payment (
        trpa_id,
        trpa_code_number,
        trpa_order_number,
        trpa_debit,
        trpa_credit,
        trpa_type,
        trpa_note,
        trpa_modified_date,
        trpa_source_id,
        trpa_target_id,
        trpa_user_entity_id
    )
VALUES (
        $1,
        $2,
        $3,
        $4,
        $5,
        $6,
        $7,
        $8,
        $9,
        $10,
        $11
    ) RETURNING trpa_id
`

type CreatePaymentTransaction_paymentParams struct {
	TrpaID           int32     `db:"trpa_id" json:"trpaId"`
	TrpaCodeNumber   string    `db:"trpa_code_number" json:"trpaCodeNumber"`
	TrpaOrderNumber  string    `db:"trpa_order_number" json:"trpaOrderNumber"`
	TrpaDebit        string    `db:"trpa_debit" json:"trpaDebit"`
	TrpaCredit       string    `db:"trpa_credit" json:"trpaCredit"`
	TrpaType         string    `db:"trpa_type" json:"trpaType"`
	TrpaNote         string    `db:"trpa_note" json:"trpaNote"`
	TrpaModifiedDate time.Time `db:"trpa_modified_date" json:"trpaModifiedDate"`
	TrpaSourceID     string    `db:"trpa_source_id" json:"trpaSourceId"`
	TrpaTargetID     string    `db:"trpa_target_id" json:"trpaTargetId"`
	TrpaUserEntityID int32     `db:"trpa_user_entity_id" json:"trpaUserEntityId"`
}

func (q *Queries) CreatePaymentTransaction_payment(ctx context.Context, arg CreatePaymentTransaction_paymentParams) (int32, error) {
	row := q.db.QueryRowContext(ctx, createPaymentTransaction_payment,
		arg.TrpaID,
		arg.TrpaCodeNumber,
		arg.TrpaOrderNumber,
		arg.TrpaDebit,
		arg.TrpaCredit,
		arg.TrpaType,
		arg.TrpaNote,
		arg.TrpaModifiedDate,
		arg.TrpaSourceID,
		arg.TrpaTargetID,
		arg.TrpaUserEntityID,
	)
	var trpa_id int32
	err := row.Scan(&trpa_id)
	return trpa_id, err
}

const deletePaymentTransaction_payment = `-- name: DeletePaymentTransaction_payment :exec

DELETE FROM payment.transaction_payment WHERE trpa_id = $1
`

func (q *Queries) DeletePaymentTransaction_payment(ctx context.Context, trpaID int32) error {
	_, err := q.db.ExecContext(ctx, deletePaymentTransaction_payment, trpaID)
	return err
}
