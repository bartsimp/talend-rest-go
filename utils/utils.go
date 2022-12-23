package utils

import (
	"fmt"

	"github.com/bartsimp/talend-rest-go/models"
)

func UnmarshalErrorResponse(e *models.ErrorResponse) string {
	message := ""
	if e.Message != nil {
		message = *e.Message
	}
	status := ""
	if e.Status != nil {
		status = fmt.Sprint(*e.Status)
	}
	return fmt.Sprintf(`{"code":"%s","details":"%s","message":"%s","requestId":"%s","status":%s,"url":"%s"}`,
		e.Code,
		e.Details,
		message,
		e.RequestID,
		status,
		e.URL)
}
