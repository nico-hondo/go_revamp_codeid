package repositories

import (
	"context"
	"time"
)

const createPaymentUsers_account = `-- name: CreatePaymentUsers_account :one

INSERT INTO
    payment.users_account (
        usac_bank_entity_id,
        usac_user_entity_id,
        usac_account_number,
        usac_saldo,
        usac_type,
        usac_start_date,
        usac_end_date,
        usac_modified_date,
        usac_status
    )
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9) RETURNING usac_bank_entity_id
`

type CreatePaymentUsers_accountParams struct {
	UsacBankEntityID  int32     `db:"usac_bank_entity_id" json:"usacBankEntityId"`
	UsacUserEntityID  int32     `db:"usac_user_entity_id" json:"usacUserEntityId"`
	UsacAccountNumber string    `db:"usac_account_number" json:"usacAccountNumber"`
	UsacSaldo         string    `db:"usac_saldo" json:"usacSaldo"`
	UsacType          string    `db:"usac_type" json:"usacType"`
	UsacStartDate     time.Time `db:"usac_start_date" json:"usacStartDate"`
	UsacEndDate       time.Time `db:"usac_end_date" json:"usacEndDate"`
	UsacModifiedDate  time.Time `db:"usac_modified_date" json:"usacModifiedDate"`
	UsacStatus        string    `db:"usac_status" json:"usacStatus"`
}

func (q *Queries) CreatePaymentUsers_account(ctx context.Context, arg CreatePaymentUsers_accountParams) (int32, error) {
	row := q.db.QueryRowContext(ctx, createPaymentUsers_account,
		arg.UsacBankEntityID,
		arg.UsacUserEntityID,
		arg.UsacAccountNumber,
		arg.UsacSaldo,
		arg.UsacType,
		arg.UsacStartDate,
		arg.UsacEndDate,
		arg.UsacModifiedDate,
		arg.UsacStatus,
	)
	var usac_bank_entity_id int32
	err := row.Scan(&usac_bank_entity_id)
	return usac_bank_entity_id, err
}

const deletePaymentUsers_account = `-- name: DeletePaymentUsers_account :exec

DELETE FROM payment.users_account WHERE usac_bank_entity_id = $1
`

func (q *Queries) DeletePaymentUsers_account(ctx context.Context, usacBankEntityID int32) error {
	_, err := q.db.ExecContext(ctx, deletePaymentUsers_account, usacBankEntityID)
	return err
}
