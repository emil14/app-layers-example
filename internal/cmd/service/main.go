package main

import (
	"log"
	"net"

	"pokemonstore/api/sdk"
	accountDummyRepo "pokemonstore/internal/domain/account/repo/dummy"
	pokemonDummyRepo "pokemonstore/internal/domain/pokemon/repo/dummy"
	grpcPort "pokemonstore/internal/port/grpc"
	"pokemonstore/internal/port/grpc/buypokemonhandler"
	"pokemonstore/internal/usecase"

	"google.golang.org/grpc"
)

type uc struct {
	usecase.BuyPokemon
}

func main() {
	log.Println("Starting listening on port 8080")
	port := ":8080"

	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	log.Printf("Listening on %s", port)

	grpcSrv := grpc.NewServer()
	sdk.RegisterPokemonStoreServer(grpcSrv, buildGRPCApp())

	if err := grpcSrv.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func buildGRPCApp() sdk.PokemonStoreServer {
	buyPokemonUseCase := usecase.MustNewBuyPokemon(
		accountDummyRepo.Repo{},
		pokemonDummyRepo.Repo{},
	)

	buyPokemonHandler := buypokemonhandler.MustNew(buyPokemonUseCase)

	return grpcPort.New(buyPokemonHandler)
}
