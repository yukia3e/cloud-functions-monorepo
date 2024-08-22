package root

import (
	"github.com/GoogleCloudPlatform/functions-framework-go/functions"

	"github.com/yukia3e/cloud-functions-monorepo/internal/config"
	eventA "github.com/yukia3e/cloud-functions-monorepo/internal/ui/function/event/a"
	eventB "github.com/yukia3e/cloud-functions-monorepo/internal/ui/function/event/b"
)

func init() {
	name := config.MustGetFunctionEntryPointName()

	switch name {
	case "function_a":
		functions.CloudEvent(
			"function_a",
			eventA.CloudEventFunc,
		)
	case "function_b":
		functions.CloudEvent(
			"function_b",
			eventB.CloudEventFunc,
		)
	}
}
