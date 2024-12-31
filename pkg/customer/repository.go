package customer

import (
	"clean-connect/pkg/entities"
	"fmt"

	"gorm.io/gorm"
)

type CustomerRepository interface {
	CreateUser(customer entities.Customer) error
	ValidateUser(email, password string) (entities.Customer, error)
	GetUser(id uint) (entities.Customer, error)
	GetUsers() ([]entities.Customer, error)
	GetUserByToken(token string) (entities.Customer, error)
	UpdateUser(customer entities.Customer) error
	UpdateCustomerByID(id uint, customer entities.Customer) error
	DeleteUser(id uint) error
	GetUserBill(userID uint, billID uint) (entities.Bill, error)
	UpdateBill(bill *entities.Bill) error
	GetUserBills(userID uint) ([]entities.Bill, error)
	PayBill(payment entities.Payment) error
	GetAllUserPaymentBills(userID uint) ([]entities.Payment, error)
}

type customerRepository struct {
	db *gorm.DB
}

func NewCustomerRepository(db *gorm.DB) *customerRepository {
	return &customerRepository{db}
}

func (r *customerRepository) CreateUser(customer entities.Customer) error {
	if r.db == nil {
		return fmt.Errorf("database connection is not initialized")
	}
	if err := r.db.Create(&customer).Error; err != nil {
		return err
	}
	return nil
}

func (r *customerRepository) ValidateUser(email, password string) (entities.Customer, error) {
	var customer entities.Customer
	if err := r.db.Where("email = ? AND password = ?", email, password).First(&customer).Error; err != nil {
		return customer, err
	}
	return customer, nil
}

func (r *customerRepository) GetUser(id uint) (entities.Customer, error) {
	var customer entities.Customer
	if err := r.db.First(&customer, id).Error; err != nil {
		return customer, err
	}
	return customer, nil
}

func (r *customerRepository) GetUsers() ([]entities.Customer, error) {
	var customers []entities.Customer
	if err := r.db.Find(&customers).Error; err != nil {
		return customers, err
	}
	return customers, nil
}

func (r *customerRepository) GetUserByToken(token string) (entities.Customer, error) {
	var customer entities.Customer
	if err := r.db.Where("verification_token = ?", token).First(&customer).Error; err != nil {
		return customer, err
	}
	return customer, nil
}

func (r *customerRepository) UpdateUser(customer entities.Customer) error {
	if err := r.db.Save(&customer).Error; err != nil {
		return err
	}
	return nil
}

func (r *customerRepository) UpdateCustomerByID(id uint, customer entities.Customer) error {
	if err := r.db.Model(&entities.Customer{}).Where("id = ?", id).Updates(&customer).Error; err != nil {
		return err
	}
	return nil
}

func (r *customerRepository) DeleteUser(id uint) error {
	if err := r.db.Delete(&entities.Customer{}, id).Error; err != nil {
		return err
	}
	return nil
}

func (r *customerRepository) GetUserBill(userID uint, billID uint) (entities.Bill, error) {
	var bill entities.Bill
	if err := r.db.Where("customer_id = ? AND id = ?", userID, billID).Preload("Customer").Preload("AdditionalBill").First(&bill).Error; err != nil {
		return bill, err
	}
	return bill, nil
}

func (r *customerRepository) UpdateBill(bill *entities.Bill) error {
	if err := r.db.Save(bill).Error; err != nil {
		return err
	}
	return nil
}

func (r *customerRepository) GetUserBills(userID uint) ([]entities.Bill, error) {
	var bills []entities.Bill
	if err := r.db.Where("customer_id = ?", userID).Order("created_at DESC").Find(&bills).Error; err != nil {
		return bills, err
	}
	return bills, nil
}

func (r *customerRepository) PayBill(payment entities.Payment) error {
	if err := r.db.Create(&payment).Error; err != nil {
		return err
	}
	return nil
}

func (r *customerRepository) GetAllUserPaymentBills(userID uint) ([]entities.Payment, error) {
	var payments []entities.Payment
	if err := r.db.
		Joins("JOIN bills ON bills.id = payments.bill_id").
		Where("bills.customer_id = ?", userID).
		Preload("Bill").
		Preload("Bill.Customer").
		Find(&payments).Error; err != nil {
		return payments, err
	}
	return payments, nil
}
