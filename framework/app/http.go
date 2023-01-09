package app

import "context"

type IController interface {
	Init(ctx context.Context)
}

type BaseController struct {
	Ctx context.Context
}

type BaseRequest struct {
	context.Context
}

func (c *BaseController) Init(ctx context.Context) {
	c.Ctx = ctx
}
