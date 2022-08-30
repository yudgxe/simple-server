package storage

import (
	"database/sql"

	_ "github.com/lib/pq"
)

type Storage struct {
	config *Config
	db     *sql.DB
}

func New(config *Config) *Storage {
	return &Storage{
		config: config,
	}
}

func (s *Storage) Open() error {
	db, err := sql.Open("postgres", s.config.DatabaseURL)

	if err != nil {
		return err
	}

	if err := db.Ping(); err != nil {
		return err
	}

	s.db = db

	return nil
}

func (s *Storage) Close() {
	s.db.Close()
}

func (s *Storage) Order() *OrderRepo {
	return &OrderRepo{
		s: s,
	}
}

func (s *Storage) Payment() *PaymentRepo {
	return &PaymentRepo{
		s: s,
	}
}

func (s *Storage) Item() *ItemRepo {
	return &ItemRepo{
		s: s,
	}
}

func (s *Storage) Delivery() *DeliveryRepo {
	return &DeliveryRepo{
		s: s,
	}
}

func (s *Storage) Cache() *CacheRepo {
	return &CacheRepo{
		s: s,
	}
}
