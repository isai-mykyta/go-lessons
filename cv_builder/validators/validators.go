package validators

import (
	"fmt"
	"regexp"
	"slices"
	"time"
)

func ValidateEmail(email string) bool {
	emailRegexp := `^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`
	return regexp.MustCompile(emailRegexp).MatchString(email)
}

func ValidatePhoneNumber(phoneNumber string) bool {
	phoneNumberRegexp := `^\+?[0-9]{10,15}$`
	return regexp.MustCompile(phoneNumberRegexp).MatchString(phoneNumber)
}

func ValidateCountryCode(code string) bool {
	allowedCountryCodes := []string{"UA", "US", "PL", "DK", "PL", "SK"}
	return slices.Contains(allowedCountryCodes, code)
}

func ValidatePostalCode(postalCode int) bool {
	postalCodeRegexp := `^\d{4,10}$`
	stringCode := fmt.Sprintf("%d", postalCode)
	return regexp.MustCompile(postalCodeRegexp).MatchString(stringCode)
}

func ValidateTimestampt(value int64) bool {
	t := time.Unix(value, 0)
	now := time.Now()
	past := time.Unix(0, 0)
	future := now.AddDate(10, 0, 0)

	return t.After(past) && t.Before(future)
}

func ValidateUrl(url string) bool {
	urlRegex := `^https?://[^\s/$.?#].[^\s]*$`
	return regexp.MustCompile(urlRegex).MatchString(url)
}
