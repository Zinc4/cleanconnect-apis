package admin

import (
	"clean-connect/config"
	"log"
	"time"
)

type Scheduler struct {
	service AdminService
}

func NewScheduler(service AdminService) *Scheduler {
	return &Scheduler{service}
}

func (s *Scheduler) Start() {
	go func() {
		for {
			s.checkOverdueBills()
			time.Sleep(24 * time.Hour)
		}
	}()
}

func (s *Scheduler) checkOverdueBills() error {
	bills, err := s.service.GetBills()
	if err != nil {
		log.Println("Error fetching bills:", err)
		return err
	}

	for _, bill := range bills {
		if bill.Status == "Belum Dibayar" && bill.BillDue.Before(time.Now()) {
			bill.Status = "Terlambat"
			if err := s.service.UpdateBill(&bill); err != nil {
				log.Println("Error updating bill status:", err)
			}
		} else if time.Until(bill.BillDue).Hours() < 48 {
			return config.SendNotification(bill.Customer.Email, bill.Description, bill.BillDue)
		}
	}
	return nil
}
