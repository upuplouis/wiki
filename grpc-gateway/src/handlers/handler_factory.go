package handlers

import (
	"context"
	"protos"
)

type Service struct {}

func NewService() *Service{
	return new(Service)
}

func (m *Service)GetAll(ctx context.Context, in *protos.Null) (*protos.Null, error) {
	panic("")
}
