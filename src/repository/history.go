package repository

import (
	"encoding/json"
	"go-merchants/src/model"
	"os"
	"time"

	"github.com/google/uuid"
)

func SaveHistory(customerID, action, merchantID string, amount int, success bool) error {
	// Create history record
	history := model.History{
		ID:         uuid.New().String(), // generate random unique ID
		CustomerID: customerID,
		Action:     action,
		MerchantID: merchantID,
		Amount:     amount,
		Timestamp:  time.Now().Format(time.RFC3339),
		Success:   success,
	}

	var histories []model.History

	// Try open history file
	file, err := os.OpenFile("data/history.json", os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	_ = decoder.Decode(&histories) // kalau gagal decode, artinya file kosong, gapapa

	histories = append(histories, history)

	bytes, err := json.MarshalIndent(histories, "", "  ")
	if err != nil {
		return err
	}

	err = os.WriteFile("data/history.json", bytes, 0644)
	if err != nil {
		return err
	}

	return nil
}
