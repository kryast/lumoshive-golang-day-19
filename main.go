package main

import (
	"day-19/database"
	"day-19/handler"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func main() {

	db, err := database.ConnectDB()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	var endpoint string
	fmt.Print("masukkan enpoint : ")
	fmt.Scan(&endpoint)

	switch endpoint {
	case "login":
		handler.Login(db)
	case "create":
		handler.CreateAdmin(db)
	}
	// repo := repository.NewAdminRepository(db)
	// adminService := service.NewAdminService(repo)

	// adminService.CreateDataAdmin("Ahmad 1", "user", "pass")

}
