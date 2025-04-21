package content

import (
	validators "cv_builder/validators"
	"errors"
)

type address struct {
	CountryCode string  `json:"country_code"`
	City        *string `json:"city,omitempty"`
	PostalCode  *int    `json:"postal_code,omitempty"`
	Address     *string `json:"address,omitempty"`
}

func NewAddress(countryCode string, city, addressLine *string, postalCode *int) (*address, error) {
	if !validators.ValidateCountryCode(countryCode) {
		return nil, errors.New("invalid country code")
	}

	if postalCode != nil && !validators.ValidatePostalCode(*postalCode) {
		return nil, errors.New("invalid postal code")
	}

	return &address{
		CountryCode: countryCode,
		City:        city,
		Address:     addressLine,
		PostalCode:  postalCode,
	}, nil
}
