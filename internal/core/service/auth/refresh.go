package auth_service

import (
	"context"
	"fmt"

	domain_entries "github.com/alfariiizi/vandor/internal/core/domain"
	"github.com/alfariiizi/vandor/internal/core/model"
	"github.com/alfariiizi/vandor/internal/core/usecase"
	"github.com/alfariiizi/vandor/internal/infrastructure/db"
	"github.com/alfariiizi/vandor/internal/infrastructure/db/session"
	"github.com/alfariiizi/vandor/pkg/validator"
)

type RefreshInput struct {
	RefreshToken string `json:"refresh_token" validate:"required"`
}
type (
	RefreshOutput LoginOutput
	Refresh       model.Service[RefreshInput, RefreshOutput]
)

type refresh struct {
	domain    *domain_entries.Domain
	client    *db.Client
	usecase   *usecase.Usecases
	validator validator.Validator
}

func NewRefresh(
	domain *domain_entries.Domain,
	client *db.Client,
	usecase *usecase.Usecases,
	validator validator.Validator,
) Refresh {
	return &refresh{
		domain:    domain,
		client:    client,
		usecase:   usecase,
		validator: validator,
	}
}

func (s *refresh) Validate(input RefreshInput) error {
	return s.validator.Validate(input)
}

func (s *refresh) Execute(ctx context.Context, input RefreshInput) (*RefreshOutput, error) {
	if err := s.Validate(input); err != nil {
		return nil, err
	}
	return s.Process(ctx, input)
}

func (s *refresh) Process(ctx context.Context, input RefreshInput) (*RefreshOutput, error) {
	sessionOne, err := s.domain.Session.One(
		s.client.Session.Query().
			Where(session.RefreshToken(input.RefreshToken)).
			WithUser().
			Only(ctx),
	)
	if err != nil {
		return nil, err
	}
	if !sessionOne.IsAvailable() {
		return nil, fmt.Errorf("session expired or revoked")
	}

	_, err = s.client.Session.Update().AddNumberOfUses(1).Where(
		session.ID(sessionOne.ID),
	).Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to update session: %w", err)
	}

	userOne := s.domain.User.Convert(sessionOne.Edges.User)
	accessToken, err := s.usecase.CreateAccessToken.Execute(ctx, usecase.CreateAccessTokenInput{
		UserID:    sessionOne.UserID.String(),
		SessionID: sessionOne.ID.String(),
		Name:      userOne.FullName(),
		Email:     userOne.Email,
		Role:      userOne.Role.String(),
	})
	if err != nil {
		return nil, fmt.Errorf("failed to create access token: %w", err)
	}

	return &RefreshOutput{
		AccessToken:           accessToken.Value,
		AccessTokenExpiresAt:  accessToken.ExpiresAt,
		RefreshToken:          sessionOne.RefreshToken,
		RefreshTokenExpiresAt: sessionOne.ExpiresAt.Unix(),
	}, nil
}
