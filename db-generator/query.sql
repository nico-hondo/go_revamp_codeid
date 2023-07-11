-- curriculum.program_entity

-- name: GetProgramEntity :one
SELECT * FROM curriculum.program_entity
WHERE prog_entity_id = $1;

-- name: ListProgramEntity :many
SELECT * FROM curriculum.program_entity
ORDER BY prog_title;

-- name: CreateProgramEntity :one

INSERT INTO curriculum.program_entity (prog_entity_id, prog_title, prog_headline, prog_type, prog_learning_type, prog_rating, prog_total_trainee, prog_image, prog_best_seller, prog_price, prog_language, prog_modified_date, prog_duration, prog_duration_type, prog_tag_skill, prog_city_id, prog_cate_id, prog_created_by, prog_status) 
VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13,$14,$15,$16,$17,$18,$19)
RETURNING prog_entity_id;

-- name: DeleteProgramEntity :exec
DELETE FROM curriculum.program_entity
WHERE prog_entity_id = $1;


UPDATE curriculum.program_entity
  set prog_title = $2,
  prog_headline = $3,
  prog_type = $4,
  prog_learning_type = $5,
  prog_rating = $6,
  prog_total_trainee = $7,
  prog_image = $8,
  prog_best_seller = $9,
  prog_price = $10,
  prog_language = $11,
  prog_modified_date = $12,
  prog_duration = $13
WHERE prog_entity_id = $1;

-- Program Entity Description

-- name: GetProgramEntityDescription :one
SELECT * FROM curriculum.program_entity_description
WHERE pred_prog_entity_id = $1;

-- name: ListProgramEntityDescription :many
SELECT * FROM curriculum.program_entity_description
ORDER BY pred_item_learning;

-- name: CreateProgramEntityDescription :one

INSERT INTO curriculum.program_entity_description 
(pred_prog_entity_id, pred_item_learning, pred_item_include, pred_requirement, pred_description, pred_target_level)
VALUES($1,$2,$3,$4,$5,$6)
RETURNING pred_prog_entity_id;

-- name: DeleteProgramEntityDescription :exec
DELETE FROM curriculum.program_entity_description
WHERE pred_prog_entity_id = $1;

-- name: UpdateProgramEntityDescription :exec
UPDATE curriculum.program_entity_description
  set pred_item_learning = $2,
  pred_item_include = $3,
  pred_requirement = $4,
  pred_description = $5,
  pred_target_level = $6
WHERE pred_prog_entity_id = $1;

-- Program Reviews

-- name: GetProgramReviews :one
SELECT * FROM curriculum.program_reviews
WHERE prow_user_entity_id = $1;

-- name: ListProgramReviews :many
SELECT * FROM curriculum.program_reviews
ORDER BY prow_prog_entity_id;

-- name: CreateProgramReviews :one

INSERT INTO curriculum.program_reviews (prow_user_entity_id, prow_prog_entity_id, prow_review, prow_rating, prow_modified_date) 
VALUES ($1,$2,$3,$4,$5)
RETURNING prow_user_entity_id;

-- name: DeleteProgramReviews :exec
DELETE FROM curriculum.program_reviews
WHERE prow_user_entity_id = $1;

-- name: UpdateProgramReviews :exec
UPDATE curriculum.program_reviews
  set prow_prog_entity_id = $2,
  prow_review = $3,
  prow_rating = $4,
  prow_modified_date = $5
WHERE prow_user_entity_id = $1;

-- Sections

-- name: GetSections :one
SELECT * FROM curriculum.sections
WHERE sect_title = $1;

-- name: ListSections :many
SELECT * FROM curriculum.sections
ORDER BY sect_description;

-- name: CreateSections :one

INSERT INTO curriculum.sections (sect_title, sect_description, sect_total_section, sect_total_lecture, sect_total_minute, sect_modified_date, sect_prog_entity_id, sect_id) 
VALUES ($1,$2,$3,$4,$5,$6,$7,$8)
RETURNING sect_title;

-- name: DeleteSections :exec
DELETE FROM curriculum.sections
WHERE sect_title = $1;

-- name: UpdateSections :exec
UPDATE curriculum.sections
  set sect_description = $2,
  sect_total_section = $3,
  sect_total_lecture = $4,
  sect_total_minute = $5,
  sect_modified_date = $6,
  sect_prog_entity_id = $7,
  sect_id = $8
WHERE sect_title = $1;

-- Section Detail 

-- name: GetSectionDetail :one
SELECT * FROM curriculum.section_detail
WHERE secd_id = $1;

-- name: ListSectionDetail :many
SELECT * FROM curriculum.section_detail
ORDER BY secd_title;

-- name: CreateSectionDetail :one

INSERT INTO curriculum.section_detail (secd_id, secd_title, secd_preview, secd_score, secd_note, secd_minute, secd_modified_date, secd_sect_id) 
VALUES ($1,$2,$3,$4,$5,$6,$7,$8)
RETURNING secd_id;

-- name: DeleteSectionDetail :exec
DELETE FROM curriculum.section_detail
WHERE secd_id = $1;

-- name: UpdateSectionDetail :exec
UPDATE curriculum.section_detail
  set secd_title = $2,
  secd_preview = $3,
  secd_score = $4,
  secd_note = $5,
  secd_minute = $6,
  secd_modified_date = $7,
  secd_sect_id = $8
WHERE secd_id = $1;


-- Section Detail Material

-- name: GetSectionDetailMaterial :one
SELECT * FROM curriculum.section_detail_material
WHERE sedm_id = $1;

-- name: ListSectionDetailMaterial :many
SELECT * FROM curriculum.section_detail_material
ORDER BY sedm_filename;

-- name: CreateSectionDetailMaterial :one

INSERT INTO curriculum.section_detail_material (sedm_id, sedm_filename, sedm_filesize, sedm_filetype, sedm_filelink, sedm_modified_date, sedm_secd_id) 
VALUES ($1,$2,$3,$4,$5,$6,$7)
RETURNING sedm_id;

-- name: DeleteSectionDetailMaterial :exec
DELETE FROM curriculum.section_detail_material
WHERE sedm_id = $1;

-- name: UpdateSectionDetailMaterial :exec
UPDATE curriculum.section_detail_material
  set sedm_filename = $2,
  sedm_filesize = $3,
  sedm_filetype = $4,
  sedm_filelink = $5,
  sedm_modified_date = $6,
  sedm_secd_id = $7
WHERE sedm_id = $1;
