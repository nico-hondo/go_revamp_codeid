package repositories

import (
	"context"
	"database/sql"
	"time"
)

const createProgramEntity = `-- name: CreateProgramEntity :one

INSERT INTO curriculum.program_entity (prog_entity_id, prog_title, prog_headline, prog_type, prog_learning_type, prog_rating, prog_total_trainee, prog_image, prog_best_seller, prog_price, prog_language, prog_modified_date, prog_duration, prog_duration_type, prog_tag_skill, prog_city_id, prog_cate_id, prog_created_by, prog_status) 
VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13,$14,$15,$16,$17,$18,$19)
RETURNING prog_entity_id
`

type CreateProgramEntityParams struct {
	ProgEntityID     int32     `db:"prog_entity_id" json:"progEntityId"`
	ProgTitle        string    `db:"prog_title" json:"progTitle"`
	ProgHeadline     string    `db:"prog_headline" json:"progHeadline"`
	ProgType         string    `db:"prog_type" json:"progType"`
	ProgLearningType string    `db:"prog_learning_type" json:"progLearningType"`
	ProgRating       int32     `db:"prog_rating" json:"progRating"`
	ProgTotalTrainee int32     `db:"prog_total_trainee" json:"progTotalTrainee"`
	ProgImage        string    `db:"prog_image" json:"progImage"`
	ProgBestSeller   string    `db:"prog_best_seller" json:"progBestSeller"`
	ProgPrice        int32     `db:"prog_price" json:"progPrice"`
	ProgLanguage     string    `db:"prog_language" json:"progLanguage"`
	ProgModifiedDate time.Time `db:"prog_modified_date" json:"progModifiedDate"`
	ProgDuration     int32     `db:"prog_duration" json:"progDuration"`
	ProgDurationType string    `db:"prog_duration_type" json:"progDurationType"`
	ProgTagSkill     string    `db:"prog_tag_skill" json:"progTagSkill"`
	ProgCityID       int32     `db:"prog_city_id" json:"progCityId"`
	ProgCateID       int32     `db:"prog_cate_id" json:"progCateId"`
	ProgCreatedBy    int32     `db:"prog_created_by" json:"progCreatedBy"`
	ProgStatus       string    `db:"prog_status" json:"progStatus"`
}

func (q *Queries) CreateProgramEntity(ctx context.Context, arg CreateProgramEntityParams) (int32, error) {
	row := q.db.QueryRowContext(ctx, createProgramEntity,
		arg.ProgEntityID,
		arg.ProgTitle,
		arg.ProgHeadline,
		arg.ProgType,
		arg.ProgLearningType,
		arg.ProgRating,
		arg.ProgTotalTrainee,
		arg.ProgImage,
		arg.ProgBestSeller,
		arg.ProgPrice,
		arg.ProgLanguage,
		arg.ProgModifiedDate,
		arg.ProgDuration,
		arg.ProgDurationType,
		arg.ProgTagSkill,
		arg.ProgCityID,
		arg.ProgCateID,
		arg.ProgCreatedBy,
		arg.ProgStatus,
	)
	var prog_entity_id int32
	err := row.Scan(&prog_entity_id)
	return prog_entity_id, err
}

const createProgramEntityDescription = `-- name: CreateProgramEntityDescription :one

INSERT INTO curriculum.program_entity_description 
(pred_prog_entity_id, pred_item_learning, pred_item_include, pred_requirement, pred_description, pred_target_level)
VALUES($1,$2,$3,$4,$5,$6)
RETURNING pred_prog_entity_id
`

type CreateProgramEntityDescriptionParams struct {
	PredProgEntityID int32          `db:"pred_prog_entity_id" json:"predProgEntityId"`
	PredItemLearning sql.NullString `db:"pred_item_learning" json:"predItemLearning"`
	PredItemInclude  sql.NullString `db:"pred_item_include" json:"predItemInclude"`
	PredRequirement  sql.NullString `db:"pred_requirement" json:"predRequirement"`
	PredDescription  sql.NullString `db:"pred_description" json:"predDescription"`
	PredTargetLevel  sql.NullString `db:"pred_target_level" json:"predTargetLevel"`
}

func (q *Queries) CreateProgramEntityDescription(ctx context.Context, arg CreateProgramEntityDescriptionParams) (int32, error) {
	row := q.db.QueryRowContext(ctx, createProgramEntityDescription,
		arg.PredProgEntityID,
		arg.PredItemLearning,
		arg.PredItemInclude,
		arg.PredRequirement,
		arg.PredDescription,
		arg.PredTargetLevel,
	)
	var pred_prog_entity_id int32
	err := row.Scan(&pred_prog_entity_id)
	return pred_prog_entity_id, err
}

const createProgramReviews = `-- name: CreateProgramReviews :one

INSERT INTO curriculum.program_reviews (prow_user_entity_id, prow_prog_entity_id, prow_review, prow_rating, prow_modified_date) 
VALUES ($1,$2,$3,$4,$5)
RETURNING prow_user_entity_id
`

type CreateProgramReviewsParams struct {
	ProwUserEntityID int32     `db:"prow_user_entity_id" json:"prowUserEntityId"`
	ProwProgEntityID int32     `db:"prow_prog_entity_id" json:"prowProgEntityId"`
	ProwReview       string    `db:"prow_review" json:"prowReview"`
	ProwRating       int32     `db:"prow_rating" json:"prowRating"`
	ProwModifiedDate time.Time `db:"prow_modified_date" json:"prowModifiedDate"`
}

func (q *Queries) CreateProgramReviews(ctx context.Context, arg CreateProgramReviewsParams) (int32, error) {
	row := q.db.QueryRowContext(ctx, createProgramReviews,
		arg.ProwUserEntityID,
		arg.ProwProgEntityID,
		arg.ProwReview,
		arg.ProwRating,
		arg.ProwModifiedDate,
	)
	var prow_user_entity_id int32
	err := row.Scan(&prow_user_entity_id)
	return prow_user_entity_id, err
}

const createSectionDetail = `-- name: CreateSectionDetail :one

INSERT INTO curriculum.section_detail (secd_id, secd_title, secd_preview, secd_score, secd_note, secd_minute, secd_modified_date, secd_sect_id) 
VALUES ($1,$2,$3,$4,$5,$6,$7,$8)
RETURNING secd_id
`

type CreateSectionDetailParams struct {
	SecdID           int32     `db:"secd_id" json:"secdId"`
	SecdTitle        string    `db:"secd_title" json:"secdTitle"`
	SecdPreview      string    `db:"secd_preview" json:"secdPreview"`
	SecdScore        int32     `db:"secd_score" json:"secdScore"`
	SecdNote         string    `db:"secd_note" json:"secdNote"`
	SecdMinute       int32     `db:"secd_minute" json:"secdMinute"`
	SecdModifiedDate time.Time `db:"secd_modified_date" json:"secdModifiedDate"`
	SecdSectID       int32     `db:"secd_sect_id" json:"secdSectId"`
}

func (q *Queries) CreateSectionDetail(ctx context.Context, arg CreateSectionDetailParams) (int32, error) {
	row := q.db.QueryRowContext(ctx, createSectionDetail,
		arg.SecdID,
		arg.SecdTitle,
		arg.SecdPreview,
		arg.SecdScore,
		arg.SecdNote,
		arg.SecdMinute,
		arg.SecdModifiedDate,
		arg.SecdSectID,
	)
	var secd_id int32
	err := row.Scan(&secd_id)
	return secd_id, err
}

const createSectionDetailMaterial = `-- name: CreateSectionDetailMaterial :one

INSERT INTO curriculum.section_detail_material (sedm_id, sedm_filename, sedm_filesize, sedm_filetype, sedm_filelink, sedm_modified_date, sedm_secd_id) 
VALUES ($1,$2,$3,$4,$5,$6,$7)
RETURNING sedm_id
`

type CreateSectionDetailMaterialParams struct {
	SedmID           int32     `db:"sedm_id" json:"sedmId"`
	SedmFilename     string    `db:"sedm_filename" json:"sedmFilename"`
	SedmFilesize     int32     `db:"sedm_filesize" json:"sedmFilesize"`
	SedmFiletype     string    `db:"sedm_filetype" json:"sedmFiletype"`
	SedmFilelink     string    `db:"sedm_filelink" json:"sedmFilelink"`
	SedmModifiedDate time.Time `db:"sedm_modified_date" json:"sedmModifiedDate"`
	SedmSecdID       int32     `db:"sedm_secd_id" json:"sedmSecdId"`
}

func (q *Queries) CreateSectionDetailMaterial(ctx context.Context, arg CreateSectionDetailMaterialParams) (int32, error) {
	row := q.db.QueryRowContext(ctx, createSectionDetailMaterial,
		arg.SedmID,
		arg.SedmFilename,
		arg.SedmFilesize,
		arg.SedmFiletype,
		arg.SedmFilelink,
		arg.SedmModifiedDate,
		arg.SedmSecdID,
	)
	var sedm_id int32
	err := row.Scan(&sedm_id)
	return sedm_id, err
}

const createSections = `-- name: CreateSections :one

INSERT INTO curriculum.sections (sect_title, sect_description, sect_total_section, sect_total_lecture, sect_total_minute, sect_modified_date, sect_prog_entity_id, sect_id) 
VALUES ($1,$2,$3,$4,$5,$6,$7,$8)
RETURNING sect_title
`

type CreateSectionsParams struct {
	SectTitle        string    `db:"sect_title" json:"sectTitle"`
	SectDescription  string    `db:"sect_description" json:"sectDescription"`
	SectTotalSection int32     `db:"sect_total_section" json:"sectTotalSection"`
	SectTotalLecture int32     `db:"sect_total_lecture" json:"sectTotalLecture"`
	SectTotalMinute  int32     `db:"sect_total_minute" json:"sectTotalMinute"`
	SectModifiedDate time.Time `db:"sect_modified_date" json:"sectModifiedDate"`
	SectProgEntityID int32     `db:"sect_prog_entity_id" json:"sectProgEntityId"`
	SectID           int32     `db:"sect_id" json:"sectId"`
}

func (q *Queries) CreateSections(ctx context.Context, arg CreateSectionsParams) (string, error) {
	row := q.db.QueryRowContext(ctx, createSections,
		arg.SectTitle,
		arg.SectDescription,
		arg.SectTotalSection,
		arg.SectTotalLecture,
		arg.SectTotalMinute,
		arg.SectModifiedDate,
		arg.SectProgEntityID,
		arg.SectID,
	)
	var sect_title string
	err := row.Scan(&sect_title)
	return sect_title, err
}

const deleteProgramEntity = `-- name: DeleteProgramEntity :exec
DELETE FROM curriculum.program_entity
WHERE prog_entity_id = $1
`

func (q *Queries) DeleteProgramEntity(ctx context.Context, progEntityID int32) error {
	_, err := q.db.ExecContext(ctx, deleteProgramEntity, progEntityID)
	return err
}

const deleteProgramEntityDescription = `-- name: DeleteProgramEntityDescription :exec
DELETE FROM curriculum.program_entity_description
WHERE pred_prog_entity_id = $1
`

func (q *Queries) DeleteProgramEntityDescription(ctx context.Context, predProgEntityID int32) error {
	_, err := q.db.ExecContext(ctx, deleteProgramEntityDescription, predProgEntityID)
	return err
}

const deleteProgramReviews = `-- name: DeleteProgramReviews :exec
DELETE FROM curriculum.program_reviews
WHERE prow_user_entity_id = $1
`

func (q *Queries) DeleteProgramReviews(ctx context.Context, prowUserEntityID int32) error {
	_, err := q.db.ExecContext(ctx, deleteProgramReviews, prowUserEntityID)
	return err
}

const deleteSectionDetail = `-- name: DeleteSectionDetail :exec
DELETE FROM curriculum.section_detail
WHERE secd_id = $1
`

func (q *Queries) DeleteSectionDetail(ctx context.Context, secdID int32) error {
	_, err := q.db.ExecContext(ctx, deleteSectionDetail, secdID)
	return err
}

const deleteSectionDetailMaterial = `-- name: DeleteSectionDetailMaterial :exec
DELETE FROM curriculum.section_detail_material
WHERE sedm_id = $1
`

func (q *Queries) DeleteSectionDetailMaterial(ctx context.Context, sedmID int32) error {
	_, err := q.db.ExecContext(ctx, deleteSectionDetailMaterial, sedmID)
	return err
}

const deleteSections = `-- name: DeleteSections :exec
DELETE FROM curriculum.sections
WHERE sect_title = $1
`

func (q *Queries) DeleteSections(ctx context.Context, sectTitle string) error {
	_, err := q.db.ExecContext(ctx, deleteSections, sectTitle)
	return err
}

const getProgramEntity = `-- name: GetProgramEntity :one

SELECT prog_entity_id, prog_title, prog_headline, prog_type, prog_learning_type, prog_rating, prog_total_trainee, prog_image, prog_best_seller, prog_price, prog_language, prog_modified_date, prog_duration, prog_duration_type, prog_tag_skill, prog_city_id, prog_cate_id, prog_created_by, prog_status FROM curriculum.program_entity
WHERE prog_entity_id = $1
`

// // curriculum.program_entity
// func (q *Queries) GetProgramEntity(ctx context.Context, progEntityID int32) (CurriculumProgramEntity, error) {
// 	row := q.db.QueryRowContext(ctx, getProgramEntity, progEntityID)
// 	var i CurriculumProgramEntity
// 	err := row.Scan(
// 		&i.ProgEntityID,
// 		&i.ProgTitle,
// 		&i.ProgHeadline,
// 		&i.ProgType,
// 		&i.ProgLearningType,
// 		&i.ProgRating,
// 		&i.ProgTotalTrainee,
// 		&i.ProgImage,
// 		&i.ProgBestSeller,
// 		&i.ProgPrice,
// 		&i.ProgLanguage,
// 		&i.ProgModifiedDate,
// 		&i.ProgDuration,
// 		&i.ProgDurationType,
// 		&i.ProgTagSkill,
// 		&i.ProgCityID,
// 		&i.ProgCateID,
// 		&i.ProgCreatedBy,
// 		&i.ProgStatus,
// 	)
// 	return i, err
// }

// const getProgramEntityDescription = `-- name: GetProgramEntityDescription :one

// SELECT pred_prog_entity_id, pred_item_learning, pred_item_include, pred_requirement, pred_description, pred_target_level FROM curriculum.program_entity_description
// WHERE pred_prog_entity_id = $1
// `

// // Program Entity Description
// func (q *Queries) GetProgramEntityDescription(ctx context.Context, predProgEntityID int32) (CurriculumProgramEntityDescription, error) {
// 	row := q.db.QueryRowContext(ctx, getProgramEntityDescription, predProgEntityID)
// 	var i CurriculumProgramEntityDescription
// 	err := row.Scan(
// 		&i.PredProgEntityID,
// 		&i.PredItemLearning,
// 		&i.PredItemInclude,
// 		&i.PredRequirement,
// 		&i.PredDescription,
// 		&i.PredTargetLevel,
// 	)
// 	return i, err
// }

// const getProgramReviews = `-- name: GetProgramReviews :one

// SELECT prow_user_entity_id, prow_prog_entity_id, prow_review, prow_rating, prow_modified_date FROM curriculum.program_reviews
// WHERE prow_user_entity_id = $1
// `

// // Program Reviews
// func (q *Queries) GetProgramReviews(ctx context.Context, prowUserEntityID int32) (CurriculumProgramReview, error) {
// 	row := q.db.QueryRowContext(ctx, getProgramReviews, prowUserEntityID)
// 	var i CurriculumProgramReview
// 	err := row.Scan(
// 		&i.ProwUserEntityID,
// 		&i.ProwProgEntityID,
// 		&i.ProwReview,
// 		&i.ProwRating,
// 		&i.ProwModifiedDate,
// 	)
// 	return i, err
// }

// const getSectionDetail = `-- name: GetSectionDetail :one

// SELECT secd_id, secd_title, secd_preview, secd_score, secd_note, secd_minute, secd_modified_date, secd_sect_id FROM curriculum.section_detail
// WHERE secd_id = $1
// `

// // Section Detail
// func (q *Queries) GetSectionDetail(ctx context.Context, secdID int32) (CurriculumSectionDetail, error) {
// 	row := q.db.QueryRowContext(ctx, getSectionDetail, secdID)
// 	var i CurriculumSectionDetail
// 	err := row.Scan(
// 		&i.SecdID,
// 		&i.SecdTitle,
// 		&i.SecdPreview,
// 		&i.SecdScore,
// 		&i.SecdNote,
// 		&i.SecdMinute,
// 		&i.SecdModifiedDate,
// 		&i.SecdSectID,
// 	)
// 	return i, err
// }

// const getSectionDetailMaterial = `-- name: GetSectionDetailMaterial :one

// SELECT sedm_id, sedm_filename, sedm_filesize, sedm_filetype, sedm_filelink, sedm_modified_date, sedm_secd_id FROM curriculum.section_detail_material
// WHERE sedm_id = $1
// `

// // Section Detail Material
// func (q *Queries) GetSectionDetailMaterial(ctx context.Context, sedmID int32) (CurriculumSectionDetailMaterial, error) {
// 	row := q.db.QueryRowContext(ctx, getSectionDetailMaterial, sedmID)
// 	var i CurriculumSectionDetailMaterial
// 	err := row.Scan(
// 		&i.SedmID,
// 		&i.SedmFilename,
// 		&i.SedmFilesize,
// 		&i.SedmFiletype,
// 		&i.SedmFilelink,
// 		&i.SedmModifiedDate,
// 		&i.SedmSecdID,
// 	)
// 	return i, err
// }

// const getSections = `-- name: GetSections :one

// SELECT sect_title, sect_description, sect_total_section, sect_total_lecture, sect_total_minute, sect_modified_date, sect_prog_entity_id, sect_id FROM curriculum.sections
// WHERE sect_title = $1
// `

// // Sections
// func (q *Queries) GetSections(ctx context.Context, sectTitle string) (CurriculumSection, error) {
// 	row := q.db.QueryRowContext(ctx, getSections, sectTitle)
// 	var i CurriculumSection
// 	err := row.Scan(
// 		&i.SectTitle,
// 		&i.SectDescription,
// 		&i.SectTotalSection,
// 		&i.SectTotalLecture,
// 		&i.SectTotalMinute,
// 		&i.SectModifiedDate,
// 		&i.SectProgEntityID,
// 		&i.SectID,
// 	)
// 	return i, err
// }

// const listProgramEntity = `-- name: ListProgramEntity :many
// SELECT prog_entity_id, prog_title, prog_headline, prog_type, prog_learning_type, prog_rating, prog_total_trainee, prog_image, prog_best_seller, prog_price, prog_language, prog_modified_date, prog_duration, prog_duration_type, prog_tag_skill, prog_city_id, prog_cate_id, prog_created_by, prog_status FROM curriculum.program_entity
// ORDER BY prog_title
// `

// func (q *Queries) ListProgramEntity(ctx context.Context) ([]CurriculumProgramEntity, error) {
// 	rows, err := q.db.QueryContext(ctx, listProgramEntity)
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer rows.Close()
// 	var items []CurriculumProgramEntity
// 	for rows.Next() {
// 		var i CurriculumProgramEntity
// 		if err := rows.Scan(
// 			&i.ProgEntityID,
// 			&i.ProgTitle,
// 			&i.ProgHeadline,
// 			&i.ProgType,
// 			&i.ProgLearningType,
// 			&i.ProgRating,
// 			&i.ProgTotalTrainee,
// 			&i.ProgImage,
// 			&i.ProgBestSeller,
// 			&i.ProgPrice,
// 			&i.ProgLanguage,
// 			&i.ProgModifiedDate,
// 			&i.ProgDuration,
// 			&i.ProgDurationType,
// 			&i.ProgTagSkill,
// 			&i.ProgCityID,
// 			&i.ProgCateID,
// 			&i.ProgCreatedBy,
// 			&i.ProgStatus,
// 		); err != nil {
// 			return nil, err
// 		}
// 		items = append(items, i)
// 	}
// 	if err := rows.Close(); err != nil {
// 		return nil, err
// 	}
// 	if err := rows.Err(); err != nil {
// 		return nil, err
// 	}
// 	return items, nil
// }

// const listProgramEntityDescription = `-- name: ListProgramEntityDescription :many
// SELECT pred_prog_entity_id, pred_item_learning, pred_item_include, pred_requirement, pred_description, pred_target_level FROM curriculum.program_entity_description
// ORDER BY pred_item_learning
// `

// func (q *Queries) ListProgramEntityDescription(ctx context.Context) ([]CurriculumProgramEntityDescription, error) {
// 	rows, err := q.db.QueryContext(ctx, listProgramEntityDescription)
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer rows.Close()
// 	var items []CurriculumProgramEntityDescription
// 	for rows.Next() {
// 		var i CurriculumProgramEntityDescription
// 		if err := rows.Scan(
// 			&i.PredProgEntityID,
// 			&i.PredItemLearning,
// 			&i.PredItemInclude,
// 			&i.PredRequirement,
// 			&i.PredDescription,
// 			&i.PredTargetLevel,
// 		); err != nil {
// 			return nil, err
// 		}
// 		items = append(items, i)
// 	}
// 	if err := rows.Close(); err != nil {
// 		return nil, err
// 	}
// 	if err := rows.Err(); err != nil {
// 		return nil, err
// 	}
// 	return items, nil
// }

// const listProgramReviews = `-- name: ListProgramReviews :many
// SELECT prow_user_entity_id, prow_prog_entity_id, prow_review, prow_rating, prow_modified_date FROM curriculum.program_reviews
// ORDER BY prow_prog_entity_id
// `

// func (q *Queries) ListProgramReviews(ctx context.Context) ([]CurriculumProgramReview, error) {
// 	rows, err := q.db.QueryContext(ctx, listProgramReviews)
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer rows.Close()
// 	var items []CurriculumProgramReview
// 	for rows.Next() {
// 		var i CurriculumProgramReview
// 		if err := rows.Scan(
// 			&i.ProwUserEntityID,
// 			&i.ProwProgEntityID,
// 			&i.ProwReview,
// 			&i.ProwRating,
// 			&i.ProwModifiedDate,
// 		); err != nil {
// 			return nil, err
// 		}
// 		items = append(items, i)
// 	}
// 	if err := rows.Close(); err != nil {
// 		return nil, err
// 	}
// 	if err := rows.Err(); err != nil {
// 		return nil, err
// 	}
// 	return items, nil
// }

// const listSectionDetail = `-- name: ListSectionDetail :many
// SELECT secd_id, secd_title, secd_preview, secd_score, secd_note, secd_minute, secd_modified_date, secd_sect_id FROM curriculum.section_detail
// ORDER BY secd_title
// `

// func (q *Queries) ListSectionDetail(ctx context.Context) ([]CurriculumSectionDetail, error) {
// 	rows, err := q.db.QueryContext(ctx, listSectionDetail)
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer rows.Close()
// 	var items []CurriculumSectionDetail
// 	for rows.Next() {
// 		var i CurriculumSectionDetail
// 		if err := rows.Scan(
// 			&i.SecdID,
// 			&i.SecdTitle,
// 			&i.SecdPreview,
// 			&i.SecdScore,
// 			&i.SecdNote,
// 			&i.SecdMinute,
// 			&i.SecdModifiedDate,
// 			&i.SecdSectID,
// 		); err != nil {
// 			return nil, err
// 		}
// 		items = append(items, i)
// 	}
// 	if err := rows.Close(); err != nil {
// 		return nil, err
// 	}
// 	if err := rows.Err(); err != nil {
// 		return nil, err
// 	}
// 	return items, nil
// }

// const listSectionDetailMaterial = `-- name: ListSectionDetailMaterial :many
// SELECT sedm_id, sedm_filename, sedm_filesize, sedm_filetype, sedm_filelink, sedm_modified_date, sedm_secd_id FROM curriculum.section_detail_material
// ORDER BY sedm_filename
// `

// func (q *Queries) ListSectionDetailMaterial(ctx context.Context) ([]CurriculumSectionDetailMaterial, error) {
// 	rows, err := q.db.QueryContext(ctx, listSectionDetailMaterial)
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer rows.Close()
// 	var items []CurriculumSectionDetailMaterial
// 	for rows.Next() {
// 		var i CurriculumSectionDetailMaterial
// 		if err := rows.Scan(
// 			&i.SedmID,
// 			&i.SedmFilename,
// 			&i.SedmFilesize,
// 			&i.SedmFiletype,
// 			&i.SedmFilelink,
// 			&i.SedmModifiedDate,
// 			&i.SedmSecdID,
// 		); err != nil {
// 			return nil, err
// 		}
// 		items = append(items, i)
// 	}
// 	if err := rows.Close(); err != nil {
// 		return nil, err
// 	}
// 	if err := rows.Err(); err != nil {
// 		return nil, err
// 	}
// 	return items, nil
// }

// const listSections = `-- name: ListSections :many
// SELECT sect_title, sect_description, sect_total_section, sect_total_lecture, sect_total_minute, sect_modified_date, sect_prog_entity_id, sect_id FROM curriculum.sections
// ORDER BY sect_description
// `

// func (q *Queries) ListSections(ctx context.Context) ([]CurriculumSection, error) {
// 	rows, err := q.db.QueryContext(ctx, listSections)
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer rows.Close()
// 	var items []CurriculumSection
// 	for rows.Next() {
// 		var i CurriculumSection
// 		if err := rows.Scan(
// 			&i.SectTitle,
// 			&i.SectDescription,
// 			&i.SectTotalSection,
// 			&i.SectTotalLecture,
// 			&i.SectTotalMinute,
// 			&i.SectModifiedDate,
// 			&i.SectProgEntityID,
// 			&i.SectID,
// 		); err != nil {
// 			return nil, err
// 		}
// 		items = append(items, i)
// 	}
// 	if err := rows.Close(); err != nil {
// 		return nil, err
// 	}
// 	if err := rows.Err(); err != nil {
// 		return nil, err
// 	}
// 	return items, nil
// }

// const updateProgramEntityDescription = `-- name: UpdateProgramEntityDescription :exec
// UPDATE curriculum.program_entity_description
//   set pred_item_learning = $2,
//   pred_item_include = $3,
//   pred_requirement = $4,
//   pred_description = $5,
//   pred_target_level = $6
// WHERE pred_prog_entity_id = $1
// `

// type UpdateProgramEntityDescriptionParams struct {
// 	PredProgEntityID int32          `db:"pred_prog_entity_id" json:"predProgEntityId"`
// 	PredItemLearning string `db:"pred_item_learning" json:"predItemLearning"`
// 	PredItemInclude  string `db:"pred_item_include" json:"predItemInclude"`
// 	PredRequirement  string `db:"pred_requirement" json:"predRequirement"`
// 	PredDescription  string `db:"pred_description" json:"predDescription"`
// 	PredTargetLevel  string `db:"pred_target_level" json:"predTargetLevel"`
// }

// func (q *Queries) UpdateProgramEntityDescription(ctx context.Context, arg UpdateProgramEntityDescriptionParams) error {
// 	_, err := q.db.ExecContext(ctx, updateProgramEntityDescription,
// 		arg.PredProgEntityID,
// 		arg.PredItemLearning,
// 		arg.PredItemInclude,
// 		arg.PredRequirement,
// 		arg.PredDescription,
// 		arg.PredTargetLevel,
// 	)
// 	return err
// }

// const updateProgramReviews = `-- name: UpdateProgramReviews :exec
// UPDATE curriculum.program_reviews
//   set prow_prog_entity_id = $2,
//   prow_review = $3,
//   prow_rating = $4,
//   prow_modified_date = $5
// WHERE prow_user_entity_id = $1
// `

// type UpdateProgramReviewsParams struct {
// 	ProwUserEntityID int32          `db:"prow_user_entity_id" json:"prowUserEntityId"`
// 	ProwProgEntityID int32          `db:"prow_prog_entity_id" json:"prowProgEntityId"`
// 	ProwReview       string `db:"prow_review" json:"prowReview"`
// 	ProwRating       int32  `db:"prow_rating" json:"prowRating"`
// 	ProwModifiedDate time.Time   `db:"prow_modified_date" json:"prowModifiedDate"`
// }

// func (q *Queries) UpdateProgramReviews(ctx context.Context, arg UpdateProgramReviewsParams) error {
// 	_, err := q.db.ExecContext(ctx, updateProgramReviews,
// 		arg.ProwUserEntityID,
// 		arg.ProwProgEntityID,
// 		arg.ProwReview,
// 		arg.ProwRating,
// 		arg.ProwModifiedDate,
// 	)
// 	return err
// }

// const updateSectionDetail = `-- name: UpdateSectionDetail :exec
// UPDATE curriculum.section_detail
//   set secd_title = $2,
//   secd_preview = $3,
//   secd_score = $4,
//   secd_note = $5,
//   secd_minute = $6,
//   secd_modified_date = $7,
//   secd_sect_id = $8
// WHERE secd_id = $1
// `

// type UpdateSectionDetailParams struct {
// 	SecdID           int32          `db:"secd_id" json:"secdId"`
// 	SecdTitle        string `db:"secd_title" json:"secdTitle"`
// 	SecdPreview      string `db:"secd_preview" json:"secdPreview"`
// 	SecdScore        int32  `db:"secd_score" json:"secdScore"`
// 	SecdNote         string `db:"secd_note" json:"secdNote"`
// 	SecdMinute       int32  `db:"secd_minute" json:"secdMinute"`
// 	SecdModifiedDate time.Time   `db:"secd_modified_date" json:"secdModifiedDate"`
// 	SecdSectID       int32  `db:"secd_sect_id" json:"secdSectId"`
// }

// func (q *Queries) UpdateSectionDetail(ctx context.Context, arg UpdateSectionDetailParams) error {
// 	_, err := q.db.ExecContext(ctx, updateSectionDetail,
// 		arg.SecdID,
// 		arg.SecdTitle,
// 		arg.SecdPreview,
// 		arg.SecdScore,
// 		arg.SecdNote,
// 		arg.SecdMinute,
// 		arg.SecdModifiedDate,
// 		arg.SecdSectID,
// 	)
// 	return err
// }

// const updateSectionDetailMaterial = `-- name: UpdateSectionDetailMaterial :exec
// UPDATE curriculum.section_detail_material
//   set sedm_filename = $2,
//   sedm_filesize = $3,
//   sedm_filetype = $4,
//   sedm_filelink = $5,
//   sedm_modified_date = $6,
//   sedm_secd_id = $7
// WHERE sedm_id = $1
// `

// type UpdateSectionDetailMaterialParams struct {
// 	SedmID           int32          `db:"sedm_id" json:"sedmId"`
// 	SedmFilename     string `db:"sedm_filename" json:"sedmFilename"`
// 	SedmFilesize     int32  `db:"sedm_filesize" json:"sedmFilesize"`
// 	SedmFiletype     string `db:"sedm_filetype" json:"sedmFiletype"`
// 	SedmFilelink     string `db:"sedm_filelink" json:"sedmFilelink"`
// 	SedmModifiedDate time.Time   `db:"sedm_modified_date" json:"sedmModifiedDate"`
// 	SedmSecdID       int32  `db:"sedm_secd_id" json:"sedmSecdId"`
// }

// func (q *Queries) UpdateSectionDetailMaterial(ctx context.Context, arg UpdateSectionDetailMaterialParams) error {
// 	_, err := q.db.ExecContext(ctx, updateSectionDetailMaterial,
// 		arg.SedmID,
// 		arg.SedmFilename,
// 		arg.SedmFilesize,
// 		arg.SedmFiletype,
// 		arg.SedmFilelink,
// 		arg.SedmModifiedDate,
// 		arg.SedmSecdID,
// 	)
// 	return err
// }

// const updateSections = `-- name: UpdateSections :exec
// UPDATE curriculum.sections
//   set sect_description = $2,
//   sect_total_section = $3,
//   sect_total_lecture = $4,
//   sect_total_minute = $5,
//   sect_modified_date = $6,
//   sect_prog_entity_id = $7,
//   sect_id = $8
// WHERE sect_title = $1
// `

// type UpdateSectionsParams struct {
// 	SectTitle        string `db:"sect_title" json:"sectTitle"`
// 	SectDescription  string `db:"sect_description" json:"sectDescription"`
// 	SectTotalSection int32  `db:"sect_total_section" json:"sectTotalSection"`
// 	SectTotalLecture int32  `db:"sect_total_lecture" json:"sectTotalLecture"`
// 	SectTotalMinute  int32  `db:"sect_total_minute" json:"sectTotalMinute"`
// 	SectModifiedDate time.Time   `db:"sect_modified_date" json:"sectModifiedDate"`
// 	SectProgEntityID int32          `db:"sect_prog_entity_id" json:"sectProgEntityId"`
// 	SectID           int32          `db:"sect_id" json:"sectId"`
// }

// func (q *Queries) UpdateSections(ctx context.Context, arg UpdateSectionsParams) error {
// 	_, err := q.db.ExecContext(ctx, updateSections,
// 		arg.SectTitle,
// 		arg.SectDescription,
// 		arg.SectTotalSection,
// 		arg.SectTotalLecture,
// 		arg.SectTotalMinute,
// 		arg.SectModifiedDate,
// 		arg.SectProgEntityID,
// 		arg.SectID,
// 	)
// 	return err
// }
