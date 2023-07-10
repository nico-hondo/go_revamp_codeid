package repositories

import "context"

const deleteCategory = `-- name: DeleteCategory :exec
DELETE FROM jobHire.job_category
WHERE joca_id = $1
`

func (q *Queries) DeleteCategory(ctx context.Context, jocaID int32) error {
	_, err := q.db.ExecContext(ctx, deleteCategory, jocaID)
	return err
}

const deleteClient = `-- name: DeleteClient :exec
DELETE FROM jobHire.client
WHERE clit_id = $1
`

func (q *Queries) DeleteClient(ctx context.Context, clitID int32) error {
	_, err := q.db.ExecContext(ctx, deleteClient, clitID)
	return err
}

const deleteEmployeeRange = `-- name: DeleteEmployeeRange :exec
DELETE FROM jobHire.employee_range
WHERE emra_id = $1
`

func (q *Queries) DeleteEmployeeRange(ctx context.Context, emraID int32) error {
	_, err := q.db.ExecContext(ctx, deleteEmployeeRange, emraID)
	return err
}

const deleteJobPost = `-- name: DeleteJobPost :exec
DELETE FROM jobHire.job_post
WHERE jopo_entity_id = $1
`

func (q *Queries) DeleteJobPost(ctx context.Context, jopoEntityID int32) error {
	_, err := q.db.ExecContext(ctx, deleteJobPost, jopoEntityID)
	return err
}

const deleteJobPostDesc = `-- name: DeleteJobPostDesc :exec
DELETE FROM jobHire.job_post_desc
WHERE jopo_entity_id = $1
`

func (q *Queries) DeleteJobPostDesc(ctx context.Context, jopoEntityID int32) error {
	_, err := q.db.ExecContext(ctx, deleteJobPostDesc, jopoEntityID)
	return err
}

const deleteTalentApply = `-- name: DeleteTalentApply :exec
DELETE FROM jobHire.talent_apply
WHERE taap_user_entity_id = $1
`

func (q *Queries) DeleteTalentApply(ctx context.Context, taapUserEntityID int32) error {
	_, err := q.db.ExecContext(ctx, deleteTalentApply, taapUserEntityID)
	return err
}

const deleteTalentApplyProgress = `-- name: DeleteTalentApplyProgress :exec
DELETE FROM jobHire.talent_apply_progress
WHERE tapr_id = $1
`

func (q *Queries) DeleteTalentApplyProgress(ctx context.Context, taprID int32) error {
	_, err := q.db.ExecContext(ctx, deleteTalentApplyProgress, taprID)
	return err
}
