package main

import (
	"context"
	"github.com/google/wire"
	ex "github.com/west2-online/domtok/app/order"
	"github.com/west2-online/domtok/app/order/controllers"
	"github.com/west2-online/domtok/app/order/domain"
	"github.com/west2-online/domtok/app/order/infrastructure"
	"github.com/west2-online/domtok/app/order/usecase"
	"github.com/west2-online/domtok/config"
	"github.com/west2-online/domtok/kitex_gen/order"
	"github.com/west2-online/domtok/pkg/logger"
)

// 正常是要返回一个server实例的封装，但我们目前没有把server的需要的各种初始化opt封装起来，所以这里只能从controller开始注入了
func InitializeApp(ctx context.Context, svcName string, level string) order.OrderService {
	panic(wire.Build(
		config.Init, // 理应返回一个cfg给后面的初始化调用
		logger.Init,
		ex.ExternalProviderSet,
		controllers.ControllerProviderSet,
		usecase.UsecaseProviderSet,
		domain.DomainProviderSet,
		infrastructure.InfraProviderSet,
	))
}
