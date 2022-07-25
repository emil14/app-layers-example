package buypokemonhandler

import (
	"context"
	"log"

	"pokemonstore/api/sdk"

	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc/status"
)

type usecases interface {
	BuyPokemon(userID, pokemonID uint64) error
}

type Handler struct {
	UseCases usecases
}

func (s *Handler) BuyPokemon(ctx context.Context, in *sdk.BuyRequest) (*empty.Empty, error) {
	if err := s.UseCases.BuyPokemon(in.UserId, in.PokemonId); err != nil {
		log.Printf("buy pokemon usecase: %v", err)
		return nil, status.Error(500, "something went wrong")
	}

	return &empty.Empty{}, nil
}

func MustNew(usecases usecases) Handler {
	if usecases == nil {
		panic("usecases == nil")
	}
	return Handler{usecases}
}
