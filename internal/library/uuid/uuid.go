package uuid

import (
	"github.com/google/uuid"
)

func GetUuid() (string, error) {
	// Creating UUID Version 4
	// panic on error
	_uuid, _ := uuid.NewRandom()
	return _uuid.String(), nil
}
