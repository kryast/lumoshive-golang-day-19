package main

import (
	"day-19/database"
	"day-19/repository"
	"day-19/service"
	"log"

	_ "github.com/lib/pq"
)

func main() {

	db, err := database.ConnectDB()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	repo := repository.NewAdminRepository(db) // Pastikan ini sesuai dengan implementasi yang ada

	// Membuat instance AdminService
	adminService := service.NewAdminService(repo)

	adminService.CreateDataAdmin("Ahmad 2")

}
