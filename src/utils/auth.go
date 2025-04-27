package utils

import (
	"errors"
	"go-merchants/src/model"
)

// AuthenticateCustomer - mengecek username dan password (pakai bcrypt compare)
func AuthenticateCustomer(customers []model.Customer, username, password string) (*model.Customer, error) {
    for _, customer := range customers {
        if customer.Username == username {
            if CheckPasswordHash(password, customer.Password) {
                return &customer, nil
            }
            return nil, errors.New("invalid password")
        }
    }
    return nil, errors.New("customer not found")
}
