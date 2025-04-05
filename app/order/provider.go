package order

import (
	"github.com/google/wire"
	"github.com/west2-online/domtok/app/order/infrastructure/mysql"
	"github.com/west2-online/domtok/pkg/base/client"
)

// ExternalProviderSet 姑且称作外部依赖的provider
var ExternalProviderSet = wire.NewSet(
	client.InitMySQL,
	mysql.NewOrderDB,
	client.InitUserRPC,
	client.InitCommodityRPC,
)
