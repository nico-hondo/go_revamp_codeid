package repositories

import (
	"context"
	"database/sql"
)

const createCategory = `-- name: CreateCategory :one

INSERT INTO jobHire.job_category(joca_id, joca_name, joca_modified_date) VALUES ($1, $2, $3)
RETURNING joca_id
`

type CreateCategoryParams struct {
	JocaID           int32        `db:"joca_id" json:"jocaId"`
	JocaName         string       `db:"joca_name" json:"jocaName"`
	JocaModifiedDate sql.NullTime `db:"joca_modified_date" json:"jocaModifiedDate"`
}

func (q *Queries) CreateCategory(ctx context.Context, arg CreateCategoryParams) (int32, error) {
	row := q.db.QueryRowContext(ctx, createCategory, arg.JocaID, arg.JocaName, arg.JocaModifiedDate)
	var joca_id int32
	err := row.Scan(&joca_id)
	return joca_id, err
}

const createClient = `-- name: CreateClient :one

INSERT INTO jobHire.client(clit_id, clit_name, clit_about, clit_modified_date, clit_addr_id, clit_emra_id) VALUES ($1, $2, $3, $4, $5, $6)
RETURNING clit_id
`

type CreateClientParams struct {
	ClitID           int32        `db:"clit_id" json:"clitId"`
	ClitName         string       `db:"clit_name" json:"clitName"`
	ClitAbout        string       `db:"clit_about" json:"clitAbout"`
	ClitModifiedDate sql.NullTime `db:"clit_modified_date" json:"clitModifiedDate"`
	ClitAddrID       int32        `db:"clit_addr_id" json:"clitAddrId"`
	ClitEmraID       int32        `db:"clit_emra_id" json:"clitEmraId"`
}

func (q *Queries) CreateClient(ctx context.Context, arg CreateClientParams) (int32, error) {
	row := q.db.QueryRowContext(ctx, createClient,
		arg.ClitID,
		arg.ClitName,
		arg.ClitAbout,
		arg.ClitModifiedDate,
		arg.ClitAddrID,
		arg.ClitEmraID,
	)
	var clit_id int32
	err := row.Scan(&clit_id)
	return clit_id, err
}

const createEmployeeRange = `-- name: CreateEmployeeRange :one
INSERT INTO jobHire.employee_range(emra_id, emra_range_min, emra_range_max, emra_modified_date) VALUES ($1, $2, $3, $4)
RETURNING emra_id
`

type CreateEmployeeRangeParams struct {
	EmraID           int32        `db:"emra_id" json:"emraId"`
	EmraRangeMin     int32        `db:"emra_range_min" json:"emraRangeMin"`
	EmraRangeMax     int32        `db:"emra_range_max" json:"emraRangeMax"`
	EmraModifiedDate sql.NullTime `db:"emra_modified_date" json:"emraModifiedDate"`
}

func (q *Queries) CreateEmployeeRange(ctx context.Context, arg CreateEmployeeRangeParams) (int32, error) {
	row := q.db.QueryRowContext(ctx, createEmployeeRange,
		arg.EmraID,
		arg.EmraRangeMin,
		arg.EmraRangeMax,
		arg.EmraModifiedDate,
	)
	var emra_id int32
	err := row.Scan(&emra_id)
	return emra_id, err
}

const createJobPost = `-- name: CreateJobPost :one
INSERT INTO jobHire.job_post(jopo_entity_id, jopo_number, jopo_title, jopo_start_date, jopo_end_date, jopo_min_salary, jopo_max_salary, jopo_min_experience, jopo_max_experience, jopo_primary_skill, jopo_secondary_skill, jopo_publish_date, jopo_modified_date, jopo_emp_entity_id, jopo_clit_id, jopo_joro_id, jopo_joty_id, jopo_joca_id, jopo_addr_id, jopo_status) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20)
RETURNING jopo_entity_id
`

type CreateJobPostParams struct {
	JopoEntityID       int32        `db:"jopo_entity_id" json:"jopoEntityId"`
	JopoNumber         string       `db:"jopo_number" json:"jopoNumber"`
	JopoTitle          string       `db:"jopo_title" json:"jopoTitle"`
	JopoStartDate      sql.NullTime `db:"jopo_start_date" json:"jopoStartDate"`
	JopoEndDate        sql.NullTime `db:"jopo_end_date" json:"jopoEndDate"`
	JopoMinSalary      uint32       `db:"jopo_min_salary" json:"jopoMinSalary"`
	JopoMaxSalary      uint32       `db:"jopo_max_salary" json:"jopoMaxSalary"`
	JopoMinExperience  int32        `db:"jopo_min_experience" json:"jopoMinExperience"`
	JopoMaxExperience  int32        `db:"jopo_max_experience" json:"jopoMaxExperience"`
	JopoPrimarySkill   string       `db:"jopo_primary_skill" json:"jopoPrimarySkill"`
	JopoSecondarySkill string       `db:"jopo_secondary_skill" json:"jopoSecondarySkill"`
	JopoPublishDate    sql.NullTime `db:"jopo_publish_date" json:"jopoPublishDate"`
	JopoModifiedDate   sql.NullTime `db:"jopo_modified_date" json:"jopoModifiedDate"`
	JopoEmpEntityID    int32        `db:"jopo_emp_entity_id" json:"jopoEmpEntityId"`
	JopoClitID         int32        `db:"jopo_clit_id" json:"jopoClitId"`
	JopoJoroID         int32        `db:"jopo_joro_id" json:"jopoJoroId"`
	JopoJotyID         int32        `db:"jopo_joty_id" json:"jopoJotyId"`
	JopoJocaID         int32        `db:"jopo_joca_id" json:"jopoJocaId"`
	JopoAddrID         int32        `db:"jopo_addr_id" json:"jopoAddrId"`
	JopoStatus         string       `db:"jopo_status" json:"jopoStatus"`
}

func (q *Queries) CreateJobPost(ctx context.Context, arg CreateJobPostParams) (int32, error) {
	row := q.db.QueryRowContext(ctx, createJobPost,
		arg.JopoEntityID,
		arg.JopoNumber,
		arg.JopoTitle,
		arg.JopoStartDate,
		arg.JopoEndDate,
		arg.JopoMinSalary,
		arg.JopoMaxSalary,
		arg.JopoMinExperience,
		arg.JopoMaxExperience,
		arg.JopoPrimarySkill,
		arg.JopoSecondarySkill,
		arg.JopoPublishDate,
		arg.JopoModifiedDate,
		arg.JopoEmpEntityID,
		arg.JopoClitID,
		arg.JopoJoroID,
		arg.JopoJotyID,
		arg.JopoJocaID,
		arg.JopoAddrID,
		arg.JopoStatus,
	)
	var jopo_entity_id int32
	err := row.Scan(&jopo_entity_id)
	return jopo_entity_id, err
}

const createJobPostDesc = `-- name: CreateJobPostDesc :one
INSERT INTO jobHire.job_post_desc(jopo_entity_id, jopo_description, jopo_responsibility, jopo_target, jopo_benefit) VALUES ($1, $2, $3, $4, $5)
RETURNING jopo_entity_id
`

type CreateJobPostDescParams struct {
	JopoEntityID       int32          `db:"jopo_entity_id" json:"jopoEntityId"`
	JopoDescription    sql.NullString `db:"jopo_description" json:"jopoDescription"`
	JopoResponsibility sql.NullString `db:"jopo_responsibility" json:"jopoResponsibility"`
	JopoTarget         sql.NullString `db:"jopo_target" json:"jopoTarget"`
	JopoBenefit        sql.NullString `db:"jopo_benefit" json:"jopoBenefit"`
}

func (q *Queries) CreateJobPostDesc(ctx context.Context, arg CreateJobPostDescParams) (int32, error) {
	row := q.db.QueryRowContext(ctx, createJobPostDesc,
		arg.JopoEntityID,
		arg.JopoDescription,
		arg.JopoResponsibility,
		arg.JopoTarget,
		arg.JopoBenefit,
	)
	var jopo_entity_id int32
	err := row.Scan(&jopo_entity_id)
	return jopo_entity_id, err
}

const createTalentApply = `-- name: CreateTalentApply :one
INSERT INTO jobHire.talent_apply(taap_user_entity_id, taap_entity_id, taap_intro, taap_scoring, taap_modified_date, taap_status) VALUES ($1, $2, $3, $4, $5, $6)
RETURNING taap_user_entity_id
`

type CreateTalentApplyParams struct {
	TaapUserEntityID int32        `db:"taap_user_entity_id" json:"taapUserEntityId"`
	TaapEntityID     int32        `db:"taap_entity_id" json:"taapEntityId"`
	TaapIntro        string       `db:"taap_intro" json:"taapIntro"`
	TaapScoring      int32        `db:"taap_scoring" json:"taapScoring"`
	TaapModifiedDate sql.NullTime `db:"taap_modified_date" json:"taapModifiedDate"`
	TaapStatus       string       `db:"taap_status" json:"taapStatus"`
}

func (q *Queries) CreateTalentApply(ctx context.Context, arg CreateTalentApplyParams) (int32, error) {
	row := q.db.QueryRowContext(ctx, createTalentApply,
		arg.TaapUserEntityID,
		arg.TaapEntityID,
		arg.TaapIntro,
		arg.TaapScoring,
		arg.TaapModifiedDate,
		arg.TaapStatus,
	)
	var taap_user_entity_id int32
	err := row.Scan(&taap_user_entity_id)
	return taap_user_entity_id, err
}

const createTalentApplyProgress = `-- name: CreateTalentApplyProgress :one
INSERT INTO jobHire.talent_apply_progress(tapr_id, taap_user_entity_id, taap_entity_id, tapr_modified_date, tapr_status) VALUES ($1, $2, $3, $4, $5)
RETURNING tapr_id
`

type CreateTalentApplyProgressParams struct {
	TaprID           int32        `db:"tapr_id" json:"taprId"`
	TaapUserEntityID int32        `db:"taap_user_entity_id" json:"taapUserEntityId"`
	TaapEntityID     int32        `db:"taap_entity_id" json:"taapEntityId"`
	TaprModifiedDate sql.NullTime `db:"tapr_modified_date" json:"taprModifiedDate"`
	TaprStatus       string       `db:"tapr_status" json:"taprStatus"`
}

func (q *Queries) CreateTalentApplyProgress(ctx context.Context, arg CreateTalentApplyProgressParams) (int32, error) {
	row := q.db.QueryRowContext(ctx, createTalentApplyProgress,
		arg.TaprID,
		arg.TaapUserEntityID,
		arg.TaapEntityID,
		arg.TaprModifiedDate,
		arg.TaprStatus,
	)
	var tapr_id int32
	err := row.Scan(&tapr_id)
	return tapr_id, err
}
