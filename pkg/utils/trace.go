package utils

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
)

func StartFiberTrace(c *fiber.Ctx, spanName string) (context.Context, trace.Span) {
	ctx := context.WithValue(c.Context(), fiber.HeaderXRequestID, c.GetRespHeader(fiber.HeaderXRequestID))

	ctx, span := otel.Tracer("").Start(ctx, spanName)
	span.SetAttributes(attribute.String(fiber.HeaderXRequestID, c.GetRespHeader(fiber.HeaderXRequestID)))

	return ctx, span
}
