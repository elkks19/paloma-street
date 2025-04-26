package db

import (
	"context"

	"github.com/uptrace/bun"
)

func TX(ctx context.Context, tx bun.Tx, fn func(context.Context) error) error {
	err := fn(ctx)
	if err != nil {
		err = tx.Rollback()
		if err != nil {
			return err
		}

		return err
	}

	return tx.Commit()
}
