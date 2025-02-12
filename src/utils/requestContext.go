package utils

import (
	"context"
	"net/http"

	"github.com/chaiyapluek/goutils-logging/src/constant"
	"github.com/google/uuid"
	"github.com/labstack/echo"
)

type RequestContext interface {
	GetContext() context.Context
	GetHeader() http.Header
	GetLogContext() map[string]any
}

func extractLogInfo(header http.Header) map[string]any {
	logContext := make(map[string]any)
	cid := header.Get(constant.CorrelationId)
	if cid == "" {
		cid = uuid.NewString()
	}

	spanId := uuid.NewString()

	logContext[constant.CorrelationId] = cid
	logContext[constant.SpanId] = spanId

	return logContext
}

type EchoContext struct {
	echo.Context
	header     http.Header
	logContext map[string]any
}

func NewEchoContext(e echo.Context) *EchoContext {
	return &EchoContext{
		e,
		e.Request().Header,
		extractLogInfo(e.Request().Header),
	}
}

func (c *EchoContext) GetContext() context.Context {
	return c.Context.Request().Context()
}

func (c *EchoContext) GetHeader() http.Header {
	return c.header
}

func (c *EchoContext) GetLogContext() map[string]any {
	return c.logContext
}
