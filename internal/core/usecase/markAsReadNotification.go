package usecase

import (
	"context"
	"log"

	domain_entries "github.com/alfariiizi/vandor/internal/core/domain"
	"github.com/alfariiizi/vandor/internal/core/model"
	"github.com/alfariiizi/vandor/internal/infrastructure/db"
	"github.com/alfariiizi/vandor/internal/pkg/validator"
)

type MarkAsReadNotificationInput struct {
	// TODO: Define fields
}
type MarkAsReadNotificationOutput struct {
	// TODO: Define fields
}
type MarkAsReadNotification model.Usecase[MarkAsReadNotificationInput, MarkAsReadNotificationOutput]

type markAsReadNotification struct {
	client    *db.Client
	domain    *domain_entries.Domain
	validator validator.Validator
}

func NewMarkAsReadNotification(
	client *db.Client,
	domain *domain_entries.Domain,
	validator validator.Validator,
) MarkAsReadNotification {
	return &markAsReadNotification{
		client:    client,
		domain:    domain,
		validator: validator,
	}
}

func (uc *markAsReadNotification) Validate(input MarkAsReadNotificationInput) error {
	return uc.validator.Validate(input)
}

func (uc *markAsReadNotification) Execute(ctx context.Context, input MarkAsReadNotificationInput) (*MarkAsReadNotificationOutput, error) {
	if err := uc.Validate(input); err != nil {
		return nil, err
	}
	if err := uc.Observer(ctx, input); err != nil {
		log.Printf("Observer usecase 'MarkAsReadNotification' error: %s", err.Error())
	}
	return uc.Process(ctx, input)
}

func (uc *markAsReadNotification) Observer(ctx context.Context, input MarkAsReadNotificationInput) error {
	// TODO: Implement observer logic
	// This is optional. You can leave this blank if not needed.

	return nil
}

func (uc *markAsReadNotification) Process(ctx context.Context, input MarkAsReadNotificationInput) (*MarkAsReadNotificationOutput, error) {
	// TODO: Implement logic

	return &MarkAsReadNotificationOutput{}, nil
}
