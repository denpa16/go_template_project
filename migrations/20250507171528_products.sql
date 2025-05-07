-- +goose Up
-- +goose StatementBegin

CREATE TABLE products
(
    id                BIGSERIAL                       PRIMARY KEY,
    name              varchar(250)                    NOT NULL,
    created_at        TIMESTAMP                       DEFAULT NOW() NOT NULL,
    updated_at        TIMESTAMP                       DEFAULT NOW() NOT NULL,
    deleted_at        TIMESTAMP                       NULL
);

CREATE INDEX  ix_products_id ON products (id);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS products CASCADE;
-- +goose StatementEnd
