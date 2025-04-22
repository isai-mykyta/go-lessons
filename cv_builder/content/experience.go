package content

import (
	"cv_builder/validators"
	"errors"
)

type experience struct {
	Posititon   string `json:"position"`
	StartFrom   int64  `json:"startFrom"`
	EndAt       int64  `json:"endAt,omitempty"`
	Company     string `json:"company"`
	Description string `json:"description,omitempty"`
}

func NewExperience(position, company, description string, startFrom, endAt int64) (*experience, error) {
	if position == "" || company == "" {
		return nil, errors.New("position and company must be provided")
	}

	if !validators.ValidateTimestampt(int64(startFrom)) {
		return nil, errors.New("invalid startFrom value")
	}

	if endAt != 0 && !validators.ValidateTimestampt(int64(endAt)) {
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
