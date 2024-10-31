package service

import (
	"day-19/model"
	"day-19/repository"
	"errors"
	"fmt"
)

type AdminService struct {
	RepoAdmin repository.RepositoryAdminDB
}

func NewAdminService(repo repository.RepositoryAdminDB) *AdminService {
	return &AdminService{RepoAdmin: repo}
}

func (cs *AdminService) CreateDataAdmin(name string, username string, password string) error {
	if name == "" {
		return errors.New("username tidak boleh kosong")
	}

	admin := model.Admin{
		AdminName: name,
		Username:  username,
		Password:  password,
	}

	err := cs.RepoAdmin.Create(&admin)
	if err != nil {
		fmt.Println("Error :", err)
	}

	fmt.Println("berhasil input data Admin dengan id ", admin.ID)

	return nil
}

func (cs *AdminService) LoginService(user model.Admin) (*model.Response, error) {

	admin, err := cs.RepoAdmin.GetAdminLogin(user)

	if err != nil {
		return nil, err
	}
	response := model.Response{
		StatusCode: 200,
		Message:    "login success",
		Data:       *admin,
	}
	return &response, nil
}
