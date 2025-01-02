package routes

import (
	"clean-connect/api/handlers"
	"clean-connect/api/middlewares"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func SetupRoutes(app *fiber.App, hu *handlers.CustomerHandler, ha *handlers.AdminHandler) {

	app.Use(cors.New(cors.Config{AllowOrigins: "*"}))

	api := app.Group("/api")

	api.Post("/webhook", ha.Webhook)

	customer := api.Group("/user")

	customer.Post("/register", hu.RegisterCustomer)
	customer.Get("/verify/:token", hu.VerifyCustomer)
	customer.Post("/login", hu.LoginCustomer)
	customer.Get("/profile", middlewares.Protected(), hu.GetCustomer)
	customer.Put("/profile", middlewares.Protected(), hu.UpdateCustomer)
	customer.Get("/bill/:id", middlewares.Protected(), hu.GetUserBill)
	customer.Post("/bill/:id", middlewares.Protected(), hu.PayBill)
	customer.Get("/bills", middlewares.Protected(), hu.GetBills)
	customer.Get("/payments", middlewares.Protected(), hu.GetCustomerPaymentBills)
	customer.Get("/logs", middlewares.Protected(), hu.GetLogsByUser)
	customer.Get("/total-dashboard", middlewares.Protected(), hu.GetTotalUserDashboard)
	customer.Get("/total-payments-history", middlewares.Protected(), hu.GetTotalHistory)

	admin := api.Group("/admin")
	admin.Get("/users", middlewares.AdminOnly(), hu.GetCustomers)
	admin.Put("/users/:id", middlewares.AdminOnly(), hu.UpdateCustomer)
	admin.Delete("/users/:id", middlewares.AdminOnly(), hu.DeleteCustomer)
	admin.Get("/profile", middlewares.AdminOnly(), hu.GetCustomer)

	admin.Get("/bills", middlewares.SuperAdminOnly(), ha.GetBills)
	admin.Get("/bills/:id", middlewares.AdminOnly(), ha.GetBill)
	admin.Delete("/bills/:id", middlewares.AdminOnly(), ha.DeleteBill)
	admin.Post("/bills", middlewares.AdminOnly(), ha.CreateBill)
	admin.Post("/bills-mass", middlewares.AdminOnly(), ha.CreateMassBill)
	admin.Post("/bills/additional", middlewares.AdminOnly(), ha.CreateAdditionalBill)
	admin.Get("/additionalbill", middlewares.AdminOnly(), ha.GetAdditionalBills)
	admin.Get("/payments/bill/:bill_id", middlewares.AdminOnly(), ha.GetPaymentByBillID)
	admin.Get("/payments/pending", middlewares.AdminOnly(), ha.GetPendingUsersPayments)
	admin.Get("/payments/success", middlewares.AdminOnly(), ha.GetSuccessUsersPayments)
	admin.Get("/payment/verify/:bill_id", middlewares.AdminOnly(), ha.VerifyPayment)
	admin.Get("/total-dashboard", middlewares.AdminOnly(), ha.GetTotalDashboard)
	admin.Get("/total-bills", middlewares.AdminOnly(), ha.GetTotalBillsManagement)
	admin.Get("/notifications", middlewares.AdminOnly(), ha.GetNotifications)
	// admin.Get("/bills", middlewares.AdminOnly(), ha.GetUsersBills)
}
