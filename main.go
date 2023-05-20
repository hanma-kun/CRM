package main

import (
	"CRM/database"
	"CRM/lead"
	"fmt"

	"github.com/gofiber/fiber"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func setupRoutes(app *fiber.App) {
	app.Get("/api/v1/lead", lead.GetLead)
	app.Get("/api/v1/lead/:id", lead.GetLeads)
	app.Post("/api/v1/lead/:id", lead.NewLead)
	app.Delete("/api/v1/lead/:id", lead.DeleteLead)
}

func initDatabase() {
	var err error
	dsn := "hanma:mysqlpassword@tcp(127.0.0.1:3306)/crm?charset=utf8mb4&parseTime=True&loc=Local"
	database.DBConn, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect Database")
	}
	fmt.Println("Connection successful")
	database.DBConn.AutoMigrate(&lead.Lead{})
	fmt.Println("Database Migrated")
}

func main() {
	app := fiber.New()
	initDatabase()
	setupRoutes(app)
	app.Listen(3000)
}
