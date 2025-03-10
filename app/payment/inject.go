/*
Copyright 2024 The west2-online Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package payment

import (
	"github.com/west2-online/domtok/app/payment/controllers/rpc"
	"github.com/west2-online/domtok/app/payment/domain/service"
	"github.com/west2-online/domtok/app/payment/infrastructure/mysql"
	"github.com/west2-online/domtok/app/payment/infrastructure/redis"
	orderRpcPkg "github.com/west2-online/domtok/app/payment/infrastructure/rpc"
	"github.com/west2-online/domtok/app/payment/usecase"
	"github.com/west2-online/domtok/kitex_gen/payment"
	"github.com/west2-online/domtok/pkg/base/client"
	"github.com/west2-online/domtok/pkg/constants"
	"github.com/west2-online/domtok/pkg/logger"
	"github.com/west2-online/domtok/pkg/utils"
)

func InjectPaymentHandler() payment.PaymentService {
	// 初始化数据库存储
	gormDB, err := client.InitMySQL()
	if err != nil {
		panic(err)
	}
	db := mysql.NewPaymentDB(gormDB)

	// 初始化 Redis 客户端
	// 初始化 Redis，使用指定的 Redis DB
	redisClient, err := client.InitRedis(constants.RedisDBOrder)
	if err != nil {
		panic(err)
	}
	// 封装 Redis 存储对象
	redisRepo := redis.NewPaymentRedis(redisClient)

	// 初始化雪花接口
	sf, err := utils.NewSnowflake(0, 0)
	if err != nil {
		panic(err)
	}

	// 初始化RPC，调用别的接口的服务
	c, err := client.InitOrderRPC()
	if err != nil {
		logger.Fatalf("api.rpc.order InitOrderRPC failed, err is %v", err)
	}
	u, err := client.InitUserRPC()
	if err != nil {
		logger.Fatalf("api.rpc.user InitUserRPC failed, err is %v", err)
	}
	orderRpc := orderRpcPkg.NewPaymentRPC(*c, *u)
	// 初始化 Service，并传入 Redis
	svc := service.NewPaymentService(db, sf, redisRepo, orderRpc)

	// 初始化 UseCase，并传入 Redis
	uc := usecase.NewPaymentCase(db, svc, redisRepo, orderRpc)

	return rpc.NewPaymentHandler(uc)
}
