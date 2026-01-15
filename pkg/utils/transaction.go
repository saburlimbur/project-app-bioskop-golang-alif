package utils

import (
	"fmt"
	"math/rand"
	"time"
)

func GenerateTransactionID() string {
	rand.Seed(time.Now().UnixNano())

	timestamp := time.Now().Format("20060102150405") // YYYYMMDDHHMMSS
	random := rand.Intn(900000) + 100000             // 6 digit random

	return fmt.Sprintf("TRX-%s-%d", timestamp, random)
}
