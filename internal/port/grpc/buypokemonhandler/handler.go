package buypokemonhandler

import (
	"context"
	"log"

	"pokemonstore/api/sdk"

	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc/status"
)

type buyPokemonUseCase interface { // на каждый юзкейс придётся добавлять такой интерфейс
	Handle(userID, pokemonID uint64) error
}

type Handler struct {
	buyPokemonUseCase buyPokemonUseCase // Инверсия зависимости. Можно замокать юзкейс и написать тест
}

func (s *Handler) BuyPokemon(ctx context.Context, in *sdk.BuyRequest) (*empty.Empty, error) {
	if err := s.buyPokemonUseCase.Handle(in.UserId, in.PokemonId); err != nil { // вызываем юзкейсы из порта
		log.Printf("buy pokemon usecase: %v", err)
		return nil, status.Error(500, "something went wrong")
	}

	return &empty.Empty{}, nil
}

func MustNew(usecases buyPokemonUseCase) Handler {
	if usecases == nil {
		panic("nil usecases")
	}
	return Handler{usecases}
}
