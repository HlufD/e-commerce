package ports

import (
	"context"
)

type HttpRequester interface {
	Get(ctx context.Context, path string, result any) error
	Post(ctx context.Context, path string, body any, result any) error
}
