package service

import "context"

type Deleter interface {
	Delete(ctx context.Context, ID uint) error
}
