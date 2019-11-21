package main

import (
	"context"

	"github.com/google/wire"
	"github.com/wxb/goLab/pkg/wire/guide/foobarbaz"
)

func initializeBaz(ctx context.Context) (foobarbaz.Baz, error) {
	wire.Build(foobarbaz.SuperSet)
	return foobarbaz.Baz{}, nil
}
