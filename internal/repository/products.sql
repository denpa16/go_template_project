-- PRODUCTS

-- name: SqlcGetProducts :many
SELECT id, created_at, updated_at, deleted_at, name
FROM products
LIMIT $1
OFFSET $2;
