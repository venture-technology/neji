package repository

import "database/sql"

type INejiRepository interface {
}

type NejiRepository struct {
	db *sql.DB
}

func NewNejiRepository(db *sql.DB) *NejiRepository {
	return &NejiRepository{
		db: db,
	}
}
