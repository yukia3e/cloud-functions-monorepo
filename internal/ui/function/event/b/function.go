package event

import (
	"context"
	"fmt"

	cloudevents "github.com/cloudevents/sdk-go/v2"
	"github.com/rs/zerolog/log"

	"github.com/yukia3e/cloud-functions-monorepo/internal/domain/model/message"
	"github.com/yukia3e/cloud-functions-monorepo/internal/usecase"
)

func CloudEventFunc(ctx context.Context, e cloudevents.Event) error {
	var msg message.TopicBMessage
	err := e.DataAs(&msg)
	if err != nil {
		log.Error().Msg(fmt.Errorf("failed to convert data: %w", err).Error())

		// Return nil to prevent re-execution because it is fatal
		return nil
	}

	usecase := usecase.New()
	if err := usecase.FunctionB(ctx, msg.ID); err != nil {
		return err
	}

	return nil
}
