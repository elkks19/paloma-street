-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS sessions (
	id SERIAL PRIMARY KEY,
	usuario_id INTEGER,
	key BYTEA,
	data BYTEA,
	token TEXT,
	created_at  TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
	last_access TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
	expires_at  TIMESTAMP,

	FOREIGN KEY (usuario_id) REFERENCES usuarios(id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS sessions;
-- +goose StatementEnd
