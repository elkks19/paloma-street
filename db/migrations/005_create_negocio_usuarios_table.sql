-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS negocio_usuarios(
	usuario_id INTEGER NOT NULL,
	negocio_id INTEGER NOT NULL,

	PRIMARY KEY(usuario_id, negocio_id),

	FOREIGN KEY(usuario_id) REFERENCES usuarios(id),
	FOREIGN KEY(negocio_id) REFERENCES negocios(id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS negocio_usuarios;
-- +goose StatementEnd
