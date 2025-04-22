package content

import (
	validators "cv_builder/validators"
	"errors"
)

type contacts struct {
	Email       string `json:"email,omitempty"`
	PhoneNumber string `json:"phone_number,omitempty"`
}

func NewContact(email, phoneNumber string) (*contacts, error) {
	if email == "" && phoneNumber == "" {
		return nil, errors.New("either email or phone number must be provided")
	}

	if email != "" && !validators.ValidateEmail(email) {
		return nil, errors.New("invalid email")
	}

	if phoneNumber != "" && !validators.ValidatePhoneNumber(phoneNumber) {
		return nil, errors.New("invalid phone number")
	}

	return &contacts{Email: email, PhoneNumber: phoneNumber}, nil
}
