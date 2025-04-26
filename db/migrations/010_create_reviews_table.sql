-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS reviews(
	id SERIAL PRIMARY KEY,
	usuario_id INTEGER NOT NULL,
	producto_id INTEGER NOT NULL,

	contenido TEXT NULL,
	calificacion INTEGER NOT NULL,

	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
	updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
	deleted_at TIMESTAMP,

	FOREIGN KEY (usuario_id) REFERENCES usuarios(id),
	FOREIGN KEY (producto_id) REFERENCES productos(id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS reviews;
-- +goose StatementEnd
