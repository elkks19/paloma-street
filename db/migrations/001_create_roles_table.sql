-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS roles(
	id SERIAL PRIMARY KEY,
	nombre VARCHAR(50) NOT NULL
);

CREATE TABLE IF NOT EXISTS permisos(
	id SERIAL PRIMARY KEY,
	nombre VARCHAR(50) NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS permisos;
DROP TABLE IF EXISTS roles;
-- +goose StatementEnd
