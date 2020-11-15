package ctxType

import (
	"context"
	"fghpdf.me/gin-project-template/internal/pkg/common"
)

type requestIdType string

var (
	requestId = requestIdType(common.REQUEST_ID)
)

// GetRequestId from context
// if null will return empty string
func GetRequestId(ctx context.Context) string {
	v := ctx.Value(requestId)
	if v == nil {
		return ""
	}

	return v.(string)
}

// WithRequestId into context
func WithRequestId(ctx context.Context, id string) context.Context {
	return context.WithValue(ctx, requestId, id)
}
