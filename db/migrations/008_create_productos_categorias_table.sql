-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS productos_categorias(
	producto_id INTEGER NOT NULL,
	categoria_id INTEGER NOT NULL,

	PRIMARY KEY (producto_id, categoria_id),
	FOREIGN KEY (producto_id) REFERENCES productos(id),
	FOREIGN KEY (categoria_id) REFERENCES categorias(id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS productos_categorias;
-- +goose StatementEnd
