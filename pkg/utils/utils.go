package utils

import "github.com/google/uuid"

func GenerateID() (uuid.UUID, error) {
	u, err := uuid.NewV7()
	if err != nil {
		return uuid.UUID{}, err
	}

	return u, nil
}

// GenerateShortID generates a short ID based on a UUID.
func GenerateShortID() (string, error) {
	u, err := uuid.NewV7()
	if err != nil {
		return "", err
	}

	shortID := u.String()[0:8]

	return shortID, nil
}
