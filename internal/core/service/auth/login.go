package auth_service

import (
	"context"
	"fmt"
	"time"

	"github.com/alfariiizi/vandor/internal/config"
	domain_entries "github.com/alfariiizi/vandor/internal/core/domain"
	"github.com/alfariiizi/vandor/internal/core/model"
	"github.com/alfariiizi/vandor/internal/core/usecase"
	"github.com/alfariiizi/vandor/internal/infrastructure/db"
	"github.com/alfariiizi/vandor/internal/infrastructure/db/user"
	"github.com/alfariiizi/vandor/internal/pkg/validator"
	"github.com/alfariiizi/vandor/internal/utils"
)

type LoginInput struct {
	Email        string `json:"email" validate:"required,email" doc:"Email address of the user"`
	Password     string `json:"password" validate:"required,min=6" doc:"Password of the user, minimum length is 8 characters"`
	IsBackoffice bool   `json:"is_backoffice" doc:"Is the user logging in from backoffice? Set to true if the user is logging in from backoffice"`
}
type LoginOutput struct {
	AccessToken           string `json:"access_token"`
	RefreshToken          string `json:"refresh_token"`
	AccessTokenExpiresAt  int64  `json:"access_token_expires_at"`
	RefreshTokenExpiresAt int64  `json:"refresh_token_expires_at"`
}
type Login model.Service[LoginInput, LoginOutput]

type login struct {
	domain    *domain_entries.Domain
	client    *db.Client
	usecase   *usecase.Usecases
	validator validator.Validator
}

func NewLogin(
	domain *domain_entries.Domain,
	client *db.Client,
	usecase *usecase.Usecases,
	validator validator.Validator,
) Login {
	return &login{
		domain:    domain,
		client:    client,
		usecase:   usecase,
		validator: validator,
	}
}

func (s *login) Validate(input LoginInput) error {
	return s.validator.Validate(input)
}

func (s *login) Execute(ctx context.Context, input LoginInput) (*LoginOutput, error) {
	if err := s.Validate(input); err != nil {
		return nil, err
	}
	if err := s.Observer(ctx, input); err != nil {
		fmt.Println("Observer error:", err)
	}
	return s.Process(ctx, input)
}

func (s *login) Observer(ctx context.Context, input LoginInput) error {
	return nil
}

func (s *login) Process(ctx context.Context, input LoginInput) (*LoginOutput, error) {
	userOne, err := s.domain.User.One(
		s.client.User.Query().
			Where(user.Email(input.Email)).
			Only(ctx),
	)
	if err != nil {
		return nil, err
	}
	if userOne == nil {
		return nil, fmt.Errorf("user not found with email: %s", input.Email)
	}
	if input.IsBackoffice && !userOne.IsAllowedToLoginBackoffice() {
		return nil, fmt.Errorf("user with email %s is not allowed to login to backoffice", input.Email)
	}
	if !userOne.IsPasswordMatches(input.Password) {
		return nil, fmt.Errorf("invalid credentials for user: %s", input.Email)
	}

	refreshToken, err := utils.GenerateRefreshToken()
	if err != nil {
		return nil, fmt.Errorf("failed to generate refresh token: %w", err)
	}

	session, err := s.domain.Session.One(
		s.client.Session.Create().
			SetRefreshToken(refreshToken).
			SetUserID(userOne.ID).
			SetExpiresAt(time.Now().Add(time.Duration(config.GetConfig().Auth.SessionDurationInDays) * time.Hour * 24)).
			Save(ctx),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create session: %w", err)
	}

	accessToken, err := utils.GenerateAccessToken(
		userOne.ID.String(),
		session.ID.String(),
		userOne.FullName(),
		userOne.Email,
		userOne.Role.String(),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to generate access token: %w", err)
	}

	return &LoginOutput{
		AccessToken:           accessToken.Token,
		RefreshToken:          session.RefreshToken,
		AccessTokenExpiresAt:  accessToken.ExpiresAt,
		RefreshTokenExpiresAt: session.ExpiresAt.Unix(),
	}, nil
}
