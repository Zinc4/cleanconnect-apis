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
	CreateLog(log entities.Log) error
	GetLogsByUser(userID uint) ([]entities.Log, error)
	GetTotalUserBills(userID uint) (int64, error)
	GetAmountSuccessfulBills(userID uint) (int64, error)
	GetActiveBills(userID uint) (int64, error)
	GetAmountPendingPaymentBills(userID uint) (int64, error)
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

func (r *customerRepository) CreateLog(log entities.Log) error {
	if err := r.db.Create(&log).Error; err != nil {
		return err
	}
	return nil
}

func (r *customerRepository) GetLogsByUser(userID uint) ([]entities.Log, error) {
	var logs []entities.Log
	if err := r.db.Where("user_id = ?", userID).Order("created_at DESC").Find(&logs).Error; err != nil {
		return logs, err
	}
	return logs, nil
}

func (r *customerRepository) GetTotalUserBills(userID uint) (int64, error) {
	var count int64
	if err := r.db.Model(&entities.Bill{}).Where("customer_id = ?", userID).Count(&count).Error; err != nil {
		return count, err
	}
	return count, nil
}

func (r *customerRepository) GetAmountSuccessfulBills(userID uint) (int64, error) {
	var totalAmount int64
	if err := r.db.Model(&entities.Bill{}).
		Joins("JOIN payments ON payments.bill_id = bills.id").
		Where("bills.customer_id = ?", userID).
		Where("payments.status = ?", "paid").
		Select("COALESCE(SUM(bills.amount), 0)").
		Scan(&totalAmount).Error; err != nil {
		return totalAmount, err
	}
	return totalAmount, nil

}

func (r *customerRepository) GetActiveBills(userID uint) (int64, error) {
	var count int64
	if err := r.db.Model(&entities.Bill{}).Where("customer_id = ?", userID).Where("status = ?", "Belum Dibayar").Count(&count).Error; err != nil {
		return count, err
	}
	return count, nil
}

func (r *customerRepository) GetAmountPendingPaymentBills(userID uint) (int64, error) {
	var totalAmount int64
	if err := r.db.Model(&entities.Payment{}).
		Joins("JOIN bills ON bills.id = payments.bill_id").
		Where("bills.customer_id = ?", userID).
		Where("payments.status = ?", "pending").
		Select("COALESCE(SUM(bills.amount), 0)").
		Scan(&totalAmount).Error; err != nil {
		return totalAmount, err
	}
	return totalAmount, nil
}
