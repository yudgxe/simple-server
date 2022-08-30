package storage

import (
	"github.com/yudgxe/simple-server/internal/model"
)

type ItemRepo struct {
	s *Storage
}

func (r *ItemRepo) Create(items []*model.Item) error {
	var responce string

	for _, item := range items {
		if err := r.s.db.QueryRow(
			`INSERT INTO items(
				chrt_id,
				track_number,
				price,
				rid,
				name,
				sale,
				size,
				total_price,
				nm_id,
				brand,
				status
			) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11) RETURNING chrt_id`,
			item.ChrtID,
			item.TrackNumber,
			item.Price,
			item.Rid,
			item.Name,
			item.Sale,
			item.Size,
			item.TotalPrice,
			item.NMID,
			item.Brand,
			item.Status,
		).Scan(&responce); err != nil {
			return err
		}
	}

	return nil
}

func (r *ItemRepo) FindByTrackNumber(track *string) ([]*model.Item, error) {
	items := []*model.Item{}

	rows, err := r.s.db.Query("SELECT * FROM items WHERE track_number = $1", track)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		item := &model.Item{}
		if err := rows.Scan(
			&item.ChrtID,
			&item.TrackNumber,
			&item.Price,
			&item.Rid,
			&item.Name,
			&item.Sale,
			&item.Size,
			&item.TotalPrice,
			&item.NMID,
			&item.Brand,
			&item.Status,
		); err != nil {
			return nil, err
		}

		items = append(items, item)
	}

	return items, nil
}
