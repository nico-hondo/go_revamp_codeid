-- name: Getcart_items :one
SELECT * FROM sales.cart_items
WHERE cait_id = $1;

-- name: ListCart_item :many
SELECT * FROM sales.cart_items
ORDER BY cait_quantity;


-- name: CreateCartItems :one
INSERT INTO sales.cart_items(cait_id, cait_quantity, cait_unit_price, cait_modified_date,cait_user_entity_id,cait_prog_entity_id) VALUES ($1, $2, $3, $4, $5, $6)
RETURNING cait_id;

-- name: DeleteCart_item :exec
DELETE FROM sales.cart_items
WHERE cait_id = $1;

-- name: UpdateCart_item :exec
UPDATE sales.cart_items
  set cait_quantity = $2,
  cait_unit_price = $3
WHERE cait_id = $1;

-- name: Getsales_order_detail :one
SELECT * FROM sales.sales_order_detail
WHERE sode_id = $1;
-- name: ListSales_order_detail :many
SELECT * FROM sales.sales_order_detail
ORDER BY sode_qty;

-- name: CreateSales_order_detail :one
INSERT INTO sales.sales_order_detail(sode_id, sode_qty, sode_unit_price, sode_unit_discount,sode_line_total,sode_modified_date,sode_sohe_id,sode_prog_entity_id) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
RETURNING sode_id;

-- name: DeleteSales_order_detail :exec
DELETE FROM sales.sales_order_detail
WHERE sode_id = $1;

-- name: UpdateSales_order_detail :exec
UPDATE sales.sales_order_detail
  set sode_qty = $2,
  sode_unit_price = $3
WHERE sode_id = $1;


-- name: Getsales_order_header :one
SELECT * FROM sales.sales_order_header
WHERE sohe_id = $1;

-- name: ListSales_order_header :many
SELECT * FROM sales.sales_order_header
ORDER BY  sohe_order_date;

-- name: CreateSales_order_header :one
INSERT INTO sales.sales_order_header(sohe_id, sohe_order_date, sohe_due_date, sohe_ship_date,sohe_order_number,sohe_account_number,sohe_trpa_code_number,sohe_subtotal,sohe_tax,sohe_total_due,sohe_license_code,sohe_modified_date,sohe_user_entity_id,sohe_status) 
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14)
RETURNING sohe_id;

-- name: DeleteSales_order_header :exec
DELETE FROM sales.sales_order_header
WHERE sohe_id = $1;

-- name: UpdateSales_order_header :exec
UPDATE sales.sales_order_header
  set sohe_order_date = $2,
  sohe_due_date = $3
WHERE sohe_id = $1;


-- name: GetcSpecial_offer :one
SELECT * FROM sales.special_offer
WHERE spof_id = $1;

-- name: ListSpecial_offer :many
SELECT * FROM sales.special_offer
ORDER BY  spof_description;

-- name: CreateSales_special_offer :one
INSERT INTO sales.special_offer(spof_id, spof_description, spof_discount, spof_type,spof_start_date,spof_end_date,spof_min_qty,spof_max_qty,spof_modified_date,spof_cate_id) 
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
RETURNING spof_id;

-- name: DeleteSales_special_offer :exec
DELETE FROM sales.special_offer
WHERE spof_id = $1;

-- name: UpdateSpecial_offer :exec
UPDATE sales.special_offer
  set spof_description = $2,
  spof_discount = $3
WHERE spof_id = $1;


-- name: GetcSpecial_offer_programs :one
SELECT * FROM sales.special_offer_programs
WHERE soco_id = $1;

-- name: ListSpecial_offer_programs :many
SELECT * FROM sales.special_offer_programs
ORDER BY soco_spof_id;

-- name: CreateSpecial_offer_programs :one
INSERT INTO sales.special_offer_programs(soco_id, soco_spof_id, soco_prog_entity_id, soco_status,soco_modified_date) VALUES ($1, $2, $3, $4, $5)
RETURNING soco_id;

-- name: DeleteSpecial_offer_programs :exec
DELETE FROM sales.special_offer_programs
WHERE soco_id = $1;

-- name: UpdateSpecial_offer_programs :exec
UPDATE sales.special_offer_programs
  set soco_spof_id = $2,
  soco_prog_entity_id = $3
WHERE soco_id = $1;