package middleware

import (
	"github.com/gofiber/contrib/fiberzerolog"
	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog/log"
)

// Logger initializes zerolog settings and returns configured middleware
func Logger() fiber.Handler {
	// Configure and return fiber-zerolog middleware
	return fiberzerolog.New(fiberzerolog.Config{
		Logger: &log.Logger,
		Fields: []string{
			fiberzerolog.FieldLatency,
			fiberzerolog.FieldStatus,
			fiberzerolog.FieldMethod,
			fiberzerolog.FieldPath,
			fiberzerolog.FieldIP,
			fiberzerolog.FieldUserAgent,
			fiberzerolog.FieldReqHeaders,
			fiberzerolog.FieldQueryParams,
			fiberzerolog.FieldBytesSent,
			fiberzerolog.FieldBytesReceived,
			fiberzerolog.FieldError,
		},
	})
}

// // ContextLogger gets the logger from context
// func ContextLogger(c *fiber.Ctx) zerolog.Logger {
// 	return log.Logger
// }
