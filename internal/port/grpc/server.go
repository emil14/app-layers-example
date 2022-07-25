package grpc

import (
	"pokemonstore/api/sdk"
	"pokemonstore/internal/port/grpc/buypokemonhandler"
)

type server struct {
	sdk.UnimplementedPokemonStoreServer
	buypokemonhandler.Handler // остальные хендлеры можно встраивать сюда же
}

func New(
	buyPokemonHandler buypokemonhandler.Handler,
) sdk.PokemonStoreServer {
	return server{
		sdk.UnimplementedPokemonStoreServer{},
		buyPokemonHandler,
	}
}
