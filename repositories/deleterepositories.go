package repositories

import (
	"context"
)

const deleteCart_item = `-- name: DeleteCart_item :exec
DELETE FROM sales.cart_items
WHERE cait_id = $1
`

func (q *Queries) DeleteCart_item(ctx context.Context, caitID int32) error {
	_, err := q.db.ExecContext(ctx, deleteCart_item, caitID)
	return err
}

const deleteSales_order_detail = `-- name: DeleteSales_order_detail :exec
DELETE FROM sales.sales_order_detail
WHERE sode_id = $1
`

func (q *Queries) DeleteSales_order_detail(ctx context.Context, sodeID int32) error {
	_, err := q.db.ExecContext(ctx, deleteSales_order_detail, sodeID)
	return err
}

const deleteSales_order_header = `-- name: DeleteSales_order_header :exec
DELETE FROM sales.sales_order_header
WHERE sohe_id = $1
`

func (q *Queries) DeleteSales_order_header(ctx context.Context, soheID int32) error {
	_, err := q.db.ExecContext(ctx, deleteSales_order_header, soheID)
	return err
}

const deleteSales_special_offer = `-- name: DeleteSales_special_offer :exec
DELETE FROM sales.special_offer
WHERE spof_id = $1
`

func (q *Queries) DeleteSales_special_offer(ctx context.Context, spofID int32) error {
	_, err := q.db.ExecContext(ctx, deleteSales_special_offer, spofID)
	return err
}
