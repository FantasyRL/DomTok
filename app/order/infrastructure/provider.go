package infrastructure

import (
	"github.com/google/wire"
	"github.com/west2-online/domtok/app/order/domain/repository"
	"github.com/west2-online/domtok/app/order/infrastructure/locker"
	"github.com/west2-online/domtok/app/order/infrastructure/mq"
	"github.com/west2-online/domtok/app/order/infrastructure/mysql"
	"github.com/west2-online/domtok/app/order/infrastructure/redis"
	"github.com/west2-online/domtok/app/order/infrastructure/rpc"
	"github.com/west2-online/domtok/pkg/utils"
)

var InfraProviderSet = wire.NewSet(
	mysql.NewOrderDB,
	wire.Bind(new(repository.OrderDB), new(*mysql.OrderDB)),
	mq.NewRocketmq,
	wire.Bind(new(repository.MQ), new(*mq.RocketMq)),
	rpc.NewOrderRpcImpl,
	wire.Bind(new(repository.RPC), new(*rpc.OrderRpcImpl)),
	redis.NewOrderCache,
	wire.Bind(new(repository.Cache), new(*redis.OrderCache)),
	locker.NewLocker,
	wire.Bind(new(repository.Locker), new(*locker.Locker)),
	utils.NewSnowflake,
	wire.Bind(new(repository.IDGenerator), new(*utils.Snowflake)),
)
