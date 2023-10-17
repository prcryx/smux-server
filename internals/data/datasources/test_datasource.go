package datasources

import (
	"smux/internals/domain/entities"
	"time"
)

type ITestDataSource interface {
	GetBlogs() []entities.BlogEntity
}
type TestDataSource struct {
}

func Init() (*TestDataSource, error) {
	return &TestDataSource{}, nil
}

func (tds *TestDataSource) GetBlogs() []entities.BlogEntity {
	return []entities.BlogEntity{
		{
			BlogId:    1,
			Title:     "Benchmarking in Golang",
			Body:      "A benchmark is a type of function that executes a code segment multiple times and compares each output against a standard, assessing the code's overall performance level Golang includes built-in tools for writing benchmarks in the testing package and the go tool, so you can write useful benchmarks without installing any dependencies.",
			CreatedAt: time.Date(2021, 2, 12, 14, 20, 30, 500, time.Local),
			AuthorId:  0,
			Author:    entities.Author{},
		},
		{
			BlogId:    2,
			Title:     "Go for Cloud & Network Services",
			Body:      "et tutorials and walkthroughs. Cloud computing basics. Learn more about ... Quickly build scalable apps using Go programming language on Google Cloud. Try it ",
			CreatedAt: time.Date(2022, 7, 9, 17, 20, 30, 500, time.Local),
		},
	}
}
