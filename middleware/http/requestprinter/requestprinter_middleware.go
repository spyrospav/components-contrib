package requestprinter

import (
	"github.com/dapr/components-contrib/middleware"
	"github.com/dapr/kit/logger"
	"github.com/valyala/fasthttp"
)

// NewRequestPrinterMiddleware returns a new requestprinter middleware.
func NewRequestPrinterMiddleware(log logger.Logger) middleware.Middleware {
	return &Middleware{logger: log}
}

// Middleware is an requestprinter middleware.
type Middleware struct {
	logger logger.Logger
}

// GetHandler returns the HTTP handler provided by the middleware.
func (m *Middleware) GetHandler(metadata middleware.Metadata) (func(h fasthttp.RequestHandler) fasthttp.RequestHandler, error) {
	return func(h fasthttp.RequestHandler) fasthttp.RequestHandler {
		return func(ctx *fasthttp.RequestCtx) {
			m.logger.Info("Request Printer: Before request")
			h(ctx)
			m.logger.Info("Request Printer: After request")
		}
	}, nil
}
