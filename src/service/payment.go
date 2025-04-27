package service

import (
	"errors"
	"go-merchants/src/repository"
)

type PaymentService struct{}

func NewPaymentService() *PaymentService {
    return &PaymentService{}
}

// Pay - Proses pembayaran dari customer ke merchant
func (ps *PaymentService) Pay(customerID, merchantID string, amount int) error {
    // Fetch semua customer
    customers, err := repository.FetchCustomers()
    if err != nil {
        return err
    }

    // Fetch merchant
    merchant, errMerchant := repository.GetMerchantByID(merchantID)

    if errMerchant != nil {
		repository.SaveHistory(customerID, "payment", merchantID, amount, false)
        return errMerchant
    }

    // Cari customer yang login
    var found bool
    for i := range customers {
        if customers[i].ID == customerID {
            found = true

            // Cek apakah logged in
            if !customers[i].IsLogged {
                return errors.New("customer not logged in")
            }

            // Cek saldo cukup
            if customers[i].Balance < amount {
                return errors.New("insufficient balance")
            }

            // Proses potong saldo
            customers[i].Balance -= amount
            break
        }
    }

    if !found {
        return errors.New("customer not found")
    }

    // Save updated customers
    err = repository.SaveCustomers(customers)
    if err != nil {

        return err
    }

    // Save payment history
    err = repository.SaveHistory(customerID, "payment", merchant.ID, amount, true)
    if err != nil {
        return err
    }

    return nil
}
