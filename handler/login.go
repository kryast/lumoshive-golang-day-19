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

func Login(db *sql.DB) {

	//input
	user := model.Admin{}
	file, err := os.Open("body.json")

	if err != nil {
		fmt.Println("Error : ", err)
	}

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&user)
	if err != nil && err != io.EOF {
		fmt.Println("error decoding JSON: ", err)
	}

	// proses
	repo := repository.NewAdminRepository(db)
	customerService := service.NewAdminService(repo)

	result, err := customerService.LoginService(user)

	if err != nil {
		fmt.Println("Error : ", err)
	}

	// output
	jsonData, err := json.MarshalIndent(result, " ", "")

	if err != nil {
		fmt.Println("err :", err)
	}

	fmt.Println(string(jsonData))
}
