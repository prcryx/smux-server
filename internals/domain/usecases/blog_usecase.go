package usecases

import (
	"smux/internals/domain/entities"
	"smux/internals/domain/repositories"
)

type IBlogUseCase interface {
	GetBlogs() []entities.BlogEntity
}

type BlogUseCase struct {
	repo repositories.BlogRepository
}

func NewBlogUseCase(repo repositories.BlogRepository) *BlogUseCase {
	return &BlogUseCase{
		repo: repo,
	}
}

func (usecase *BlogUseCase) GetBlogs() []entities.BlogEntity {
	return usecase.repo.GetBlogs()
}
