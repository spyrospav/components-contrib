package requestprinter

import (
	"encoding/json"
    "os"

	"github.com/valyala/fasthttp"
	"github.com/dapr/components-contrib/middleware"
	"github.com/dapr/kit/logger"
)

// Metadata is the requestprinter middleware config.
type requestprinterMiddlewareMetadata struct {
	// Path is where to save the printing of the request.
	Path string `json:"path"`
}

const (
	// todo: enter valid path
	defaultPath = "~/request.txt"
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
func (m *Middleware) GetHandler(metadata middleware.Metadata) (func(h fasthttp.RequestHandler) fasthttp.RequestHandler, error,) {

	meta, err := m.getNativeMetadata(metadata)
	if err != nil {
		return nil, err
	}
	m.logger.Info("requestprinter middleware loaded")
	m.logger.Info(meta.Path)

	return func(h fasthttp.RequestHandler) fasthttp.RequestHandler {
		return func(ctx *fasthttp.RequestCtx) {

			err := os.WriteFile(meta.Path, []byte(ctx.Request.String()), 0644)
			if err != nil {
				m.logger.Error(err)
				return
			}

			h(ctx)

		}
	}, nil
}

func (m *Middleware) getNativeMetadata(metadata middleware.Metadata) (*requestprinterMiddlewareMetadata, error) {
	var middlewareMetadata requestprinterMiddlewareMetadata

	b, err := json.Marshal(metadata.Properties)

	err = json.Unmarshal(b, &middlewareMetadata)
	
	if err != nil {
		return nil, err
	}

	if middlewareMetadata.Path == "" {
		middlewareMetadata.Path = defaultPath
	}

	return &middlewareMetadata, nil
}