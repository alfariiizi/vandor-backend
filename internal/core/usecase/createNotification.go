package usecase

import (
	"context"
	"fmt"
	"time"

	domain_entries "github.com/alfariiizi/vandor/internal/core/domain"
	"github.com/alfariiizi/vandor/internal/core/model"
	"github.com/alfariiizi/vandor/internal/enum"
	"github.com/alfariiizi/vandor/internal/infrastructure/db"
	"github.com/alfariiizi/vandor/internal/infrastructure/db/notification"
	"github.com/alfariiizi/vandor/internal/infrastructure/sse"
	"github.com/alfariiizi/vandor/internal/pkg/logger"
	"github.com/alfariiizi/vandor/internal/pkg/validator"
	"github.com/alfariiizi/vandor/internal/utils"
)

type CreateNotificationInput struct {
	RecipientID  string `validate:"required"`
	Title        string `validate:"required"`
	Message      string
	Type         notification.Type     // info|success|warning|error
	Priority     notification.Priority // optional
	Channel      notification.Channel  // optional
	Event        enum.Event
	Link         *string
	Action       *string
	ResourceType *string
	ResourceID   *string
	Meta         map[string]any
	DedupeKey    *string
}
type CreateNotificationOutput struct {
	ID       string  `json:"id"`
	Message  string  `json:"message"`
	Priority string  `json:"priority"`
	Type     string  `json:"type"`
	Action   *string `json:"action,omitempty"`
}
type CreateNotification model.Usecase[CreateNotificationInput, CreateNotificationOutput]

type createNotification struct {
	client    *db.Client
	domain    *domain_entries.Domain
	validator validator.Validator
	sse       *sse.Manager
}

func NewCreateNotification(
	client *db.Client,
	domain *domain_entries.Domain,
	validator validator.Validator,
	sse *sse.Manager,
) CreateNotification {
	return &createNotification{
		client:    client,
		domain:    domain,
		validator: validator,
		sse:       sse,
	}
}

func (uc *createNotification) Validate(input CreateNotificationInput) error {
	return uc.validator.Validate(input)
}

func (uc *createNotification) Execute(ctx context.Context, input CreateNotificationInput) (*CreateNotificationOutput, error) {
	log := logger.Get()

	if err := uc.Validate(input); err != nil {
		return nil, err
	}
	res, err := uc.Process(ctx, input)
	if err != nil {
		log.Error().
			Str("usecase", "CreateNotification").
			Str("process", "CreateNotification").
			Str("error", err.Error()).
			Msg("Failed to process CreateNotification")
		return nil, err
	}

	if err := uc.Observer(ctx, input); err != nil {
		log.Error().
			Str("usecase", "CreateNotification").
			Str("observer", "CreateNotification").
			Str("error", err.Error()).
			Msg("Failed to observe CreateNotification")
		return nil, err
	}

	go func() {
		if err := uc.SendEvent(ctx, input, *res); err != nil {
			log.Error().
				Str("usecase", "CreateNotification").
				Str("send_event", "CreateNotification").
				Str("error", err.Error()).
				Msg("Failed to send event in CreateNotification")
		}
	}()

	return res, nil
}

func (uc *createNotification) Observer(ctx context.Context, input CreateNotificationInput) error {
	// TODO: Implement observer logic
	// This is optional. You can leave this blank if not needed.

	return nil
}

func (uc *createNotification) SendEvent(ctx context.Context, input CreateNotificationInput, output CreateNotificationOutput) error {
	uc.sse.PublishToUser(input.RecipientID, sse.Event{
		ID:    fmt.Sprintf("notif-%d", time.Now().UnixNano()),
		Event: input.Event.Label(),
		Data:  output,
	})

	return nil
}

func (uc *createNotification) Process(ctx context.Context, input CreateNotificationInput) (*CreateNotificationOutput, error) {
	notif, err := utils.WithTxResult(ctx, uc.client, func(tx *db.Tx) (*db.Notification, error) {
		recipientID, err := utils.IDParser(input.RecipientID)
		if err != nil {
			return nil, err
		}

		b := tx.Notification.Create().
			SetUserID(*recipientID).
			SetTitle(input.Title).
			SetMessage(input.Message).
			SetType(input.Type).
			SetPriority("normal").
			SetChannel("in_app")
		if input.Priority != "" {
			b.SetPriority(input.Priority)
		}
		if input.Channel != "" {
			b.SetChannel(input.Channel)
		}
		if input.Link != nil {
			b.SetLink(*input.Link)
		}
		if input.Action != nil {
			b.SetAction(*input.Action)
		}
		if input.ResourceType != nil {
			b.SetResourceType(*input.ResourceType)
		}
		if input.ResourceID != nil {
			b.SetResourceID(*input.ResourceID)
		}
		if input.Meta != nil {
			b.SetMeta(input.Meta)
		}
		if input.DedupeKey != nil {
			b.SetDedupeKey(*input.DedupeKey)
		}

		notif, err := b.Save(ctx)
		if err != nil {
			return nil, err
		}

		return notif, nil
	})
	if err != nil {
		return nil, err
	}

	return &CreateNotificationOutput{
		ID:       notif.ID.String(),
		Message:  "Succesfully create notification",
		Priority: notif.Priority.String(),
		Type:     notif.Type.String(),
		Action:   notif.Action,
	}, nil
}
