package repositories

import (
	"context"
	"time"
)

const createCartItems = `-- name: CreateCartItems :one
INSERT INTO sales.cart_items(cait_id, cait_quantity, cait_unit_price, cait_modified_date,cait_user_entity_id,cait_prog_entity_id) VALUES ($1, $2, $3, $4, $5, $6)
RETURNING cait_id
`

type CreateCartItemsParams struct {
	CaitID           int32     `db:"cait_id" json:"caitId"`
	CaitQuantity     int32     `db:"cait_quantity" json:"caitQuantity"`
	CaitUnitPrice    int32     `db:"cait_unit_price" json:"caitUnitPrice"`
	CaitModifiedDate time.Time `db:"cait_modified_date" json:"caitModifiedDate"`
	CaitUserEntityID int32     `db:"cait_user_entity_id" json:"caitUserEntityId"`
	CaitProgEntityID int32     `db:"cait_prog_entity_id" json:"caitProgEntityId"`
}

func (q *Queries) CreateCartItems(ctx context.Context, arg CreateCartItemsParams) (int32, error) {
	row := q.db.QueryRowContext(ctx, createCartItems,
		arg.CaitID,
		arg.CaitQuantity,
		arg.CaitUnitPrice,
		arg.CaitModifiedDate,
		arg.CaitUserEntityID,
		arg.CaitProgEntityID,
	)
	var cait_id int32
	err := row.Scan(&cait_id)
	return cait_id, err
}

const createSales_order_detail = `-- name: CreateSales_order_detail :one
INSERT INTO sales.sales_order_detail(sode_id, sode_qty, sode_unit_price, sode_unit_discount,sode_line_total,sode_modified_date,sode_sohe_id,sode_prog_entity_id) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
RETURNING sode_id
`

type CreateSales_order_detailParams struct {
	SodeID           int32     `db:"sode_id" json:"sodeId"`
	SodeQty          int32     `db:"sode_qty" json:"sodeQty"`
	SodeUnitPrice    int32     `db:"sode_unit_price" json:"sodeUnitPrice"`
	SodeUnitDiscount int32     `db:"sode_unit_discount" json:"sodeUnitDiscount"`
	SodeLineTotal    int32     `db:"sode_line_total" json:"sodeLineTotal"`
	SodeModifiedDate time.Time `db:"sode_modified_date" json:"sodeModifiedDate"`
	SodeSoheID       int32     `db:"sode_sohe_id" json:"sodeSoheId"`
	SodeProgEntityID int32     `db:"sode_prog_entity_id" json:"sodeProgEntityId"`
}

func (q *Queries) CreateSales_order_detail(ctx context.Context, arg CreateSales_order_detailParams) (int32, error) {
	row := q.db.QueryRowContext(ctx, createSales_order_detail,
		arg.SodeID,
		arg.SodeQty,
		arg.SodeUnitPrice,
		arg.SodeUnitDiscount,
		arg.SodeLineTotal,
		arg.SodeModifiedDate,
		arg.SodeSoheID,
		arg.SodeProgEntityID,
	)
	var sode_id int32
	err := row.Scan(&sode_id)
	return sode_id, err
}

const createSales_order_header = `-- name: CreateSales_order_header :one
INSERT INTO sales.sales_order_header(sohe_id, sohe_order_date, sohe_due_date, sohe_ship_date,sohe_order_number,sohe_account_number,sohe_trpa_code_number,sohe_subtotal,sohe_tax,sohe_total_due,sohe_license_code,sohe_modified_date,sohe_user_entity_id,sohe_status) 
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14)
RETURNING sohe_id
`

type CreateSales_order_headerParams struct {
	SoheID             int32     `db:"sohe_id" json:"soheId"`
	SoheOrderDate      time.Time `db:"sohe_order_date" json:"soheOrderDate"`
	SoheDueDate        time.Time `db:"sohe_due_date" json:"soheDueDate"`
	SoheShipDate       time.Time `db:"sohe_ship_date" json:"soheShipDate"`
	SoheOrderNumber    string    `db:"sohe_order_number" json:"soheOrderNumber"`
	SoheAccountNumber  string    `db:"sohe_account_number" json:"soheAccountNumber"`
	SoheTrpaCodeNumber string    `db:"sohe_trpa_code_number" json:"soheTrpaCodeNumber"`
	SoheSubtotal       string    `db:"sohe_subtotal" json:"soheSubtotal"`
	SoheTax            string    `db:"sohe_tax" json:"soheTax"`
	SoheTotalDue       int32     `db:"sohe_total_due" json:"soheTotalDue"`
	SoheLicenseCode    string    `db:"sohe_license_code" json:"soheLicenseCode"`
	SoheModifiedDate   time.Time `db:"sohe_modified_date" json:"soheModifiedDate"`
	SoheUserEntityID   int32     `db:"sohe_user_entity_id" json:"soheUserEntityId"`
	SoheStatus         string    `db:"sohe_status" json:"soheStatus"`
}

func (q *Queries) CreateSales_order_header(ctx context.Context, arg CreateSales_order_headerParams) (int32, error) {
	row := q.db.QueryRowContext(ctx, createSales_order_header,
		arg.SoheID,
		arg.SoheOrderDate,
		arg.SoheDueDate,
		arg.SoheShipDate,
		arg.SoheOrderNumber,
		arg.SoheAccountNumber,
		arg.SoheTrpaCodeNumber,
		arg.SoheSubtotal,
		arg.SoheTax,
		arg.SoheTotalDue,
		arg.SoheLicenseCode,
		arg.SoheModifiedDate,
		arg.SoheUserEntityID,
		arg.SoheStatus,
	)
	var sohe_id int32
	err := row.Scan(&sohe_id)
	return sohe_id, err
}

const createSales_special_offer = `-- name: CreateSales_special_offer :one
INSERT INTO sales.special_offer(spof_id, spof_description, spof_discount, spof_type,spof_start_date,spof_end_date,spof_min_qty,spof_max_qty,spof_modified_date,spof_cate_id) 
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
RETURNING spof_id
`

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

func (q *Queries) CreateSales_special_offer(ctx context.Context, arg CreateSales_special_offerParams) (int32, error) {
	row := q.db.QueryRowContext(ctx, createSales_special_offer,
		arg.SpofID,
		arg.SpofDescription,
		arg.SpofDiscount,
		arg.SpofType,
		arg.SpofStartDate,
		arg.SpofEndDate,
		arg.SpofMinQty,
		arg.SpofMaxQty,
		arg.SpofModifiedDate,
		arg.SpofCateID,
	)
	var spof_id int32
	err := row.Scan(&spof_id)
	return spof_id, err
}

const createSpecial_offer_programs = `-- name: CreateSpecial_offer_programs :one
INSERT INTO sales.special_offer_programs(soco_id, soco_spof_id, soco_prog_entity_id, soco_status,soco_modified_date) VALUES ($1, $2, $3, $4, $5)
RETURNING soco_id
`

type CreateSpecial_offer_programsParams struct {
	SocoID           int32     `db:"soco_id" json:"socoId"`
	SocoSpofID       int32     `db:"soco_spof_id" json:"socoSpofId"`
	SocoProgEntityID int32     `db:"soco_prog_entity_id" json:"socoProgEntityId"`
	SocoStatus       string    `db:"soco_status" json:"socoStatus"`
	SocoModifiedDate time.Time `db:"soco_modified_date" json:"socoModifiedDate"`
}

func (q *Queries) CreateSpecial_offer_programs(ctx context.Context, arg CreateSpecial_offer_programsParams) (int32, error) {
	row := q.db.QueryRowContext(ctx, createSpecial_offer_programs,
		arg.SocoID,
		arg.SocoSpofID,
		arg.SocoProgEntityID,
		arg.SocoStatus,
		arg.SocoModifiedDate,
	)
	var soco_id int32
	err := row.Scan(&soco_id)
	return soco_id, err
}
