package repository

import (
	"database/sql"
	"day-19/model"
)

type RepositoryAdminDB struct {
	DB *sql.DB
}

func NewAdminRepository(db *sql.DB) RepositoryAdminDB {
	return RepositoryAdminDB{DB: db}
}

func (r *RepositoryAdminDB) Create(Admin *model.Admin) error {
	query := `INSERT INTO admin (name, username, password) VALUES ($1, $2, $3) RETURNING id`
	err := r.DB.QueryRow(query, Admin.AdminName, Admin.Username, Admin.Password).Scan(&Admin.ID)
	if err != nil {
		return err
	}

	return nil
}

func (r *RepositoryAdminDB) GetAdminLogin(admin model.Admin) (*model.Admin, error) {
	query := `SELECT id, name, username, password FROM admin WHERE username=$1 AND password=$2`
	var adminResponse model.Admin
	err := r.DB.QueryRow(query, admin.Username, admin.Password).Scan(&adminResponse.ID, &adminResponse.AdminName, &adminResponse.Username, &adminResponse.Password)

	if err != nil {
		return nil, err
	}

	return &adminResponse, nil
}
