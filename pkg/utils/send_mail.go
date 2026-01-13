package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"go.uber.org/zap"
)

type EmailRequest struct {
	To      string `json:"to"`
	Name    string `json:"name"`
	Subject string `json:"subject"`
	Message string `json:"body"`
}

type EmailSender struct {
	Config EmailConfig
	Logger *zap.Logger
}

func NewEmailSender(cfgMail EmailConfig, logger *zap.Logger) *EmailSender {
	return &EmailSender{
		Config: cfgMail,
		Logger: logger,
	}
}

func (e *EmailSender) SendEmail(to, name, subject, message string) error {
	payload := EmailRequest{
		To:      to,
		Name:    name,
		Subject: subject,
		Message: message,
	}

	body, err := json.Marshal(payload)
	if err != nil {
		e.Logger.Error("marshal email payload failed", zap.Error(err))
		return err
	}

	req, err := http.NewRequest(
		http.MethodPost,
		e.Config.RequestURL,
		bytes.NewBuffer(body),
	)
	if err != nil {
		e.Logger.Error("create email request failed", zap.Error(err))
		return err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("x-api-key", e.Config.APIKey)

	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	resp, err := client.Do(req)
	if err != nil {
		e.Logger.Error("send email failed", zap.Error(err))
		return err
	}
	defer resp.Body.Close()

	respBody, _ := io.ReadAll(resp.Body)

	if resp.StatusCode != http.StatusOK {
		e.Logger.Error(
			"email api error",
			zap.Int("status", resp.StatusCode),
			zap.String("response", string(respBody)),
		)
		return fmt.Errorf("email api error: %s", resp.Status)
	}

	e.Logger.Info(
		"email sent",
		zap.String("to", to),
		zap.String("subject", subject),
	)

	return nil
}
