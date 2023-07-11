package dbContext

import (
	"context"
	"time"
)

type CreateSales_special_offerParams struct {
	SpofID           int32     `db:"spof_id" json:"spofId"`
	SpofDescription  string    `db:"spof_description" json:"spofDescription"`
	SpofDiscount     int32     `db:"spof_discount" json:"spofDiscount"`
	SpofType         string    `db:"spof_type" json:"spofType"`
	SpofStartDate    time.Time `db:"spof_start_date" json:"spofStartDate"`
	SpofEndDate      time.Time `db:"spof_end_date" json:"spofEndDate"`
	SpofMinQty       int32     `db:"spof_min_qty" json:"spofMinQty"`
	SpofMaxQty       int32     `db:"spof_max_qty" json:"spofMaxQty"`
	SpofModifiedDate time.Time `db:"spof_modified_date" json:"spofModifiedDate"`
	SpofCateID       int32     `db:"spof_cate_id" json:"spofCateId"`
}

const getcSpecial_offer = `-- name: GetcSpecial_offer :one
SELECT spof_id, spof_description, spof_discount, spof_type, spof_start_date, spof_end_date, spof_min_qty, spof_max_qty, spof_modified_date, spof_cate_id FROM sales.special_offer
WHERE spof_id = $1
`

func (q *Queries) GetcSpecial_offer(ctx context.Context, spofID int32) (SalesSpecialOffer, error) {
	row := q.db.QueryRowContext(ctx, getcSpecial_offer, spofID)
	var i SalesSpecialOffer
	err := row.Scan(
		&i.SpofID,
		&i.SpofDescription,
		&i.SpofDiscount,
		&i.SpofType,
		&i.SpofStartDate,
		&i.SpofEndDate,
		&i.SpofMinQty,
		&i.SpofMaxQty,
		&i.SpofModifiedDate,
		&i.SpofCateID,
	)
	return i, err
}
