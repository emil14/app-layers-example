package grpc

import (
	"pokemonstore/api/sdk"
	"pokemonstore/internal/port/grpc/buypokemonhandler"
)

type server struct {
	sdk.UnimplementedPokemonStoreServer
	buypokemonhandler.Handler
}

func New(
	buyPokemonHandler buypokemonhandler.Handler,
) sdk.PokemonStoreServer {
	return server{
		sdk.UnimplementedPokemonStoreServer{},
		buyPokemonHandler,
	}
}
