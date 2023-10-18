package container

import "smux/internals/application/controllers/blog"

// this will grow over time
type ControllerRegistry struct {
	BlogController *blog.BlogController
}

// add new controllers to the ControllerRegistry
func NewControllerRegistry(blogController *blog.BlogController) *ControllerRegistry {
	return &ControllerRegistry{
		BlogController: blogController,
	}
}
