package repository

import (
	"database/sql"
	"day-19/model"
)

// type RepositoryAdminInterface interface {
// 	Create(Admin *model.Admin) error
// }

type RepositoryAdminDB struct {
	DB *sql.DB
}

func NewAdminRepository(db *sql.DB) RepositoryAdminDB {
	return RepositoryAdminDB{DB: db}
}

func (r *RepositoryAdminDB) Create(Admin *model.Admin) error {
	query := `INSERT INTO admin (name) VALUES ($1) RETURNING id`
	err := r.DB.QueryRow(query, Admin.AdminName).Scan(&Admin.ID)
	if err != nil {
		return err
	}

	return nil
}
