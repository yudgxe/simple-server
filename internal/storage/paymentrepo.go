package storage

import "github.com/yudgxe/simple-server/internal/model"

type PaymentRepo struct {
	s *Storage
}

func (r *PaymentRepo) Create(p *model.Payment) error {
	var responce string

	if err := r.s.db.QueryRow(
		`INSERT INTO payment(
			transaction,
			request_id,
			currency,
			provider,
			amount,
			payment_dt,
			bank,
			delivery_cost,
			goods_total,
			custom_fee
		) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10) RETURNING transaction`,
		p.Transaction,
		p.RequestID,
		p.Currency,
		p.Provider,
		p.Amount,
		p.PaymentDt,
		p.Bank,
		p.DeliveryCost,
		p.GoodsTotal,
		p.CustomFee,
	).Scan(&responce); err != nil {
		return err
	}

	return nil
}

func (r *PaymentRepo) FindByTransaction(t *string) (*model.Payment, error) {
	p := &model.Payment{}

	if err := r.s.db.QueryRow(
		`SELECT * FROM payment WHERE transaction = $1`,
		t,
	).Scan(
		&p.Transaction,
		&p.RequestID,
		&p.Currency,
		&p.Provider,
		&p.Amount,
		&p.PaymentDt,
		&p.Bank,
		&p.DeliveryCost,
		&p.GoodsTotal,
		&p.CustomFee,
	); err != nil {
		return nil, err
	}

	return p, nil
}
