package dummy

type Repo struct{}

func (r Repo) Price(pokemonID uint64) (float64, error)
func (r Repo) SetOwner(pokemonID, userID uint64) error
