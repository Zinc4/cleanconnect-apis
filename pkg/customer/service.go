package customer

import (
	"clean-connect/config"
	"clean-connect/pkg/entities"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type CustomerService interface {
	RegisterUser(customer entities.Customer) error
	LoginUser(email, password string) (string, error)
	GetUser(id uint) (entities.Customer, error)
	GetUsers() ([]entities.Customer, error)
	UpdateCustomer(id uint, customer entities.Customer) error
	VerifyUser(token string) error
	DeleteUser(id uint) error
	GetUserBill(userID uint, billID uint) (entities.Bill, error)
	UpdateBill(bill *entities.Bill) error
	GetUserBills(userID uint) ([]entities.Bill, error)
	PayBill(payment entities.Payment) error
	GetAllUserPaymentBills(userID uint) ([]entities.Payment, error)
	CreateLog(log entities.Log) error
	GetLogsByUser(userID uint) ([]entities.Log, error)
	GetTotalUserBills(userID uint) (int64, error)
	GetAmountSuccessfulBills(userID uint) (int64, error)
	GetActiveBills(userID uint) (int64, error)
	GetAmountPendingPaymentBills(userID uint) (int64, error)
}

type customerService struct {
	repository CustomerRepository
}

func NewCustomerService(repository CustomerRepository) *customerService {
	return &customerService{repository}
}

func (s *customerService) RegisterUser(customer entities.Customer) error {
	if customer.NIK != "" && len(customer.NIK) != 16 {
		return fmt.Errorf("NIK must be 16 characters")
	}

	token := config.GenerateVerificationToken()
	customer.VerificationToken = token

	if err := s.repository.CreateUser(customer); err != nil {
		return err
	}
	return config.SendVerificationEmail(customer.Email, token)

}

func (s *customerService) LoginUser(email, password string) (string, error) {
	user, err := s.repository.ValidateUser(email, password)
	if err != nil {
		return "", err
	}

	if !user.Verified {
		return "", fmt.Errorf("user not verified")
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email":  email,
		"userID": user.ID,
		"role":   user.Role,
		"exp":    time.Now().Add(time.Hour * 24).Unix(),
	})

	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		return "", err
	}

	return t, err
}

func (s *customerService) GetUser(id uint) (entities.Customer, error) {
	return s.repository.GetUser(id)
}

func (s *customerService) GetUsers() ([]entities.Customer, error) {
	return s.repository.GetUsers()
}

func (s *customerService) UpdateCustomer(id uint, customer entities.Customer) error {
	return s.repository.UpdateCustomerByID(id, customer)
}

func (s *customerService) VerifyUser(token string) error {
	user, err := s.repository.GetUserByToken(token)
	if err != nil {
		return fmt.Errorf("invalid or expired token")
	}

	user.Verified = true
	user.VerificationToken = ""
	return s.repository.UpdateUser(user)
}

func (s *customerService) DeleteUser(id uint) error {
	return s.repository.DeleteUser(id)
}

func (s *customerService) GetUserBill(userID uint, billID uint) (entities.Bill, error) {
	return s.repository.GetUserBill(userID, billID)
}

func (s *customerService) UpdateBill(bill *entities.Bill) error {
	return s.repository.UpdateBill(bill)
}

func (s *customerService) GetUserBills(userID uint) ([]entities.Bill, error) {
	return s.repository.GetUserBills(userID)
}

func (s *customerService) PayBill(payment entities.Payment) error {
	return s.repository.PayBill(payment)
}

func (s *customerService) GetAllUserPaymentBills(userID uint) ([]entities.Payment, error) {
	return s.repository.GetAllUserPaymentBills(userID)
}

func (s *customerService) CreateLog(log entities.Log) error {
	return s.repository.CreateLog(log)
}

func (s *customerService) GetLogsByUser(userID uint) ([]entities.Log, error) {
	return s.repository.GetLogsByUser(userID)
}

func (s *customerService) GetTotalUserBills(userID uint) (int64, error) {
	return s.repository.GetTotalUserBills(userID)
}

func (s *customerService) GetAmountSuccessfulBills(userID uint) (int64, error) {
	return s.repository.GetAmountSuccessfulBills(userID)
}

func (s *customerService) GetActiveBills(userID uint) (int64, error) {
	return s.repository.GetActiveBills(userID)
}

func (s *customerService) GetAmountPendingPaymentBills(userID uint) (int64, error) {
	return s.repository.GetAmountPendingPaymentBills(userID)
}
