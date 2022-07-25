package dummy

import "pokemonstore/internal/domain/account" // реализация зависит от домена

type Repo struct{}

func (r Repo) ByUserID(userID uint64) (account.Account, error)
func (r Repo) WithdrawByID(userID uint64, amount float64) error
