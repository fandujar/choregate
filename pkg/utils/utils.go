package utils

import "github.com/google/uuid"

func GenerateID() (uuid.UUID, error) {
	u, err := uuid.NewV7()
	if err != nil {
		return uuid.UUID{}, err
	}

	return u, nil
}
