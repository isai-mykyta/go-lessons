package content

import (
	"cv_builder/validators"
	"errors"
)

type experience struct {
	Posititon   string  `json:"position"`
	StartFrom   int     `json:"startFrom"`
	EndAt       *int    `json:"endAt,omitempty"`
	Company     string  `json:"company"`
	Description *string `json:"description,omitempty"`
}

func NewExperience(position, company string, description *string, startFrom int, endAt *int) (*experience, error) {
	if position == "" || company == "" {
		return nil, errors.New("position and company must be provided")
	}

	if !validators.ValidateTimestampt(int64(startFrom)) {
		return nil, errors.New("invalid startFrom value")
	}

	if endAt != nil && !validators.ValidateTimestampt(int64(*endAt)) {
		return nil, errors.New("invalid endAt value")
	}

	return &experience{
		Posititon:   position,
		Description: description,
		Company:     company,
		StartFrom:   startFrom,
		EndAt:       endAt,
	}, nil
}
