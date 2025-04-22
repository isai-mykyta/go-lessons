package content

import (
	"testing"
	"time"
)

func assertEqual[T comparable](t *testing.T, expect, recieve T) {
	if expect != recieve {
		t.Helper()
		t.Errorf("Expected %v but got %v", expect, recieve)
	}
}

func TestNewAddressWithInvalidCountryCode(t *testing.T) {
	_, err := NewAddress("JJ", "", "", 39000)
	assertEqual(t, "invalid country code", err.Error())
}

func TestNewAddressWithInvalidPostalCode(t *testing.T) {
	_, err := NewAddress("UA", "", "", 11)
	assertEqual(t, "invalid postal code", err.Error())
}

func TestNewAddressWithValidData(t *testing.T) {
	address, err := NewAddress("UA", "", "", 39000)
	assertEqual(t, nil, err)
	assertEqual(t, "UA", address.CountryCode)
	assertEqual(t, 39000, address.PostalCode)
}

func TestNewContactsWithInvalidData(t *testing.T) {
	_, err := NewContact("", "")
	assertEqual(t, "either email or phone number must be provided", err.Error())
}

func TestNewContactsWithInvalidEmail(t *testing.T) {
	_, err := NewContact("example.com", "")
	assertEqual(t, "invalid email", err.Error())
}

func TestNewContactsWithInvalidPhoneNumber(t *testing.T) {
	_, err := NewContact("", "333")
	assertEqual(t, "invalid phone number", err.Error())
}

func TestNewContactsWithPhoneNumber(t *testing.T) {
	contacts, err := NewContact("", "+380000000000")
	assertEqual(t, nil, err)
	assertEqual(t, "+380000000000", contacts.PhoneNumber)
}

func TestNewContactsWithEmail(t *testing.T) {
	contacts, err := NewContact("example@mail.com", "")
	assertEqual(t, nil, err)
	assertEqual(t, "example@mail.com", contacts.Email)
}

func TestNewContactsWithEmailAndPhoneNumber(t *testing.T) {
	contacts, err := NewContact("example@mail.com", "+380000000000")
	assertEqual(t, nil, err)
	assertEqual(t, "example@mail.com", contacts.Email)
	assertEqual(t, "+380000000000", contacts.PhoneNumber)
}

func TestNewExperienceWithInvalidData(t *testing.T) {
	_, err := NewExperience("", "", "", 0, 0)
	assertEqual(t, "position and company must be provided", err.Error())
}

func TestNewExperienceWithInvalidCompany(t *testing.T) {
	_, err := NewExperience("engineer", "", "", 0, 0)
	assertEqual(t, "position and company must be provided", err.Error())
}

func TestNewExperienceWithInvalidPosition(t *testing.T) {
	_, err := NewExperience("", "Apple", "", 0, 0)
	assertEqual(t, "position and company must be provided", err.Error())
}

func TestNewExperienceWithInvalidStartFrom(t *testing.T) {
	_, err := NewExperience("engineer", "Apple", "", 0, 0)
	assertEqual(t, "invalid startFrom value", err.Error())
}

func TestNewExperienceWithInvalidEndAt(t *testing.T) {
	_, err := NewExperience("engineer", "Apple", "", time.Now().Unix(), 99999999999)
	assertEqual(t, "invalid endAt value", err.Error())
}

func TestNewExperienceWithDefaultEndAt(t *testing.T) {
	experience, _ := NewExperience("engineer", "Apple", "", time.Now().Unix(), 0)
	assertEqual(t, experience.EndAt, 0)
}

func TestNewExperienceWithValidData(t *testing.T) {
	startFrom := time.Now().Unix()
	endAt := time.Now().Unix()
	experience, _ := NewExperience("engineer", "Apple", "Lorem Ipsum", startFrom, endAt)

	assertEqual(t, experience.Posititon, "engineer")
	assertEqual(t, experience.Company, "Apple")
	assertEqual(t, experience.Description, "Lorem Ipsum")
	assertEqual(t, experience.StartFrom, startFrom)
	assertEqual(t, experience.EndAt, endAt)
}

func TestNewExternalRefWithInvalidTitle(t *testing.T) {
	_, err := NewExternalRef("", "http://127.0.0.1")
	assertEqual(t, "title and url must be provided", err.Error())
}

func TestNewExternalRefWithInvalidUrl(t *testing.T) {
	_, err := NewExternalRef("title", "invalid.url")
	assertEqual(t, "title and url must be provided", err.Error())
}

func TestNewExternalRefWithValidData(t *testing.T) {
	ref, _ := NewExternalRef("title", "https://example.com")
	assertEqual(t, "title", ref.Title)
	assertEqual(t, "https://example.com", ref.Url)
}
