package handlers

import (
	"clean-connect/api/presenters"
	"clean-connect/config"
	"clean-connect/pkg/admin"
	"clean-connect/pkg/customer"
	"clean-connect/pkg/entities"
	"database/sql"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
)

type AdminHandler struct {
	adminService    admin.AdminService
	customerService customer.CustomerService
}

func NewAdminHandler(adminService admin.AdminService, customerService customer.CustomerService) *AdminHandler {
	return &AdminHandler{adminService, customerService}
}

func (h *AdminHandler) GetBills(c *fiber.Ctx) error {
	bills, err := h.adminService.GetBills()
	if err != nil {
		return err
	}
	return c.Status(200).JSON(presenters.GetBillsUserSuccessResponse(bills))
}

func (h *AdminHandler) GetBill(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return err
	}
	bill, err := h.adminService.GetBill(uint(id))
	if err != nil {
		return err
	}
	return c.JSON(bill)
}

func (h *AdminHandler) DeleteBill(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return err
	}
	err = h.adminService.DeleteBill(uint(id))
	if err != nil {
		return err
	}
	return c.JSON(fiber.Map{"message": "Bill deleted successfully"})
}

func (h *AdminHandler) CreateBill(c *fiber.Ctx) error {
	type RequsetBody struct {
		CustomerID       uint      `json:"customer_id"`
		AdditionalBillID uint      `json:"additional_bill_id"`
		Description      string    `json:"description"`
		Amount           int       `json:"amount"`
		BillDate         time.Time `json:"bill_date"`
		BillDue          time.Time `json:"bill_due"`
		Status           string    `json:"status"`
	}

	var body RequsetBody
	if err := c.BodyParser(&body); err != nil {
		return err
	}

	bill := entities.Bill{
		AdditionalBillID: body.AdditionalBillID,
		CustomerID:       body.CustomerID,
		Description:      body.Description,
		Amount:           body.Amount,
		BillDate:         body.BillDate,
		BillDue:          body.BillDue,
		Status:           "Belum Dibayar",
	}

	if bill.AdditionalBillID == 0 {
		bill.AdditionalBillID = 2
	}

	additionalBill, err := h.adminService.GetAdditionalBill(bill.AdditionalBillID)
	if err != nil && err != sql.ErrNoRows {
		return err
	}

	totalAmount := additionalBill.Price + bill.Amount

	qrUrl, err := config.GenerateMayarQRCode(totalAmount)
	if err != nil {
		return err
	}
	bill.QrUrl = qrUrl

	bill.Amount = totalAmount

	err = h.adminService.CreateBill(bill)
	if err != nil {
		return err
	}
	return c.JSON(fiber.Map{"message": "Bill created successfully"})
}

func (h *AdminHandler) CreateAdditionalBill(c *fiber.Ctx) error {

	var bill entities.AdditionalBill

	if err := c.BodyParser(&bill); err != nil {
		return err
	}
	err := h.adminService.CreateAdditionalBill(bill)
	if err != nil {
		return err
	}
	return c.JSON(fiber.Map{"message": "Bill created successfully"})
}

func (h *AdminHandler) CreateMassBill(c *fiber.Ctx) error {
	customers, err := h.customerService.GetUsers()
	if err != nil {
		return err
	}

	type RequsetBody struct {
		AdditionalBillID uint      `json:"additional_bill_id"`
		Description      string    `json:"description"`
		Amount           uint      `json:"amount"`
		BillDate         time.Time `json:"bill_date"`
		BillDue          time.Time `json:"bill_due"`
	}

	var body RequsetBody
	if err := c.BodyParser(&body); err != nil {
		return err
	}

	var bills []entities.Bill

	for _, customer := range customers {
		bill := entities.Bill{
			CustomerID:       customer.ID,
			AdditionalBillID: body.AdditionalBillID,
			Status:           "Belum Dibayar",
			Description:      body.Description,
			Amount:           int(body.Amount),
			BillDate:         body.BillDate,
			BillDue:          body.BillDue,
		}

		additionalBill, err := h.adminService.GetAdditionalBill(bill.AdditionalBillID)
		if err != nil {
			return err
		}
		bill.Amount += additionalBill.Price

		bills = append(bills, bill)
	}

	err = h.adminService.CreateMassBill(bills)
	if err != nil {
		return c.Status(500).JSON(err)
	}
	return c.JSON(fiber.Map{"message": "Bill created successfully"})
}

func (h *AdminHandler) GetAdditionalBills(c *fiber.Ctx) error {
	bills, err := h.adminService.GetAdditionalBills()
	if err != nil {
		return c.Status(500).JSON(err)
	}
	return c.Status(200).JSON(presenters.GetAdditionalBillsSuccessResponse(bills))
}

func (h *AdminHandler) GetPaymentByBillID(c *fiber.Ctx) error {
	billID, err := c.ParamsInt("bill_id")
	if err != nil {
		return c.Status(500).JSON(presenters.AdminErrorResponse(err))
	}

	bill, err := h.adminService.GetBill(uint(billID))
	if err != nil {
		return c.Status(500).JSON(presenters.AdminErrorResponse(err))
	}

	payments, err := h.adminService.GetPaymentByBillIDAndCustomerID(uint(billID), bill.CustomerID)
	if err != nil {
		return c.Status(500).JSON(presenters.AdminErrorResponse(err))
	}
	return c.Status(200).JSON(presenters.GetPaymentSuccessResponse(payments))
}

func (h *AdminHandler) GetPendingUsersPayments(c *fiber.Ctx) error {
	payments, err := h.adminService.GetPendingUsersPayments()
	if err != nil {
		return err
	}

	return c.Status(200).JSON(presenters.GetPendingUsersPaymentsSuccessResponse(payments))
}

func (h *AdminHandler) GetSuccessUsersPayments(c *fiber.Ctx) error {
	payments, err := h.adminService.GetSuccessPayments()
	if err != nil {
		return err
	}
	return c.Status(200).JSON(presenters.GetSuccessUsersPaymentsSuccessResponse(payments))
}

func (h *AdminHandler) GetUsersBills(c *fiber.Ctx) error {
	bills, err := h.adminService.GetAllPaymentBills()
	if err != nil {
		return err
	}
	return c.Status(200).JSON(presenters.GetBillsSuccessResponse(bills))
}

func (h *AdminHandler) VerifyPayment(c *fiber.Ctx) error {
	billID, err := c.ParamsInt("bill_id")
	if err != nil {
		return err
	}

	bill, err := h.adminService.GetBill(uint(billID))
	if err != nil {
		return err
	}
	bill.Status = "Dibayar"
	err = h.adminService.UpdateBill(&bill)
	if err != nil {
		return err
	}

	payment, err := h.adminService.GetPaymentByBillIDAndCustomerID(uint(billID), bill.CustomerID)
	if err != nil {
		return err
	}
	payment.Status = "paid"

	err = h.adminService.UpdatePayment(&payment)
	if err != nil {
		return err
	}

	return c.Status(200).JSON(presenters.SuccessResponse("Payment verified successfully"))
}

func (h *AdminHandler) Webhook(c *fiber.Ctx) error {
	type WebhookPayload struct {
		Event      string `json:"event"`       // Nama event (e.g., payment.success)
		MerchantID string `json:"merchant_id"` // ID Merchant
		Amount     int64  `json:"amount"`      // Jumlah transaksi
		Status     string `json:"status"`      // Status transaksi
	}

	var payload WebhookPayload

	// Parse JSON payload
	if err := c.BodyParser(&payload); err != nil {
		log.Printf("Error parsing webhook: %v", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid payload",
		})
	}

	// Log payload
	log.Printf("Webhook received: %+v", payload)

	// Handle specific event
	if payload.Event == "payment.success" {
		log.Printf("Payment success for  Amount: %d", payload.Amount)
	} else {
		log.Printf("Unhandled event: %s", payload.Event)
	}

	return c.SendStatus(fiber.StatusOK)

}
