package placeholder

import (
	"platform/http"
	"platform/http/handling"
	"platform/pipeline"
	"platform/pipeline/basic"
	"platform/services"
	"sync"
)

func createPipeline() pipeline.RequestPipeline {
	return pipeline.CreatePipeline(
		&basic.ServiceComponent{},
		&basic.LoggingComponent{},
		&basic.ErrorComponent{},
		&basic.StaticFileComponent{},
		//&SimpleMessageComponent{},
		handling.NewRouter(
			handling.HandlerEntry{Prefix: "", Handler: NameHandler{}},
		),
	)
}

func Start() {
	result, err := services.Call(http.Serve, createPipeline())
	if err == nil {
		(result[0].(*sync.WaitGroup)).Wait()
	} else {
		panic(err)
	}
}
