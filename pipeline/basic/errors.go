package basic

import (
	"fmt"
	"net/http"
	"platform/logging"
	"platform/pipeline"
)

type ErrorComponent struct{}

func recoveryFunc(ctx *pipeline.ComponentContext, logger logging.Logger) {
	if arg := recover(); arg != nil {
		logger.Debugf("Error: %v", fmt.Sprint(arg))
		ctx.ResponseWriter.WriteHeader(http.StatusInternalServerError)
	}
}
func (c *ErrorComponent) ImplementsProcessRequestWithServices() {}
func (c *ErrorComponent) Init()                                 {}
func (c *ErrorComponent) ProcessRequestWithServices(ctx *pipeline.ComponentContext, next func(*pipeline.ComponentContext), logger logging.Logger) {
	defer recoveryFunc(ctx, logger)
	next(ctx)
	if ctx.GetError() != nil {
		logger.Debugf("Error: %v", ctx.GetError())
		ctx.ResponseWriter.WriteHeader(http.StatusInternalServerError)
	}
}
