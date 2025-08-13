package utils

import (
	"context"
	"fmt"
)

func ExtractTenant(ctx context.Context) (string, error) {
	tenant, ok := ctx.Value("tenant").(string)
	if !ok || tenant == "" {
		return "", fmt.Errorf("tenant not found in context")
	}
	return tenant, nil
}
