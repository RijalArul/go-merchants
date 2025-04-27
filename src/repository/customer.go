package repository

import (
	"encoding/json"
	"os"

	"go-merchants/src/model"
)

const customerFile = "data/customers.json"

// FetchCustomers - membaca semua customer dari file JSON
func FetchCustomers() ([]model.Customer, error) {
    bytes, err := os.ReadFile(customerFile)
    if err != nil {
        return nil, err
    }

    var customers []model.Customer
    err = json.Unmarshal(bytes, &customers)
    if err != nil {
        return nil, err
    }

    return customers, nil
}

func SaveCustomers(customers []model.Customer) error {
    bytes, err := json.MarshalIndent(customers, "", "  ")
    if err != nil {
        return err
    }

    err = os.WriteFile(customerFile, bytes, 0644)
    if err != nil {
        return err
    }

    return nil
}
