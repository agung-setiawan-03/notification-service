package services

import (
	"bytes"
	"context"
	"html/template"
	"notification-service/external"
	"notification-service/internal/interfaces"
	"notification-service/internal/models"

	"github.com/pkg/errors"
)

type EmailService struct {
	EmailRepo interfaces.IEmailRepo
}

func (s *EmailService) SendEmail(ctx context.Context, req models.InternalNotificationRequest) error {
	emailTemplate, err := s.EmailRepo.GetTemplate(ctx, req.TemplateName)
	if err != nil {
		return errors.Wrap(err, "failed to get template email")
	}

	tmpl, err := template.New("emailTemplate").Parse(emailTemplate.Body)
	if err != nil {
		return errors.Wrap(err, "failed to parse email template")
	}

	var tpl bytes.Buffer
	err = tmpl.Execute(&tpl, req.Placeholders)
	if err != nil {
		return errors.Wrap(err, "failed to execute the placeholder")
	}

	email := external.Email{
		To:      req.Recipient,
		Subject: emailTemplate.Subject,
		Body:    tpl.String(),
	}
	err = email.SendEmail()
	if err != nil {
		notifHistory := &models.NotificationHistory{
			Recipient:    req.Recipient,
			TemplateID:   emailTemplate.ID,
			Status:       "failed",
			ErrorMessage: err.Error(),
		}
		s.EmailRepo.InsertNotificationHistory(ctx, notifHistory)

		return errors.Wrap(err, "failed to send email")
	}

	notifHistory := &models.NotificationHistory{
		Recipient:  req.Recipient,
		TemplateID: emailTemplate.ID,
		Status:     "success",
	}
	s.EmailRepo.InsertNotificationHistory(ctx, notifHistory)

	return nil
}
