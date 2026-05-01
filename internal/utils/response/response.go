package response

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/spiderocious/medcord-backend/internal/shared/constants"
)

type Envelope struct {
	Success bool   `json:"success"`
	Data    any    `json:"data,omitempty"`
	Error   string `json:"error,omitempty"`
	Message string `json:"message,omitempty"`
	Details any    `json:"details,omitempty"`
}

func detectLang(c *gin.Context) constants.Language {
	if l := c.Query("lang"); l != "" {
		return constants.LangOf(l)
	}
	if l := c.GetHeader("Accept-Language"); l != "" {
		return constants.LangOf(l)
	}
	return constants.LangEN
}

func Success(c *gin.Context, data any, key constants.MessageKey) {
	c.JSON(http.StatusOK, Envelope{
		Success: true,
		Data:    data,
		Message: constants.Translate(key, detectLang(c)),
	})
}

func Created(c *gin.Context, data any, key constants.MessageKey) {
	c.JSON(http.StatusCreated, Envelope{
		Success: true,
		Data:    data,
		Message: constants.Translate(key, detectLang(c)),
	})
}

func errorResp(c *gin.Context, status int, key constants.MessageKey, details any) {
	c.JSON(status, Envelope{
		Success: false,
		Error:   constants.Translate(key, detectLang(c)),
		Details: details,
	})
}

func BadRequest(c *gin.Context, key constants.MessageKey, details any) {
	errorResp(c, http.StatusBadRequest, key, details)
}

func Unauthorized(c *gin.Context, key constants.MessageKey) {
	errorResp(c, http.StatusUnauthorized, key, nil)
}

func Forbidden(c *gin.Context, key constants.MessageKey) {
	errorResp(c, http.StatusForbidden, key, nil)
}

func NotFound(c *gin.Context, key constants.MessageKey) {
	errorResp(c, http.StatusNotFound, key, nil)
}

func Conflict(c *gin.Context, key constants.MessageKey) {
	errorResp(c, http.StatusConflict, key, nil)
}

func ServerError(c *gin.Context, key constants.MessageKey) {
	errorResp(c, http.StatusInternalServerError, key, nil)
}
