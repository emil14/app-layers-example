package pokemon

type Repo interface {
	Price(pokemonID uint64) (float64, error)
	SetOwner(pokemonID, userID uint64) error
}
