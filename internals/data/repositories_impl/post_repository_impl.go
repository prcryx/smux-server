package repositories_impl

import (
	"smux/internals/data/datasources"
	"smux/internals/domain/entities"
)

type BlogRepositoryImpl struct {
	src *datasources.TestDataSource
}

func NewBlogRepositoryImpl(dataSrc *datasources.TestDataSource) *BlogRepositoryImpl {
	return &BlogRepositoryImpl{
		src: dataSrc,
	}
}

func (repoImpl *BlogRepositoryImpl) GetBlogs() []entities.BlogEntity {
	return repoImpl.src.GetBlogs()
}
