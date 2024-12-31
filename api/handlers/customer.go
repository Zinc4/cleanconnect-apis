package handlers

import (
	"clean-connect/api/presenters"
	"clean-connect/config"
	"clean-connect/pkg/customer"
	"clean-connect/pkg/entities"
	"errors"

	"github.com/gofiber/fiber/v2"
)

type CustomerHandler struct {
	customerService customer.CustomerService
}

func NewCustomerHandler(service customer.CustomerService) *CustomerHandler {
	return &CustomerHandler{
		customerService: service,
	}
}

func (h *CustomerHandler) RegisterCustomer(c *fiber.Ctx) error {
	var customer entities.Customer
	customer.Role = "customer"
	if err := c.BodyParser(&customer); err != nil {
		return c.Status(500).JSON(presenters.CustomerErrorResponse(err))
	}
	if err := h.customerService.RegisterUser(customer); err != nil {
		return c.Status(500).JSON(presenters.CustomerErrorResponse(err))
	}
	return c.Status(200).JSON(presenters.SuccessResponse("user created successfully"))
}

func (h *CustomerHandler) VerifyCustomer(c *fiber.Ctx) error {
	token := c.Params("token")
	if err := h.customerService.VerifyUser(token); err != nil {
		return c.Status(500).JSON(presenters.CustomerErrorResponse(err))
	}
	return c.Status(200).JSON(presenters.SuccessResponse("user verified successfully"))
}

func (h *CustomerHandler) LoginCustomer(c *fiber.Ctx) error {
	var customer entities.Customer
	if err := c.BodyParser(&customer); err != nil {
		return c.Status(500).JSON(presenters.CustomerErrorResponse(err))
	}
	token, err := h.customerService.LoginUser(customer.Email, customer.Password)
	if err != nil {
		return c.Status(500).JSON(presenters.CustomerErrorResponse(err))
	}
	return c.Status(200).JSON(&fiber.Map{
		"status":  "success",
		"message": "user found successfully",
		"data":    token,
	})
}

func (h *CustomerHandler) GetCustomer(c *fiber.Ctx) error {
	userID := c.Locals("userID").(uint)
	customer, err := h.customerService.GetUser(userID)
	if err != nil {
		return c.Status(500).JSON(presenters.CustomerErrorResponse(err))
	}
	return c.Status(200).JSON(presenters.GetCustomerSuccessResponse(customer))
}

func (h *CustomerHandler) GetCustomers(c *fiber.Ctx) error {
	customers, err := h.customerService.GetUsers()
	if err != nil {
		return c.Status(500).JSON(presenters.CustomerErrorResponse(err))
	}
	return c.Status(200).JSON(presenters.GetUsersSuccessResponse(customers))
}

func (h *CustomerHandler) UpdateCustomer(c *fiber.Ctx) error {
	userID := c.Locals("userID").(uint)
	type ReqUpdate struct {
		FirstName string `form:"first_name"`
		LastName  string `form:"last_name"`
		Email     string `form:"email"`
		NIK       string `form:"nik"`
		Address   string `form:"address"`
		NoHP      string `form:"no_hp"`
		Kategori  string `form:"kategori"`
		Image     string `form:"image"`
	}
	var request ReqUpdate
	if err := c.BodyParser(&request); err != nil {
		return c.Status(500).JSON(presenters.CustomerErrorResponse(err))
	}

	imgFile, err := c.FormFile("image")
	if err == nil {
		imgUrl, err := config.UploadToCloudinary(imgFile)
		if err != nil {
			return c.Status(500).JSON(presenters.CustomerErrorResponse(err))
		}
		request.Image = imgUrl
	}

	customer := entities.Customer{
		FirstName: request.FirstName,
		LastName:  request.LastName,
		Email:     request.Email,
		NIK:       request.NIK,
		Address:   request.Address,
		NoHP:      request.NoHP,
		Kategori:  request.Kategori,
		Image:     request.Image,
	}

	if err := h.customerService.UpdateCustomer(userID, customer); err != nil {
		return c.Status(500).JSON(presenters.CustomerErrorResponse(err))
	}
	return c.Status(200).JSON(presenters.SuccessResponse("user updated successfully"))
}

func (h *CustomerHandler) DeleteCustomer(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(500).JSON(presenters.CustomerErrorResponse(err))
	}
	if err := h.customerService.DeleteUser(uint(id)); err != nil {
		return c.Status(500).JSON(presenters.CustomerErrorResponse(err))
	}
	return c.Status(200).JSON(presenters.SuccessResponse("user deleted successfully"))
}

func (h *CustomerHandler) GetUserBill(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	userID := c.Locals("userID").(uint)
	if err != nil {
		return c.Status(500).JSON(presenters.CustomerErrorResponse(err))
	}
	bill, err := h.customerService.GetUserBill(userID, uint(id))
	if err != nil {
		return c.Status(500).JSON(presenters.CustomerErrorResponse(err))
	}
	return c.Status(200).JSON(presenters.GetBillCustomerSuccessResponse(bill))
}

func (h *CustomerHandler) GetBills(c *fiber.Ctx) error {
	userID := c.Locals("userID").(uint)
	bills, err := h.customerService.GetUserBills(userID)
	if err != nil {
		return c.Status(500).JSON(presenters.CustomerErrorResponse(err))
	}
	return c.Status(200).JSON(presenters.GetUserBillsSuccessResponse(bills))
}

func (h *CustomerHandler) PayBill(c *fiber.Ctx) error {
	var req presenters.PaymentRequest
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(500).JSON(presenters.CustomerErrorResponse(err))
	}

	userID := c.Locals("userID").(uint)

	if err := c.BodyParser(&req); err != nil {
		return c.Status(500).JSON(presenters.CustomerErrorResponse(err))
	}

	bill, err := h.customerService.GetUserBill(userID, uint(id))
	if err != nil {
		return c.Status(500).JSON(presenters.CustomerErrorResponse(err))
	}
	if bill.Status == "pending" {
		return c.Status(500).JSON(presenters.CustomerErrorResponse(errors.New("invalid or already paid bill")))
	}

	imgFile, err := c.FormFile("image")
	if err != nil {
		return c.Status(500).JSON(presenters.CustomerErrorResponse(err))
	}

	imgUrl, err := config.UploadToCloudinary(imgFile)
	if err != nil {
		return c.Status(500).JSON(presenters.CustomerErrorResponse(err))
	}

	payment := entities.Payment{
		Status:     "pending",
		Image:      imgUrl,
		BillID:     bill.ID,
		CustomerID: userID,
	}

	if err := h.customerService.PayBill(payment); err != nil {
		return c.Status(500).JSON(presenters.CustomerErrorResponse(err))
	}

	bill.Status = "pending"
	if err := h.customerService.UpdateBill(&bill); err != nil {
		return c.Status(500).JSON(presenters.CustomerErrorResponse(err))
	}

	return c.Status(200).JSON(presenters.SuccessResponse("Payment submitted successfully, pending verification"))
}

func (h *CustomerHandler) GetCustomerPaymentBills(c *fiber.Ctx) error {
	userID := c.Locals("userID").(uint)
	bills, err := h.customerService.GetAllUserPaymentBills(userID)
	if err != nil {
		return c.Status(500).JSON(presenters.CustomerErrorResponse(err))
	}
	return c.Status(200).JSON(presenters.GetUserPaymentBillsSuccessResponse(bills))

}
