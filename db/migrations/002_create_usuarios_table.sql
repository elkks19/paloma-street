-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS usuarios(
	id SERIAL PRIMARY KEY,
	rol_id INTEGER NOT NULL,
	nombre VARCHAR(255) NOT NULL,
	email VARCHAR(255) NOT NULL UNIQUE,
	email_verificado TIMESTAMP NULL,
	password VARCHAR(255) NOT NULL,

	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
	updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
	deleted_at TIMESTAMP,

	FOREIGN KEY (rol_id) REFERENCES roles(id)
);

CREATE TABLE IF NOT EXISTS usuario_has_permiso(
	usuario_id INTEGER NOT NULL,
	permiso_id INTEGER NOT NULL,

	PRIMARY KEY(usuario_id, permiso_id),
	FOREIGN KEY(usuario_id) REFERENCES usuarios(id),
	FOREIGN KEY(permiso_id) REFERENCES permisos(id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS usuario_has_permiso;
DROP TABLE IF EXISTS usuarios;
-- +goose StatementEnd

