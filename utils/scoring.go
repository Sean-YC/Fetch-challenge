package utils

import (
	"math"
	"regexp"
	"strconv"
	"strings"
	"time"

	"receipt-processor/models"
)

// Order retailer name
func countAlphanumeric(s string) int {
	re := regexp.MustCompile(`[a-zA-Z0-9]`)
	return len(re.FindAllString(s, -1))
}

// Check whole number
func isWholeNumber(total string) bool {
	f, err := strconv.ParseFloat(total, 64)
	return err == nil && f == float64(int(f))
}

// Check multiply of 0.25
func isMultipleOfQuarter(total string) bool {
	f, err := strconv.ParseFloat(total, 64)
	return err == nil && math.Mod(f, 0.25) == 0
}

// Calculate bonus
func descriptionBonus(item models.Item) int {
	if len(strings.TrimSpace(item.ShortDescription))%3 == 0 {
		price, _ := strconv.ParseFloat(item.Price, 64)
		return int(math.Ceil(price * 0.2))
	}
	return 0
}

// Calculate points for a receipt
func CalculatePoints(receipt models.Receipt) int {
	points := countAlphanumeric(receipt.Retailer)

	if isWholeNumber(receipt.Total) {
		points += 50
	}
	if isMultipleOfQuarter(receipt.Total) {
		points += 25
	}
	points += (len(receipt.Items) / 2) * 5

	for _, item := range receipt.Items {
		points += descriptionBonus(item)
	}

	purchaseDate, _ := time.Parse("2006-01-02", receipt.PurchaseDate)
	if purchaseDate.Day()%2 != 0 {
		points += 6
	}

	purchaseTime, _ := time.Parse("15:04", receipt.PurchaseTime)
	if purchaseTime.Hour() >= 14 && purchaseTime.Hour() < 16 {
		points += 10
	}

	return points
}
