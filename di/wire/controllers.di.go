package wire

import (
	"smux/internals/application/controllers/blog"

	"github.com/google/wire"
)

var (
	BlogControllerSet = wire.NewSet(
		blog.NewBlogController,
		wire.Bind(new(blog.IBlogController), new(*blog.BlogController)),
	)
)

var (
	ControllerSet = wire.NewSet(
		BlogControllerSet,
	)
)
