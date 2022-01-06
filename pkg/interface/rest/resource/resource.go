package resource

import "github.com/google/wire"

var ResourceSet = wire.NewSet(NewUserResource)
