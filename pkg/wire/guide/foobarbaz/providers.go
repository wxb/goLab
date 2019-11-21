package foobarbaz

import "github.com/google/wire"

var SuperSet = wire.NewSet(ProvideFoo, ProvideBar, ProvideBaz)
