package utils

import (
	"fmt"

	"github.com/bartsimp/talend-rest-go/models"
)

func UnmarshalErrorResponse(e *models.ErrorResponse) string {
	return fmt.Sprintf(`{"code":"%s","details":"%s","message":"%s","requestId":"%s","status":%d,"url":"%s"}`,
		e.Code,
		e.Details,
		*e.Message,
		e.RequestID,
		*e.Status,
		e.URL)
}
