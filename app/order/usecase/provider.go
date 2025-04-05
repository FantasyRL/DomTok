package usecase

import (
	"github.com/google/wire"
)

var UsecaseProviderSet = wire.NewSet(
	NewOrderCase,
	//wire.Bind(new(OrderUseCase), new(*UseCase)),
)
