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


func GenerateCustomers() {
    var customers []model.Customer

    for i := 0; i < 10; i++ { // misal mau generate 10 customer
        passwordPlain := "password123" // password default
        hashedPassword, err := utils.HashPassword(passwordPlain)
        if err != nil {
            log.Fatal(err)
        }

        customer := model.Customer{
            ID:       fmt.Sprintf("customer-%03d", i+1),
            Username: faker.Username(),
            Password: hashedPassword,
            Balance:  rand.Intn(1000000), // saldo random 0 - 999999
            IsLogged: false,
        }
        customers = append(customers, customer)
    }

    err := os.MkdirAll("data", os.ModePerm) // Pastikan folder data/ ada
    if err != nil {
        log.Fatal(err)
    }

    file, err := os.Create("data/customers.json")
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    encoder := json.NewEncoder(file)
    encoder.SetIndent("", "  ")
    err = encoder.Encode(customers)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Println("✅ Successfully generated customers.json 🚀")
}
