package wire

import (
	"smux/internals/data/datasources"

	"github.com/google/wire"
)

var (
	TestDataSourceSet = wire.NewSet(
		datasources.Init,
		wire.Bind(new(datasources.ITestDataSource), new(*datasources.TestDataSource)),
	)
)

var (
	DataSourceSet = wire.NewSet(
		TestDataSourceSet,
	)
)
