-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS favoritos(
	usuario_id INTEGER NOT NULL,
	negocio_id INTEGER NOT NULL,

	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
	updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
	deleted_at TIMESTAMP,

	PRIMARY KEY (usuario_id, negocio_id),
	FOREIGN KEY (usuario_id) REFERENCES usuarios(id),
	FOREIGN KEY (negocio_id) REFERENCES negocios(id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS favoritos;
-- +goose StatementEnd
