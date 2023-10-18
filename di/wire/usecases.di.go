package wire

import (
	"smux/internals/domain/usecases"

	"github.com/google/wire"
)

var (
	BlogUseCaseSet = wire.NewSet(
		usecases.NewBlogUseCase,
		wire.Bind(new(usecases.IBlogUseCase), new(*usecases.BlogUseCase)),
	)
)

var (
	UseCaseSet = wire.NewSet(
		BlogUseCaseSet,
	)
)
