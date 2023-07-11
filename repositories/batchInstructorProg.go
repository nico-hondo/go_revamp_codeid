package repositories

import (
	"context"
	"time"
)

const createInstructorProgram = `-- name: CreateInstructorProgram :one
INSERT INTO bootcamp.instructor_programs
(batch_id, inpro_entity_id, inpro_emp_entity_id, inpro_modified_date)
VALUES ($1, $2, $3, $4)
RETURNING batch_id
`

type CreateInstructorProgramParams struct {
	BatchID           int32     `db:"batch_id" json:"batchId"`
	InproEntityID     int32     `db:"inpro_entity_id" json:"inproEntityId"`
	InproEmpEntityID  int32     `db:"inpro_emp_entity_id" json:"inproEmpEntityId"`
	InproModifiedDate time.Time `db:"inpro_modified_date" json:"inproModifiedDate"`
}

func (q *Queries) CreateInstructorProgram(ctx context.Context, arg CreateInstructorProgramParams) (int32, error) {
	row := q.db.QueryRowContext(ctx, createInstructorProgram,
		arg.BatchID,
		arg.InproEntityID,
		arg.InproEmpEntityID,
		arg.InproModifiedDate,
	)
	var batch_id int32
	err := row.Scan(&batch_id)
	return batch_id, err
}
