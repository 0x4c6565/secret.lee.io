package storage

import (
	"context"
)

type Storage interface {
	Get(ctx context.Context, uuid string) (string, error)
	Set(ctx context.Context, uuid string, content string) error
	Delete(ctx context.Context, uuid string) error
}
