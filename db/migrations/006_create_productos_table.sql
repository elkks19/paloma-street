-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS productos(
	id SERIAL PRIMARY KEY,
	negocio_id INTEGER NOT NULL,

	nombre VARCHAR(100) NOT NULL,
	descripcion TEXT NOT NULL,
	precio FLOAT(2) NOT NULL,

	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
	updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
	deleted_at TIMESTAMP NULL,

	FOREIGN KEY (negocio_id) REFERENCES negocios(id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS productos;
-- +goose StatementEnd
