package httpapi

import (
	"github.com/allentom/haruka"
	"myservice/service"
)

var helloHandler haruka.RequestHandler = func(context *haruka.Context) {
	context.JSON(haruka.JSON{
		"message": service.ExampleService(),
	})
}
