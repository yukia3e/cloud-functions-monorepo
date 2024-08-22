package usecase

import (
	"context"

	"github.com/rs/zerolog/log"
)

func (u *usecase) FunctionA(ctx context.Context, id string) error {
	log.Debug().Str("id", id).Msg("FunctionA")
	return nil
}
