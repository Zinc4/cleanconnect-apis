package presenters

import (
	"clean-connect/pkg/entities"
	"time"

	"github.com/gofiber/fiber/v2"
)

type Customer struct {
	ID       uint   `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Address  string `json:"address"`
	NoHP     string `json:"no_hp"`
	Kategori string `json:"kategori"`
	Image    string `json:"image"`
	NIK      string `json:"nik"`
}

func SuccessResponse(message string) fiber.Map {
	return fiber.Map{
		"status":  true,
		"message": message,
	}
}

func GetCustomerSuccessResponse(data entities.Customer) fiber.Map {

	customer := Customer{
		ID:       data.ID,
		Name:     data.FirstName + " " + data.LastName,
		NoHP:     data.NoHP,
		Address:  data.Address,
		Email:    data.Email,
		Kategori: data.Kategori,
		Image:    data.Image,
		NIK:      data.NIK,
	}

	return fiber.Map{
		"status":  true,
		"message": "user found successfully",
		"data":    customer,
	}
}

func GetCustomersSuccessResponse(data []entities.Customer) fiber.Map {
	return fiber.Map{
		"status":  true,
		"message": "users found successfully",
		"data":    data,
	}
}

func CustomerErrorResponse(err error) fiber.Map {
	return fiber.Map{
		"status":  false,
		"message": err.Error(),
		"data":    "",
	}
}

type Bill struct {
	ID                 uint      `json:"id"`
	CustomerName       string    `json:"customer_name"`
	Description        string    `json:"description"`
	Amount             int       `json:"amount"`
	BillDate           time.Time `json:"bill_date"`
	BillDue            time.Time `json:"bill_due"`
	Status             string    `json:"status"`
	Category           string    `json:"category"`
	AdditionalBillName string    `json:"additional_bill_name"`
	AdditionalAmount   int       `json:"additional_amount"`
	QrUrl              string    `json:"qr_url"`
}

func GetBillCustomerSuccessResponse(data entities.Bill) fiber.Map {

	bill := Bill{
		ID:                 data.ID,
		CustomerName:       data.Customer.FirstName,
		Description:        data.Description,
		Amount:             data.Amount,
		BillDate:           data.BillDate,
		BillDue:            data.BillDue,
		Status:             data.Status,
		Category:           data.Category,
		AdditionalBillName: data.AdditionalBill.Name,
		AdditionalAmount:   data.AdditionalBill.Price,
		QrUrl:              data.QrUrl,
	}

	return fiber.Map{
		"status":  true,
		"message": "bill found successfully",
		"data":    bill,
	}
}

type PaymentRequest struct {
	Image string `json:"image" form:"image"`
}

type Bills struct {
	BillNo  uint   `json:"bill_no"`
	Name    string `json:"name"`
	Amount  int    `json:"amount"`
	DueDate string `json:"due_date"`
	Status  string `json:"status"`
}

func GetUserBillsSuccessResponse(data []entities.Bill) fiber.Map {
	var bills []Bills

	for _, bill := range data {
		bills = append(bills, Bills{
			BillNo:  bill.ID,
			Name:    bill.Description,
			Amount:  bill.Amount + bill.AdditionalBill.Price,
			DueDate: bill.BillDue.Format("2006-01-02"),
			Status:  bill.Status,
		})
	}

	return fiber.Map{
		"status":  true,
		"message": "bills found successfully",
		"data":    bills,
	}
}

type PaymentBills struct {
	BillNo  uint   `json:"bill_no"`
	Amount  int    `json:"amount"`
	Date    string `json:"date"`
	DueDate string `json:"due_date"`
	Status  string `json:"status"`
}

func GetUserPaymentBillsSuccessResponse(data []entities.Payment) fiber.Map {
	var bills []PaymentBills

	for _, payment := range data {
		bill := PaymentBills{
			BillNo:  payment.BillID,
			Amount:  payment.Bill.Amount + payment.Bill.AdditionalBill.Price,
			Date:    payment.CreatedAt.Format("2006-01-02"),
			DueDate: payment.Bill.BillDue.Format("2006-01-02"),
			Status:  payment.Status,
		}
		bills = append(bills, bill)
	}

	return fiber.Map{
		"status":  true,
		"message": "bills found successfully",
		"data":    bills,
	}
}

type UserLogs struct {
	ID         uint      `json:"id"`
	CreatedAt  time.Time `json:"created_at"`
	ChangeType string    `json:"change_type"`
	OldValue   string    `json:"old_value"`
	NewValue   string    `json:"new_value"`
	Status     string    `json:"status"`
}

func GetUserLogsSuccessResponse(data []entities.Log) fiber.Map {
	var logs []UserLogs

	for _, log := range data {
		logs = append(logs, UserLogs{
			ID:         log.ID,
			CreatedAt:  log.CreatedAt,
			ChangeType: log.ChangeType,
			OldValue:   log.OldValue,
			NewValue:   log.NewValue,
			Status:     log.Status,
		})
	}

	return fiber.Map{
		"status":  true,
		"message": "logs found successfully",
		"data":    logs,
	}
}
