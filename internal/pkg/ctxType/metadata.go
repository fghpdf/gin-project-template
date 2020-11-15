package ctxType

import (
	"context"
	"fghpdf.me/gin-project-template/internal/pkg/common"
	"google.golang.org/grpc/metadata"
	"strings"
)

// GetRequestIdFromMetadata return the request id in grpc metadata
// if no exist will return empty string
// the key will by lower in serialization
func GetRequestIdFromMetadata(ctx context.Context) string {
	// get data from metadata
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		md = metadata.Pairs()
	}

	// get request id
	requestIds := md.Get(strings.ToLower(common.REQUEST_ID_HEADER))
	if len(requestIds) >= 1 {
		return requestIds[0]
	}

	return ""
}

// WithRequestIdIntoMetadata with value into grpc metadata
func WithRequestIdIntoMetadata(ctx context.Context, id string) context.Context {
	return metadata.AppendToOutgoingContext(ctx, common.REQUEST_ID_HEADER, id)
}
