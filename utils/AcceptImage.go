package utils

import (
	"errors"
)

func AcceptImage(Typeimage string) error {
	var acceptedImages = map[string]struct{}{
		"image/png":  {},
		"image/jpg":  {},
		"image/jpeg": {},
	}

	if _, ok := acceptedImages[Typeimage]; !ok {
		return errors.New("unsupported image type")
	}
	return nil
}
