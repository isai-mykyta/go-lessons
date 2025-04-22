package content

import (
	"cv_builder/validators"
	"errors"
)

type externalRef struct {
	Title string `json:"title"`
	Url   string `json:"url"`
}

func NewExternalRef(title, url string) (*externalRef, error) {
	if title == "" || !validators.ValidateUrl(url) {
		return nil, errors.New("title and url must be provided")
	}

	return &externalRef{Title: title, Url: url}, nil
}
