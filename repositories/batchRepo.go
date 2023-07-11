package repositories

import (
	"context"
	"time"
)

const createBatch = `-- name: CreateBatch :one
INSERT INTO bootcamp.batch
(batch_id, batch_entity_id, batch_name, batch_description, batch_start_date, batch_end_date, batch_reason, batch_type, batch_modified_date, batch_status, batch_pic_id)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)
RETURNING batch_id
`

type CreateBatchParams struct {
	BatchID           int32     `db:"batch_id" json:"batchId"`
	BatchEntityID     int32     `db:"batch_entity_id" json:"batchEntityId"`
	BatchName         string    `db:"batch_name" json:"batchName"`
	BatchDescription  string    `db:"batch_description" json:"batchDescription"`
	BatchStartDate    time.Time `db:"batch_start_date" json:"batchStartDate"`
	BatchEndDate      time.Time `db:"batch_end_date" json:"batchEndDate"`
	BatchReason       string    `db:"batch_reason" json:"batchReason"`
	BatchType         string    `db:"batch_type" json:"batchType"`
	BatchModifiedDate time.Time `db:"batch_modified_date" json:"batchModifiedDate"`
	BatchStatus       string    `db:"batch_status" json:"batchStatus"`
	BatchPicID        int32     `db:"batch_pic_id" json:"batchPicId"`
}

func (q *Queries) CreateBatch(ctx context.Context, arg CreateBatchParams) (int32, error) {
	row := q.db.QueryRowContext(ctx, createBatch,
		arg.BatchID,
		arg.BatchEntityID,
		arg.BatchName,
		arg.BatchDescription,
		arg.BatchStartDate,
		arg.BatchEndDate,
		arg.BatchReason,
		arg.BatchType,
		arg.BatchModifiedDate,
		arg.BatchStatus,
		arg.BatchPicID,
	)
	var batch_id int32
	err := row.Scan(&batch_id)
	return batch_id, err
}
