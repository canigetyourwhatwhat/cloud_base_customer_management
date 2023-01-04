package service

import (
	"erply/entity"
	"erply/infra/database"
)

func (s *Service) CreateCustomer(filter map[string]string) error {
	err := database.InsertCustomer(s.db, filter)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) GetCustomerByCustomerID(customerId string) (*entity.Customer, error) {
	customer, err := database.GetCustomerByCustomerID(s.db, customerId)
	if err != nil {
		return nil, err
	}
	return customer, nil
}

func (s *Service) UpdateCustomerByCustomerID(filter map[string]string) error {
	err := database.UpdateCustomerByCustomerID(s.db, filter)
	if err != nil {
		return err
	}

	return nil
}

func (s *Service) DeleteCustomerByCustomerID(filter map[string]string) error {
	err := database.DeleteCustomerByCustomerID(s.db, filter["customerID"])
	if err != nil {
		return err
	}

	return nil
}
