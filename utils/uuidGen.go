package utils

import "github.com/google/uuid"

func GenId() uuid.UUID {

	return uuid.New()
}
