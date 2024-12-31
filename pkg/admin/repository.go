package admin

import (
	"clean-connect/pkg/entities"

	"gorm.io/gorm"
)

type AdminRepository interface {
	CreateBill(bill entities.Bill) error
	CreateBills(bills []entities.Bill) error
	CreateAdditionalBill(additionalBill entities.AdditionalBill) error
	GetBills() ([]entities.Bill, error)
	GetBill(id uint) (entities.Bill, error)
	GetAdditionalBills() ([]entities.AdditionalBill, error)
	GetAdditionalBill(id uint) (entities.AdditionalBill, error)
	DeleteBill(id uint) error
	UpdateBill(bill *entities.Bill) error
	UpdatePayment(payment *entities.Payment) error
	GetPaymentByBillIDAndCustomerID(billID uint, customerID uint) (entities.Payment, error)
	GetPendingUsersPayments() ([]entities.Payment, error)
	GetSuccessPayments() ([]entities.Payment, error)
	GetAllPaymentBills() ([]entities.Payment, error)
}

type adminRepository struct {
	db *gorm.DB
}

func NewAdminRepository(db *gorm.DB) *adminRepository {
	return &adminRepository{db}
}

func (r *adminRepository) CreateBill(bill entities.Bill) error {
	if err := r.db.Create(&bill).Error; err != nil {
		return err
	}
	return nil
}

func (r *adminRepository) CreateBills(bills []entities.Bill) error {
	if err := r.db.Create(&bills).Error; err != nil {
		return err
	}
	return nil
}

func (r *adminRepository) CreateAdditionalBill(additionalBill entities.AdditionalBill) error {
	if err := r.db.Create(&additionalBill).Error; err != nil {
		return err
	}
	return nil
}

func (r *adminRepository) GetBills() ([]entities.Bill, error) {
	var bills []entities.Bill
	if err := r.db.Preload("Customer").Order("created_at DESC").Find(&bills).Error; err != nil {
		return nil, err
	}
	return bills, nil
}

func (r *adminRepository) GetBill(id uint) (entities.Bill, error) {
	var bill entities.Bill
	if err := r.db.First(&bill, id).Error; err != nil {
		return bill, err
	}
	return bill, nil
}

func (r *adminRepository) GetAdditionalBills() ([]entities.AdditionalBill, error) {
	var additionalBills []entities.AdditionalBill
	if err := r.db.Find(&additionalBills).Error; err != nil {
		return nil, err
	}
	return additionalBills, nil
}

func (r *adminRepository) GetAdditionalBill(id uint) (entities.AdditionalBill, error) {
	var additionalBill entities.AdditionalBill
	if err := r.db.Where("id = ? ", id).First(&additionalBill).Error; err != nil {
		return additionalBill, err
	}
	return additionalBill, nil
}

func (r *adminRepository) DeleteBill(id uint) error {
	if err := r.db.Delete(&entities.Bill{}, id).Error; err != nil {
		return err
	}
	return nil
}

func (r *adminRepository) UpdateBill(bill *entities.Bill) error {
	if err := r.db.Save(bill).Error; err != nil {
		return err
	}
	return nil
}

func (r *adminRepository) UpdatePayment(payment *entities.Payment) error {
	if err := r.db.Save(payment).Error; err != nil {
		return err
	}
	return nil
}

func (r *adminRepository) GetPaymentByBillIDAndCustomerID(billID uint, customerID uint) (entities.Payment, error) {
	var payments entities.Payment
	if err := r.db.Where("bill_id = ? AND customer_id = ?", billID, customerID).Preload("Bill").Preload("Bill.Customer").First(&payments).Error; err != nil {
		return payments, err
	}
	return payments, nil
}

func (r *adminRepository) GetPendingUsersPayments() ([]entities.Payment, error) {
	var payments []entities.Payment
	if err := r.db.Where("status = ?", "pending").Preload("Bill").Preload("Bill.Customer").Find(&payments).Error; err != nil {
		return nil, err
	}
	return payments, nil
}

func (r *adminRepository) GetSuccessPayments() ([]entities.Payment, error) {
	var payments []entities.Payment
	if err := r.db.Where("status = ?", "paid").Preload("Bill").Preload("Bill.Customer").Find(&payments).Error; err != nil {
		return nil, err
	}
	return payments, nil
}

func (r *adminRepository) GetAllPaymentBills() ([]entities.Payment, error) {
	var payments []entities.Payment
	if err := r.db.Preload("Bill").Preload("Bill.Customer").Find(&payments).Error; err != nil {
		return nil, err
	}
	return payments, nil
}
