package utils

import (
	"crypto/rand"
	"fmt"
)

func GenerateOTP() string {
	var otp int
	rand.Read([]byte{byte(otp)})
	return fmt.Sprintf("%06d", otp%1000000)
}
