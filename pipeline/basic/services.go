package basic

import (
	"platform/pipeline"
	"platform/services"
)

type ServiceComponent struct{}

func (c *ServiceComponent) Init() {}
func (c *ServiceComponent) ProcessRequest(ctx *pipeline.ComponentContext, next func(*pipeline.ComponentContext)) {
	reqContext := ctx.Request.Context()
	_ = ctx.Request.WithContext(services.NewServiceContext(reqContext))
	next(ctx)
}
