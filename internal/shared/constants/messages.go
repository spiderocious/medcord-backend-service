package constants

type MessageKey string

const (
	MsgSuccess             MessageKey = "SUCCESS"
	MsgInternalServerError MessageKey = "INTERNAL_SERVER_ERROR"
	MsgNotFound            MessageKey = "NOT_FOUND"
	MsgInvalidPayload      MessageKey = "INVALID_PAYLOAD"
	MsgValidationFailed    MessageKey = "VALIDATION_FAILED"
	MsgUnauthorized        MessageKey = "UNAUTHORIZED"
	MsgForbidden           MessageKey = "FORBIDDEN"
	MsgTokenRequired       MessageKey = "TOKEN_REQUIRED"
	MsgInvalidToken        MessageKey = "INVALID_TOKEN"
	MsgHealthOK            MessageKey = "HEALTH_OK"
)
