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
	GetTotalDashboard() (int64, int64, int64, error)
	GetTotalBillsManagement() (int64, int64, int64, int64, error)
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

func (s *adminService) GetTotalDashboard() (int64, int64, int64, error) {
	totalUser, err := s.repository.GetTotalUsers()
	if err != nil {
		return 0, 0, 0, err
	}

	totalPendingPayment, err := s.repository.GetTotalPendingPayment()
	if err != nil {
		return 0, 0, 0, err
	}

	totalSuccessPayment, err := s.repository.GetTotalAmountSuccessPayment()
	if err != nil {
		return 0, 0, 0, err
	}

	return totalUser, totalPendingPayment, totalSuccessPayment, nil
}

func (s *adminService) GetTotalBillsManagement() (int64, int64, int64, int64, error) {
	totalBills, err := s.repository.GetTotalBills()
	if err != nil {
		return 0, 0, 0, 0, err
	}

	totalAmountBills, err := s.repository.GetTotalAmountBills()
	if err != nil {
		return 0, 0, 0, 0, err
	}

	totalSuccessPayment, err := s.repository.GetTotalAmountSuccessPayment()
	if err != nil {
		return 0, 0, 0, 0, err
	}

	totalPendingPayment, err := s.repository.GetPendingAmountPayments()
	if err != nil {
		return 0, 0, 0, 0, err
	}

	return totalBills, totalAmountBills, totalSuccessPayment, totalPendingPayment, nil

}
