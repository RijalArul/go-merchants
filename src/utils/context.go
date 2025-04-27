package utils

import (
	"context"
)

type contextKey string

const customerIDKey = contextKey("customer_id")

func InjectCustomerID(ctx context.Context, customerID string) context.Context {
    return context.WithValue(ctx, customerIDKey, customerID)
}

func GetCustomerID(ctx context.Context) (string, bool) {
    customerID, ok := ctx.Value(customerIDKey).(string)
    return customerID, ok
}