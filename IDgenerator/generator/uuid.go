package generator

import (
	"github.com/google/uuid"
)

func getUUID() string {
	id := uuid.New()
	return id.String()
}
