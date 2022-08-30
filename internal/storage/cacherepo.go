package storage

import (
	"database/sql"

	"github.com/yudgxe/simple-server/internal/model"
)

type CacheRepo struct {
	s *Storage
}

func (r *CacheRepo) Create(c *model.Cache) error {
	var responce string

	if err := r.s.db.QueryRow(
		`INSERT INTO cache(order_uid) VALUES($1) RETURNING order_uid`,
		c.OrderID,
	).Scan(&responce); err != nil {
		return err
	}

	return nil
}

func (r *CacheRepo) Get() (*model.Cache, error) {
	c := &model.Cache{}

	if err := r.s.db.QueryRow(
		`SELECT * FROM cache`,
	).Scan(&c.OrderID); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return c, nil
}
