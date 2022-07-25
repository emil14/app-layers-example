package usecase

import (
	"fmt"

	accountdomain "pokemonstore/internal/domain/account"
	pokemondomain "pokemonstore/internal/domain/pokemon"
)

type BuyPokemon struct {
	accountRepo accountdomain.Repo
	pokemonRepo pokemondomain.Repo
}

func (uc BuyPokemon) Handle(userID, pokemonID uint64) error {
	account, err := uc.accountRepo.ByUserID(userID)
	if err != nil {
		return fmt.Errorf("account by user id: %w", err)
	}

	pokemonPrice, err := uc.pokemonRepo.Price(pokemonID)
	if err != nil {
		return fmt.Errorf("pokemon price: %w", err)
	}

	if err := uc.accountRepo.WithdrawByID(account.ID, pokemonPrice); err != nil {
		return fmt.Errorf("withdraw: %w", err)
	}

	if err := uc.pokemonRepo.SetOwner(userID, pokemonID); err != nil {
		return fmt.Errorf("set pokemon owner: %w", err)
	}

	return nil
}

func MustNewBuyPokemon(
	accountRepo accountdomain.Repo,
	pokemonRepo pokemondomain.Repo,
) BuyPokemon {
	if accountRepo == nil || pokemonRepo == nil {
		panic("nil deps")
	}

	return BuyPokemon{accountRepo, pokemonRepo}
}
