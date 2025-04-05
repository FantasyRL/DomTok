package controllers

import (
	"github.com/google/wire"
	"github.com/west2-online/domtok/app/order/controllers/rpc"
)

var ControllerProviderSet = wire.NewSet(
	rpc.NewOrderHandler,
)
