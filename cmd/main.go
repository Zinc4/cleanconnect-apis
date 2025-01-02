package main

import (
	"clean-connect/api/handlers"
	"clean-connect/api/routes"
	"clean-connect/database"
	"clean-connect/pkg/admin"
	"clean-connect/pkg/customer"

	"github.com/gofiber/fiber/v2"
)

func main() {

	database.Connect()

	app := fiber.New()

	customerRepo := customer.NewCustomerRepository(database.DB)
	adminRepo := admin.NewAdminRepository(database.DB)
	customerService := customer.NewCustomerService(customerRepo)
	adminService := admin.NewAdminService(adminRepo)
	customerHandler := handlers.NewCustomerHandler(customerService, adminService)
	adminHandler := handlers.NewAdminHandler(adminService, customerService)

	scheleduler := admin.NewScheduler(adminService)
	scheleduler.Start()

	routes.SetupRoutes(app, customerHandler, adminHandler)

	app.Listen(":8080")
}
