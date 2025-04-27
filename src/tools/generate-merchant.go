package tools

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"os"

	"go-merchants/src/model"
	"go-merchants/src/utils"

	"github.com/bxcodec/faker/v4"
)

func RandomAccountNumber() string {
    // Generate account number random 10 digit
    return fmt.Sprintf("%010d", rand.Intn(10000000000))
}

func GenerateMerchants() {
    var merchants []model.Merchant

    for i := 0; i < 5; i++ { // generate 5 merchants, bisa diubah sesuai kebutuhan
        plainAccountNumber := RandomAccountNumber()

        encryptedAccount, err := utils.Encrypt(plainAccountNumber)
        if err != nil {
            log.Fatalf("Failed to encrypt account number: %v", err)
        }

        merchant := model.Merchant{
            ID:            fmt.Sprintf("merchant-%03d", i+1),
            Name:          faker.DomainName(),
            AccountNumber: encryptedAccount, // save yang sudah dienkripsi
        }
        merchants = append(merchants, merchant)
    }

    // Pastikan folder data/ ada
    err := os.MkdirAll("data", os.ModePerm)
    if err != nil {
        log.Fatal(err)
    }

    file, err := os.Create("data/merchants.json")
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    encoder := json.NewEncoder(file)
    encoder.SetIndent("", "  ")
    err = encoder.Encode(merchants)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Println("✅ Successfully generated merchants.json with encrypted account numbers 🚀")
}
