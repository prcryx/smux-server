package wire

import (
	ri "smux/internals/data/repositories_impl"
	"smux/internals/domain/repositories"

	"github.com/google/wire"
)

var (
	BlogRepositorySet = wire.NewSet(
		ri.NewBlogRepositoryImpl,
		wire.Bind(new(repositories.BlogRepository), new(*ri.BlogRepositoryImpl)),
	)
)

var (
	RepositorySet = wire.NewSet(
		BlogRepositorySet,
	)
)
