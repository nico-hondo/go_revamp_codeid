-- name: GetCategory :one
SELECT * FROM jobHire.job_category
WHERE joca_id = $1;

-- name: ListCategories :many
SELECT * FROM jobHire.job_category
ORDER BY joca_name;

-- name: CreateCategory :one

INSERT INTO jobHire.job_category(joca_id, joca_name, joca_modified_date) VALUES ($1, $2, $3)
RETURNING joca_id;

-- name: DeleteCategory :exec
DELETE FROM jobHire.job_category
WHERE joca_id = $1;


-- name: GetEmployeeRange :one
SELECT * FROM jobHire.employee_range
WHERE emra_id = $1;

-- name: ListEmployeesRange :many
SELECT * FROM jobHire.employee_range
ORDER BY emra_range_max;

-- name: CreateEmployeeRange :one
INSERT INTO jobHire.employee_range(emra_id, emra_range_min, emra_range_max, emra_modified_date) VALUES ($1, $2, $3, $4)
RETURNING emra_id;

-- name: DeleteEmployeeRange :exec
DELETE FROM jobHire.employee_range
WHERE emra_id = $1;


-- name: GetClient :one
SELECT * FROM jobHire.client
WHERE clit_id = $1;

-- name: ListClient :many
SELECT * FROM jobHire.client
ORDER BY clit_name;

-- name: CreateClient :one

INSERT INTO jobHire.client(clit_id, clit_name, clit_about, clit_modified_date, clit_addr_id, clit_emra_id) VALUES ($1, $2, $3, $4, $5, $6)
RETURNING clit_id;

-- name: DeleteClient :exec
DELETE FROM jobHire.client
WHERE clit_id = $1;


-- name: GetJobPost :one
SELECT * FROM jobHire.job_post
WHERE jopo_entity_id = $1;

-- name: ListJobPost :many
SELECT * FROM jobHire.job_post
ORDER BY jopo_title;

-- name: CreateJobPost :one
INSERT INTO jobHire.job_post(jopo_entity_id, jopo_number, jopo_title, jopo_start_date, jopo_end_date, jopo_min_salary, jopo_max_salary, jopo_min_experience, jopo_max_experience, jopo_primary_skill, jopo_secondary_skill, jopo_publish_date, jopo_modified_date, jopo_emp_entity_id, jopo_clit_id, jopo_joro_id, jopo_joty_id, jopo_joca_id, jopo_addr_id, jopo_status) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20)
RETURNING jopo_entity_id;

-- name: DeleteJobPost :exec
DELETE FROM jobHire.job_post
WHERE jopo_entity_id = $1;


-- name: GetJobPostDesc :one
SELECT * FROM jobHire.job_post_desc
WHERE jopo_entity_id = $1;

-- name: ListJobPostDesc :many
SELECT * FROM jobHire.job_post_desc
ORDER BY jopo_entity_id;

-- name: CreateJobPostDesc :one
INSERT INTO jobHire.job_post_desc(jopo_entity_id, jopo_description, jopo_responsibility, jopo_target, jopo_benefit) VALUES ($1, $2, $3, $4, $5)
RETURNING jopo_entity_id;

-- name: DeleteJobPostDesc :exec
DELETE FROM jobHire.job_post_desc
WHERE jopo_entity_id = $1;


-- name: GetTalentApply :one
SELECT * FROM jobHire.talent_apply
WHERE taap_user_entity_id = $1;

-- name: ListTalentApply :many
SELECT * FROM jobHire.talent_apply
ORDER BY taap_entity_id;

-- name: CreateTalentApply :one
INSERT INTO jobHire.talent_apply(taap_user_entity_id, taap_entity_id, taap_intro, taap_scoring, taap_modified_date, taap_status) VALUES ($1, $2, $3, $4, $5, $6)
RETURNING taap_user_entity_id;

-- name: DeleteTalentApply :exec
DELETE FROM jobHire.talent_apply
WHERE taap_user_entity_id = $1;


-- name: GetTalentApplyProgress :one
SELECT * FROM jobHire.talent_apply_progress
WHERE tapr_id = $1;

-- name: ListTalentApplyProgress :many
SELECT * FROM jobHire.talent_apply_progress
ORDER BY tapr_status;

-- name: CreateTalentApplyProgress :one
INSERT INTO jobHire.talent_apply_progress(tapr_id, taap_user_entity_id, taap_entity_id, tapr_modified_date, tapr_status) VALUES ($1, $2, $3, $4, $5)
RETURNING tapr_id;

-- name: DeleteTalentApplyProgress :exec
DELETE FROM jobHire.talent_apply_progress
WHERE tapr_id = $1;