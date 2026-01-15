package utils

import (
	"fmt"
	"math/rand"
	"time"
)

func GenerateBookingCode() string {
	rand.Seed(time.Now().UnixNano())
	return fmt.Sprintf("BK-%06d", rand.Intn(1000000))
}
