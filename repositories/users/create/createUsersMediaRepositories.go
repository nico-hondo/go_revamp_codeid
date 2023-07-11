package users

import (
	"context"
	"time"
)

const createMedia = `-- name: CreateMedia :one

INSERT INTO users.users_media
(usme_id, usme_entity_id, usme_file_link, usme_filename,
usme_filesize, usme_filetype, usme_note, usme_modified_date)
VALUES($1,$2,$3,$4,$5,$6,$7,$8)
RETURNING usme_id
`

type CreateMediaParams struct {
	UsmeID           int32          `db:"usme_id" json:"usmeId"`
	UsmeEntityID     int32          `db:"usme_entity_id" json:"usmeEntityId"`
	UsmeFileLink     string `db:"usme_file_link" json:"usmeFileLink"`
	UsmeFilename     string `db:"usme_filename" json:"usmeFilename"`
	UsmeFilesize     int32  `db:"usme_filesize" json:"usmeFilesize"`
	UsmeFiletype     string `db:"usme_filetype" json:"usmeFiletype"`
	UsmeNote         string `db:"usme_note" json:"usmeNote"`
	UsmeModifiedDate time.Time   `db:"usme_modified_date" json:"usmeModifiedDate"`
}

func (q *Queries) CreateMedia(ctx context.Context, arg CreateMediaParams) (int32, error) {
	row := q.db.QueryRowContext(ctx, createMedia,
		arg.UsmeID,
		arg.UsmeEntityID,
		arg.UsmeFileLink,
		arg.UsmeFilename,
		arg.UsmeFilesize,
		arg.UsmeFiletype,
		arg.UsmeNote,
		arg.UsmeModifiedDate,
	)
	var usme_id int32
	err := row.Scan(&usme_id)
	return usme_id, err
}