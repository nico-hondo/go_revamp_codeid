package models

import (
	"time"
)

type BootcampBatch struct {
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

type BootcampBatchTrainee struct {
	BatrID               int32     `db:"batr_id" json:"batrId"`
	BatrStatus           string    `db:"batr_status" json:"batrStatus"`
	BatrCertificated     string    `db:"batr_certificated" json:"batrCertificated"`
	BatreCertificateLink string    `db:"batre_certificate_link" json:"batreCertificateLink"`
	BatrAccessToken      string    `db:"batr_access_token" json:"batrAccessToken"`
	BatrAccessGrant      string    `db:"batr_access_grant" json:"batrAccessGrant"`
	BatrReview           string    `db:"batr_review" json:"batrReview"`
	BatrTotalScore       int32     `db:"batr_total_score" json:"batrTotalScore"`
	BatrModifiedDate     time.Time `db:"batr_modified_date" json:"batrModifiedDate"`
	BatrTraineeEntityID  int32     `db:"batr_trainee_entity_id" json:"batrTraineeEntityId"`
	BatrBatchID          int32     `db:"batr_batch_id" json:"batrBatchId"`
}

type BootcampBatchTraineeEvaluation struct {
	BtevID              int32     `db:"btev_id" json:"btevId"`
	BtevType            string    `db:"btev_type" json:"btevType"`
	BtevHeader          string    `db:"btev_header" json:"btevHeader"`
	BtevSection         string    `db:"btev_section" json:"btevSection"`
	BtevSkill           string    `db:"btev_skill" json:"btevSkill"`
	BtevWeek            int32     `db:"btev_week" json:"btevWeek"`
	BtevSkor            int32     `db:"btev_skor" json:"btevSkor"`
	BtevNote            string    `db:"btev_note" json:"btevNote"`
	BtevModifiedDate    time.Time `db:"btev_modified_date" json:"btevModifiedDate"`
	BtevBatchID         int32     `db:"btev_batch_id" json:"btevBatchId"`
	BtevTraineeEntityID int32     `db:"btev_trainee_entity_id" json:"btevTraineeEntityId"`
}

type BootcampInstructorProgram struct {
	BatchID           int32     `db:"batch_id" json:"batchId"`
	InproEntityID     int32     `db:"inpro_entity_id" json:"inproEntityId"`
	InproEmpEntityID  int32     `db:"inpro_emp_entity_id" json:"inproEmpEntityId"`
	InproModifiedDate time.Time `db:"inpro_modified_date" json:"inproModifiedDate"`
}

type BootcampProgramApply struct {
	PrapUserEntityID int32     `db:"prap_user_entity_id" json:"prapUserEntityId"`
	PrapProgEntityID int32     `db:"prap_prog_entity_id" json:"prapProgEntityId"`
	PrapTestScore    int32     `db:"prap_test_score" json:"prapTestScore"`
	PrapGpa          int32     `db:"prap_gpa" json:"prapGpa"`
	PrapIqTest       int32     `db:"prap_iq_test" json:"prapIqTest"`
	PrapReview       string    `db:"prap_review" json:"prapReview"`
	PrapModifiedDate time.Time `db:"prap_modified_date" json:"prapModifiedDate"`
	PrapStatus       string    `db:"prap_status" json:"prapStatus"`
}

type BootcampProgramApplyProgress struct {
	ParogID           int32     `db:"parog_id" json:"parogId"`
	ParogUserEntityID int32     `db:"parog_user_entity_id" json:"parogUserEntityId"`
	ParogProgEntityID int32     `db:"parog_prog_entity_id" json:"parogProgEntityId"`
	ParogActionDate   time.Time `db:"parog_action_date" json:"parogActionDate"`
	ParogModifiedDate time.Time `db:"parog_modified_date" json:"parogModifiedDate"`
	ParogComment      string    `db:"parog_comment" json:"parogComment"`
	ParogProgressName string    `db:"parog_progress_name" json:"parogProgressName"`
	ParogEmpEntityID  int32     `db:"parog_emp_entity_id" json:"parogEmpEntityId"`
	ParogStatus       string    `db:"parog_status" json:"parogStatus"`
}
