package repository

import (
	"encoding/json"
	"errors"
	"os"

	"go-merchants/src/model"
)

const merchantFile = "data/merchants.json"

// FetchMerchants - Load all merchants from merchants.json
func FetchMerchants() ([]model.Merchant, error) {
    var merchants []model.Merchant

    file, err := os.OpenFile(merchantFile, os.O_RDONLY|os.O_CREATE, 0644)
    if err != nil {
        return nil, err
    }
    defer file.Close()

    decoder := json.NewDecoder(file)
    err = decoder.Decode(&merchants)
    if err != nil && err.Error() != "EOF" {
        return nil, err
    }

    return merchants, nil
}

// SaveMerchants - Save all merchants back to merchants.json
func SaveMerchants(merchants []model.Merchant) error {
    bytes, err := json.MarshalIndent(merchants, "", "  ")
    if err != nil {
        return err
    }

    err = os.WriteFile(merchantFile, bytes, 0644)
    if err != nil {
        return err
    }

    return nil
}

// GetMerchantByID - Find one merchant by ID
func GetMerchantByID(merchantID string) (*model.Merchant, error) {
    merchants, err := FetchMerchants()
    if err != nil {
        return nil, err
    }

    for _, m := range merchants {
        if m.ID == merchantID {
            return &m, nil
        }
    }

    return nil, errors.New("merchant not found")
}
