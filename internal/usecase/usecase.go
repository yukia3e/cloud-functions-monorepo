package usecase

import "context"

type Usecase interface {
	FunctionA(ctx context.Context, id string) error
	FunctionB(ctx context.Context, id string) error
}

type usecase struct{}

func New() Usecase {
	return &usecase{}
}
