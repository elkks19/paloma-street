-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS producto_imagenes(
	id SERIAL PRIMARY KEY,
	producto_id INTEGER NOT NULL,
	imagen_url TEXT NOT NULL,

	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
	updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
	deleted_at TIMESTAMP NULL,

	FOREIGN KEY (producto_id) REFERENCES productos(id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS producto_imagenes;
-- +goose StatementEnd
