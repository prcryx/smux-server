package repositories

import "smux/internals/domain/entities"

type BlogRepository interface {
	GetBlogs() []entities.BlogEntity
}
