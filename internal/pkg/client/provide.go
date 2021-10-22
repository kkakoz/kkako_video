package client

import "go.uber.org/fx"

var Provider = fx.Provide(NewUserClient)