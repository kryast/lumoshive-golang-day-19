package handler

import (
	"database/sql"
	"day-19/model"
	"day-19/repository"
	"day-19/service"
	"encoding/json"
	"fmt"
	"io"
	"os"
)

func CreateAdmin(db *sql.DB) {
	// Input dari file JSON
	var admin model.Admin
	file, err := os.Open("body.json")
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&admin)
	if err != nil && err != io.EOF {
		fmt.Println("Error decoding JSON: ", err)
		return
	}

	// Cek apakah username kosong
	if admin.Username == "" {
		fmt.Println("Error: username tidak boleh kosong")
		return
	}

	// Proses
	repo := repository.NewAdminRepository(db)
	adminService := service.NewAdminService(repo)

	err = adminService.CreateDataAdmin(admin.AdminName, admin.Username, admin.Password)
	if err != nil {
		fmt.Println("Error while creating admin: ", err)
		return
	}

	// Output yang diinginkan
	response := model.Response{
		StatusCode: 200,
		Message:    "create success",
		Data:       admin,
	}

	jsonData, err := json.MarshalIndent(response, "", "  ")
	if err != nil {
		fmt.Println("Error marshaling response: ", err)
		return
	}

	fmt.Println(string(jsonData))
}
