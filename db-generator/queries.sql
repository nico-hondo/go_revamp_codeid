-- payment.bank

-- name: GetPaymentBank :one

SELECT * FROM payment.bank WHERE bank_entity_id = $1;

-- name: ListPaymentBank :many

SELECT * FROM payment.bank ORDER BY bank_name;

-- name: CreatePaymentBank :one

INSERT INTO
    payment.bank(
        bank_entity_id,
        bank_code,
        bank_name,
        bank_modified_date
    )
VALUES ($1, $2, $3, $4) RETURNING bank_entity_id;

-- name: DeletePaymentBank :exec

DELETE FROM payment.bank WHERE bank_entity_id = $1;

UPDATE payment.bank
set
    bank_code = $2,
    bank_name = $3
WHERE bank_entity_id = $1;

-- payment.fintech

-- name: GetPaymentFintech :one

SELECT * FROM payment.fintech WHERE fint_entity_id = $1;

-- name: ListPaymentFintech :many

SELECT * FROM payment.fintech ORDER BY fint_name;

-- name: CreatePaymentFintech :one

INSERT INTO
    payment.fintech (
        fint_entity_id,
        fint_code,
        fint_name,
        fint_modified_date
    )
VALUES ($1, $2, $3, $4) RETURNING fint_entity_id;

-- name: DeletePaymentFintech :exec

DELETE FROM payment.fintech WHERE fint_entity_id = $1;

-- name: UpdatePaymentFintech :exec

UPDATE payment.fintech
set
    fint_code = $2,
    fint_name = $3
WHERE fint_entity_id = $1;

-- payment.transaction_payment

-- name: GetPaymentTransaction_payment :one

SELECT * FROM payment.transaction_payment WHERE trpa_id = $1;

-- name: ListPaymentTransaction_payment :many

SELECT * FROM payment.transaction_payment ORDER BY trpa_code_number;

-- name: CreatePaymentTransaction_payment :one

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
    ) RETURNING trpa_id;

-- name: DeletePaymentTransaction_payment :exec

DELETE FROM payment.transaction_payment WHERE trpa_id = $1;

-- name: UpdatePaymentTransaction_payment :exec

UPDATE
    payment.transaction_payment
set
    trpa_code_number = $2,
    trpa_order_number = $3,
    trpa_debit = $4,
    trpa_credit = $5,
    trpa_type = $6,
    trpa_note = $7,
    trpa_modified_date = $8,
    trpa_source_id = $9,
    trpa_target_id = $10,
    trpa_user_entity_id = $11
WHERE trpa_id = $1;

-- payment.users_account

-- name: GetPaymentUsers_account :one

SELECT * FROM payment.users_account WHERE usac_bank_entity_id = $1;

-- name: ListPaymentUsers_account :many

SELECT * FROM payment.users_account ORDER BY usac_account_number;

-- name: CreatePaymentUsers_account :one

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
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9) RETURNING usac_bank_entity_id;

-- name: DeletePaymentUsers_account :exec

DELETE FROM payment.users_account WHERE usac_bank_entity_id = $1;

-- name: UpdatePaymentUsers_account :exec

UPDATE payment.users_account
set
    usac_user_entity_id = $2,
    usac_account_number = $3,
    usac_saldo = $4,
    usac_type = $5,
    usac_start_date = $6,
    usac_end_date = $7,
    usac_modified_date = $8,
    usac_status = $9
WHERE usac_bank_entity_id = $1;