package admin

import "clean-connect/pkg/entities"

type AdminService interface {
	CreateBill(bill entities.Bill) error
	CreateMassBill(bills []entities.Bill) error
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

type adminService struct {
	repository AdminRepository
}

func NewAdminService(repository AdminRepository) *adminService {
	return &adminService{repository}
}

func (s *adminService) CreateBill(bill entities.Bill) error {
	return s.repository.CreateBill(bill)
}

func (s *adminService) CreateMassBill(bills []entities.Bill) error {
	return s.repository.CreateBills(bills)
}

func (s *adminService) CreateAdditionalBill(additionalBill entities.AdditionalBill) error {
	return s.repository.CreateAdditionalBill(additionalBill)
}

func (s *adminService) GetBills() ([]entities.Bill, error) {
	return s.repository.GetBills()
}

func (s *adminService) GetBill(id uint) (entities.Bill, error) {
	return s.repository.GetBill(id)
}

func (s *adminService) GetAdditionalBills() ([]entities.AdditionalBill, error) {
	return s.repository.GetAdditionalBills()
}

func (s *adminService) GetAdditionalBill(id uint) (entities.AdditionalBill, error) {
	return s.repository.GetAdditionalBill(id)
}

func (s *adminService) DeleteBill(id uint) error {
	return s.repository.DeleteBill(id)
}

func (s *adminService) UpdateBill(bill *entities.Bill) error {
	return s.repository.UpdateBill(bill)
}

func (s *adminService) UpdatePayment(payment *entities.Payment) error {
	return s.repository.UpdatePayment(payment)
}

func (s *adminService) GetPaymentByBillIDAndCustomerID(billID uint, customerID uint) (entities.Payment, error) {
	return s.repository.GetPaymentByBillIDAndCustomerID(billID, customerID)
}

func (s *adminService) GetPendingUsersPayments() ([]entities.Payment, error) {
	return s.repository.GetPendingUsersPayments()
}

func (s *adminService) GetSuccessPayments() ([]entities.Payment, error) {
	return s.repository.GetSuccessPayments()
}

func (s *adminService) GetAllPaymentBills() ([]entities.Payment, error) {
	return s.repository.GetAllPaymentBills()
}
