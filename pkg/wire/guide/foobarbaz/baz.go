package foobarbaz

import (
	"context"
	"errors"
)

type Baz struct {
	X int
}

func ProvideBaz(ctx context.Context, bar Bar) (Baz, error) {
	if bar.X == 0 {
		return Baz{}, errors.New("cannot provide baz when bar is zero")
	}

	return Baz{X: bar.X}, nil
}
