package domain

import (
	"github.com/google/wire"
	"github.com/west2-online/domtok/app/order/domain/service"
)

var DomainProviderSet = wire.NewSet(
	service.NewOrderService,
)
