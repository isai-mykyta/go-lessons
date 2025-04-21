package validators

import (
	"testing"
	"time"
)

func assertEqual[T comparable](expect T, recieve T, t *testing.T) {
	if expect != recieve {
		t.Helper()
		t.Errorf("Expected %v but got %v", expect, recieve)
	}
}

func TestValidateEmail(t *testing.T) {
	assertEqual(true, ValidateEmail("example@mail.com"), t)
	assertEqual(false, ValidateEmail("invalidemail.com"), t)
}

func TestValidatePhoneNumber(t *testing.T) {
	assertEqual(true, ValidatePhoneNumber("+380000000000"), t)
	assertEqual(false, ValidatePhoneNumber("1234"), t)
}

func TestValidateCountryCode(t *testing.T) {
	assertEqual(true, ValidateCountryCode("UA"), t)
	assertEqual(false, ValidateCountryCode("JJ"), t)
}

func TestValidatePostalCode(t *testing.T) {
	assertEqual(true, ValidatePostalCode(345500), t)
	assertEqual(false, ValidatePostalCode(00), t)
}

func TestValidateTimestampt(t *testing.T) {
	assertEqual(true, ValidateTimestampt(time.Now().Unix()), t)
	assertEqual(false, ValidateTimestampt(time.Date(1900, time.January, 1, 0, 0, 0, 0, time.UTC).Unix()), t)
}
