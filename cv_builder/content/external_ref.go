package content

import "errors"

type externalRef struct {
	Title string `json:"title"`
	Url   string `json:"url"`
}

func NewExternalRef(title, url string) (*externalRef, error) {
	if title == "" || url == "" {
		return nil, errors.New("title and url must be provided")
	}

	return &externalRef{Title: title, Url: url}, nil
}
