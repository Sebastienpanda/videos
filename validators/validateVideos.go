package validators

import (
	"errors"
)

func ValidateVideos(title string, content string) error {
	if title == "" {
		return errors.New("title is required")
	}
	if content == "" {
		return errors.New("content is required")
	}
	return nil
}
