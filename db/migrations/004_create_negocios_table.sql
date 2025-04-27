-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS negocios(
	id SERIAL PRIMARY KEY,
	nombre VARCHAR(100) NOT NULL UNIQUE,
	descripcion TEXT NOT NULL,
	ubicacion JSONB NOT NULL,
	menu JSONB NULL,
	calificacion FLOAT(2) DEFAULT 0,
	tipo VARCHAR(20) NOT NULL,
	contacto VARCHAR(8) NULL,

	imagen_path TEXT NULL,

	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
	updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
	deleted_at TIMESTAMP NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS negocios;
-- +goose StatementEnd
