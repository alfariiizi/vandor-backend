package usecase

import (
	"context"
	"fmt"
	"time"

	domain_entries "github.com/alfariiizi/vandor/internal/core/domain"
	"github.com/alfariiizi/vandor/internal/core/model"
	"github.com/alfariiizi/vandor/internal/infrastructure/db"
	"github.com/alfariiizi/vandor/internal/utils"
	"github.com/alfariiizi/vandor/pkg/validator"
)

type CreateAccessTokenInput struct {
	UserID    string `json:"user_id" validate:"required"`
	SessionID string `json:"session_id" validate:"required"`
	Name      string `json:"name" validate:"required"`
	Email     string `json:"email" validate:"required,email"`
	Role      string `json:"role" validate:"required"`
}

type CreateAccessTokenOutput struct {
	Value     string
	ExpiresAt int64
}
type CreateAccessToken model.Usecase[CreateAccessTokenInput, CreateAccessTokenOutput]

type createAccessToken struct {
	client    *db.Client
	domain    *domain_entries.Domain
	validator validator.Validator
}

func NewCreateAccessToken(
	client *db.Client,
	domain *domain_entries.Domain,
	validator validator.Validator,
) CreateAccessToken {
	return &createAccessToken{
		client:    client,
		domain:    domain,
		validator: validator,
	}
}

func (uc *createAccessToken) Validate(input CreateAccessTokenInput) error {
	return uc.validator.Validate(input)
}

func (uc *createAccessToken) Execute(ctx context.Context, input CreateAccessTokenInput) (*CreateAccessTokenOutput, error) {
	if err := uc.Validate(input); err != nil {
		return nil, err
	}
	return uc.Process(ctx, input)
}

func (uc *createAccessToken) Process(ctx context.Context, input CreateAccessTokenInput) (*CreateAccessTokenOutput, error) {
	accessTokenExpiresAt := time.Now().Add(time.Minute * 17)
	accessToken, err := utils.GenerateAccessToken(
		input.UserID,
		input.SessionID,
		input.Name,
		input.Email,
		input.Role,
		accessTokenExpiresAt,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to generate access token: %w", err)
	}

	return &CreateAccessTokenOutput{
		Value:     accessToken,
		ExpiresAt: accessTokenExpiresAt.Unix(),
	}, nil
}
