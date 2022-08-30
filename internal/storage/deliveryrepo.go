package storage

import "github.com/yudgxe/simple-server/internal/model"

type DeliveryRepo struct {
	s *Storage
}

func (r *DeliveryRepo) Create(d *model.Delivery, orderID *string) error {
	var responce string

	if err := r.s.db.QueryRow(
		`INSERT INTO delivery(
			order_uid,
			name,
			phone,
			zip,
			city,
			address,
			region,
			email
		) VALUES ($1, $2, $3, $4, $5, $6, $7, $8) RETURNING order_uid`,
		orderID,
		d.Name,
		d.Phone,
		d.Zip,
		d.City,
		d.Address,
		d.Region,
		d.Email,
	).Scan(&responce); err != nil {
		return err
	}

	return nil
}

func (r *DeliveryRepo) FindByOrderID(ID *string) (*model.Delivery, error) {
	d := &model.Delivery{}
	var orderID string

	if err := r.s.db.QueryRow(
		`SELECT * FROM delivery WHERE order_uid = $1`,
		ID,
	).Scan(
		&orderID,
		&d.Name,
		&d.Phone,
		&d.Zip,
		&d.City,
		&d.Address,
		&d.Region,
		&d.Email,
	); err != nil {
		return nil, err
	}

	return d, nil
}
