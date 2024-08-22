package usecase

import (
	"context"

	"github.com/rs/zerolog/log"
)

func (u *usecase) FunctionB(ctx context.Context, id string) error {
	log.Debug().Str("id", id).Msg("FunctionB")
	return nil
}
