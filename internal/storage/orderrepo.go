package storage

import (
	"database/sql"

	"github.com/yudgxe/simple-server/internal/model"
)

type OrderRepo struct {
	s *Storage
}

func (r *OrderRepo) Create(o *model.Order) error {
	var responce string

	if err := r.s.db.QueryRow(
		`INSERT INTO orders(
			order_uid, 
			track_number, 
			entry, locale, 
			internal_signature, 
			customer_id, 
			delivery_service, 
			shardkey, 
			sm_id, 
			date_created, 
			oof_shard
		) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11) RETURNING order_uid`,
		o.UID,
		o.TrackNumber,
		o.Entry,
		o.Locale,
		o.Signature,
		o.CustomerID,
		o.DeliveryService,
		o.Shardkey,
		o.SMID,
		o.DateCreated,
		o.OffShard,
	).Scan(&responce); err != nil {
		return err
	}

	return nil
}

func (r *OrderRepo) FindByID(ID *string) (*model.Order, error) {
	o := &model.Order{}

	if err := r.s.db.QueryRow(
		`SELECT * FROM orders WHERE order_uid = $1`,
		ID,
	).Scan(
		&o.UID,
		&o.TrackNumber,
		&o.Entry,
		&o.Locale,
		&o.Signature,
		&o.CustomerID,
		&o.DeliveryService,
		&o.Shardkey,
		&o.SMID,
		&o.DateCreated,
		&o.OffShard,
	); err != nil {
		return nil, err
	}

	return o, nil
}

func (r *OrderRepo) GetLastRecordID() (*string, error) {
	var UID string

	if err := r.s.db.QueryRow(`SELECT order_uid FROM orders ORDER BY date_created DESC LIMIT 1`).Scan(&UID); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}

		return nil, err
	}

	return &UID, nil
}
