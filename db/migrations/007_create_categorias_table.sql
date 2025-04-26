-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS categorias(
	id SERIAL PRIMARY KEY,
	nombre VARCHAR(100) NOT NULL UNIQUE,

	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS categorias;
-- +goose StatementEnd
