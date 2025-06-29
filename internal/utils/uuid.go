package utils

import (
	"fmt"

	"github.com/google/uuid"
)

func IDParser(id string) (*uuid.UUID, error) {
	parseId, err := uuid.Parse(id)
	if err != nil {
		return nil, fmt.Errorf("id is not uuid: %w", err)
	}

	return &parseId, nil
}
