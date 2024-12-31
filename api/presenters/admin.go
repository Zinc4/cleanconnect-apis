package presenters

import (
	"clean-connect/pkg/entities"

	"github.com/gofiber/fiber/v2"
)

func AdminErrorResponse(err error) fiber.Map {
	return fiber.Map{
		"status":  false,
		"message": err.Error(),
		"data":    "",
	}
}

type PendingPayment struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
	Image    string `json:"image"`
	Amount   int    `json:"amount"`
}

func GetPendingUsersPaymentsSuccessResponse(data []entities.Payment) fiber.Map {
	var pendingPayments []PendingPayment

	for _, payment := range data {
		pendingPayment := PendingPayment{
			ID:       payment.Bill.CustomerID,
			Username: payment.Bill.Customer.FirstName + " " + payment.Bill.Customer.LastName,
			Image:    payment.Bill.Customer.Image,
			Amount:   payment.Bill.Amount + payment.Bill.AdditionalBill.Price,
		}

		pendingPayments = append(pendingPayments, pendingPayment)
	}

	return fiber.Map{
		"status":  true,
		"message": "payments found successfully",
		"data":    pendingPayments,
	}
}

type SuccessPayment struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
	Address  string `json:"address"`
	Date     string `json:"date"`
	Amount   int    `json:"amount"`
}

func GetSuccessUsersPaymentsSuccessResponse(data []entities.Payment) fiber.Map {
	var successPayments []SuccessPayment

	for _, payment := range data {
		successPayment := SuccessPayment{
			ID:       payment.Bill.CustomerID,
			Username: payment.Bill.Customer.FirstName + " " + payment.Bill.Customer.LastName,
			Address:  payment.Bill.Customer.Address,
			Date:     payment.CreatedAt.Format("2006-01-02"),
			Amount:   payment.Bill.Amount + payment.Bill.AdditionalBill.Price,
		}

		successPayments = append(successPayments, successPayment)
	}

	return fiber.Map{
		"status":  true,
		"message": "payments found successfully",
		"data":    successPayments,
	}
}

type UserBills struct {
	BillNo   uint   `json:"bill_no"`
	Username string `json:"username"`
	Date     string `json:"date"`
	DueDate  string `json:"due_date"`
	Amount   int    `json:"amount"`
	Status   string `json:"status"`
}

func GetBillsSuccessResponse(data []entities.Payment) fiber.Map {
	var bills []UserBills

	for _, payment := range data {
		bill := UserBills{
			BillNo:   payment.BillID,
			Username: payment.Bill.Customer.FirstName + " " + payment.Bill.Customer.LastName,
			Date:     payment.Bill.BillDate.Format("2006-01-02"),
			DueDate:  payment.Bill.BillDue.Format("2006-01-02"),
			Amount:   payment.Bill.Amount + payment.Bill.AdditionalBill.Price,
			Status:   payment.Status,
		}

		bills = append(bills, bill)
	}

	return fiber.Map{
		"status":  true,
		"message": "bills found successfully",
		"data":    bills,
	}
}

func GetBillsUserSuccessResponse(data []entities.Bill) fiber.Map {
	var bills []UserBills

	for _, payment := range data {
		bill := UserBills{
			BillNo:   payment.ID,
			Username: payment.Customer.FirstName + " " + payment.Customer.LastName,
			Date:     payment.BillDate.Format("2006-01-02"),
			DueDate:  payment.BillDue.Format("2006-01-02"),
			Amount:   payment.Amount + payment.AdditionalBill.Price,
			Status:   payment.Status,
		}

		bills = append(bills, bill)
	}

	return fiber.Map{
		"status":  true,
		"message": "bills found successfully",
		"data":    bills,
	}
}

type AdditionalBill struct {
	ID    uint   `json:"id"`
	Name  string `json:"name"`
	Price int    `json:"price"`
}

func GetAdditionalBillsSuccessResponse(data []entities.AdditionalBill) fiber.Map {
	var additionalBills []AdditionalBill

	for _, additionalBill := range data {
		additionalBills = append(additionalBills, AdditionalBill{
			ID:    additionalBill.ID,
			Name:  additionalBill.Name,
			Price: additionalBill.Price,
		})
	}

	return fiber.Map{
		"status":  true,
		"message": "additional bills found successfully",
		"data":    additionalBills,
	}
}

type Users struct {
	ID       uint   `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	NIK      string `json:"nik"`
	Address  string `json:"address"`
	Role     string `json:"role"`
	Kategori string `json:"kategori"`
}

func GetUsersSuccessResponse(data []entities.Customer) fiber.Map {
	var users []Users

	for _, user := range data {
		users = append(users, Users{
			ID:       user.ID,
			Name:     user.FirstName + " " + user.LastName,
			Email:    user.Email,
			NIK:      user.NIK,
			Address:  user.Address,
			Role:     user.Role,
			Kategori: user.Kategori,
		})
	}

	return fiber.Map{
		"status":  true,
		"message": "users found successfully",
		"data":    users,
	}
}

type Payment struct {
	ID         uint   `json:"id"`
	BillID     uint   `json:"bill_id"`
	CustomerID uint   `json:"customer_id"`
	Status     string `json:"status"`
	Image      string `json:"image"`
	Name       string `json:"name"`
}

func GetPaymentSuccessResponse(data entities.Payment) fiber.Map {
	return fiber.Map{
		"status":  true,
		"message": "payment found successfully",
		"data":    Payment{ID: data.ID, BillID: data.BillID, CustomerID: data.Bill.CustomerID, Status: data.Status, Image: data.Image, Name: data.Bill.Customer.FirstName + " " + data.Bill.Customer.LastName},
	}
}
