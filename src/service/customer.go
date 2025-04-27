package service

import (
	"errors"
	"go-merchants/src/model"
	"go-merchants/src/repository"
	"go-merchants/src/utils"
)

type CustomerService struct{}

func NewCustomerService() *CustomerService {
    return &CustomerService{}
}

// Login - Authenticate customer dan set IsLogged = true
func (cs *CustomerService) Login(username, password string) (string, error) {
    customers, err := repository.FetchCustomers()
    if err != nil {
        return "", err
    }
	
    authenticatedCustomer, err := utils.AuthenticateCustomer(customers, username, password)
    if err != nil {
		repository.SaveHistory(authenticatedCustomer.ID, "login", "", 0, false)
		return "", err
    }

	_ = repository.SaveHistory(authenticatedCustomer.ID, "login", "", 0, true)

    // Set IsLogged = true
    for i, c := range customers {
        if c.ID == authenticatedCustomer.ID {
            customers[i].IsLogged = true
            break
        }
    }

    // Save updated customers.json
    err = repository.SaveCustomers(customers)

    if err != nil {
        return "", err
    }
    token, err := utils.GenerateJWT(authenticatedCustomer.ID)
    if err != nil {
        return "", err
    }

    return token, nil
}



// Logout - Set IsLogged menjadi false
func (cs *CustomerService) Logout(customerID string) error {
    customers, err := repository.FetchCustomers()
    if err != nil {
        return err
    }

    found := false
    for i, c := range customers {
        if c.ID == customerID {
            customers[i].IsLogged = false
            found = true
            break
        }
    }

    if !found {
        return errors.New("customer not found")
    }

    err = repository.SaveCustomers(customers)
    if err != nil {
        return err
    }

    return nil
}

// GetCustomerByID - Cari customer berdasarkan ID
func (cs *CustomerService) GetCustomerByID(customerID string) (*model.Customer, error) {
    customers, err := repository.FetchCustomers()
    if err != nil {
        return nil, err
    }

    for _, c := range customers {
        if c.ID == customerID {
            return &c, nil
        }
    }

    return nil, errors.New("customer not found")
}
