package helpers

import (
	"fmt"
	"math/rand"
	"time"
)

func GenerateBookCode() string {
	now := time.Now()
	nowFormat := now.Format("20060102150405")
	randomNumber := rand.Intn(20)
	bookCode := fmt.Sprintf("%s%d", nowFormat, randomNumber)
	return bookCode
}
